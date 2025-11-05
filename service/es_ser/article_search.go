package es_ser

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/service/redis_ser"
	"strings"
)

func CommList(option Option) (list []models.ArticleModel, count int, err error) {
	boolSearch := elastic.NewBoolQuery()
	if option.Key != "" {
		boolSearch.Must(
			elastic.NewMultiMatchQuery(option.Key, option.Fields...),
		)
	}
	if option.Tag != "" {
		boolSearch.Must(
			elastic.NewMultiMatchQuery(option.Tag, "tags"))
	}
	type SortField struct {
		Field     string
		Ascending bool
	}
	sortFields := SortField{
		Field:     "created_at",
		Ascending: true,
	}
	if option.Sort != "" {
		_list := strings.Split(option.Sort, " ")
		if len(_list) == 2 && (_list[1] == "desc" || _list[1] == "asc") {
			sortFields.Field = _list[0]
			if _list[1] == "desc" {
				sortFields.Ascending = false
			}
			if _list[1] == "asc" {
				sortFields.Ascending = true

			}
		}
	}
	res, err := global.ESClient.Search(models.ArticleModel{}.Index()).Query(boolSearch).
		Highlight(elastic.NewHighlight().Field("title")).FetchSourceContext(
		elastic.NewFetchSourceContext(true).Exclude("content")).
		From(option.GetForm()).Sort(sortFields.Field, sortFields.Ascending).Size(option.Limit).Do(context.Background())
	if err != nil {
		logrus.Error(err)
		return
	}
	count = int(res.Hits.TotalHits.Value)
	domoList := []models.ArticleModel{}

	diggInfo := redis_ser.NewDigg().GetInfo()
	lookInfo := redis_ser.NewArticleLook().GetInfo()
	commentInfo := redis_ser.NewCommentCount().GetInfo()
	for _, hit := range res.Hits.Hits {
		var model models.ArticleModel
		data, err := hit.Source.MarshalJSON()
		if err != nil {
			logrus.Error(err)
			continue
		}

		err = json.Unmarshal(data, &model)
		if err != nil {
			logrus.Error(err)
			continue
		}
		title, ok := hit.Highlight["title"]
		if ok {
			model.Title = title[0]
		}
		//fmt.Println(hit.Highlight)
		model.ID = hit.Id
		digg := diggInfo[hit.Id]
		look := lookInfo[hit.Id]
		comment := commentInfo[hit.Id]
		model.DiggCount = digg + model.DiggCount
		model.LookCount = model.LookCount + look
		model.CommentCount = model.CommentCount + comment
		domoList = append(domoList, model)

	}
	return domoList, count, err
}

func CommDetail(id string) (model models.ArticleModel, err error) {
	res, err := global.ESClient.Get().Index(models.ArticleModel{}.Index()).Id(id).Do(context.Background())
	if err != nil {
		logrus.Error(err)
		return
	}
	err = json.Unmarshal(res.Source, &model)
	if err != nil {
		return
	}
	model.ID = res.Id
	model.LookCount = model.LookCount + redis_ser.NewArticleLook().Get(res.Id)
	return model, err
}

func CommDetailByKeyword(key string) (model models.ArticleModel, err error) {
	res, err := global.ESClient.Search().
		Index(models.ArticleModel{}.Index()).
		Query(elastic.NewTermQuery("keyword", key)).
		Size(1).
		Do(context.Background())
	if err != nil {
		return
	}
	if res.Hits.TotalHits.Value == 0 {
		return model, errors.New("文章不存在")
	}
	hit := res.Hits.Hits[0]

	err = json.Unmarshal(hit.Source, &model)
	if err != nil {
		return
	}
	model.ID = hit.Id
	return
}
func ArticleUpdate(id string, data map[string]any) error {
	_, err := global.ESClient.Update().
		Index(models.ArticleModel{}.Index()).
		Id(id).
		Doc(data).Refresh("true").
		Do(context.Background())

	return err
}
