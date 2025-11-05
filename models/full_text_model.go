package models

import (
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"gvb_server/global"
)

type FullTextModel struct {
	ID    string `json:"id" structs:"id"`       // es的id
	Key   string `json:"key" structs:"key"`     //文章关联的id
	Title string `json:"title" structs:"title"` // 文章标题
	Slug  string `json:"slug" structs:"slug"`   //标题的跳转地址
	Body  string `json:"body" structs:"body"`   //文章内容
}

func (FullTextModel) Index() string {
	return "full_text_index"
}
func (FullTextModel) Mapping() string {
	return `{
		"settings": {
			"index":{
				"max_result_window": "100000"
			}
		},
		"mappings": {
			"properties": {
                "key": {
					"type": "keyword"
				},
				"title": {
					"type": "text"
				},
				"slug": {
					"type": "keyword"
				},
				"body": {
					"type": "text"
				}
				
			}
		}
	}
	`
}
func (a FullTextModel) IndexExists() bool {
	exists, err := global.ESClient.IndexExists(a.Index()).Do(context.Background())
	if err != nil {
		logrus.Error(err.Error())
		return exists
	}
	return exists
}

func (a FullTextModel) CreatIndex() error {
	if a.IndexExists() {
		err := a.RemoveIndex()
		if err != nil {
			return err
		}
	}
	creatIndex, err := global.ESClient.CreateIndex(a.Index()).BodyString(a.Mapping()).Do(context.Background())
	if err != nil {
		logrus.Error("创建索引失败")
		logrus.Error(err.Error())
		return err
	}
	if !creatIndex.Acknowledged {
		logrus.Error("创建失败")
		return err
	}
	logrus.Infof("索引%s创建成功", a.Index())
	return nil
}
func (a FullTextModel) RemoveIndex() error {
	logrus.Info("索引存在，删除索引")
	indexDelete, err := global.ESClient.DeleteIndex(a.Index()).Do(context.Background())
	if err != nil {
		logrus.Error("删除索引失败")
		logrus.Error(err.Error())
		return err
	}
	if !indexDelete.Acknowledged {
		logrus.Error("索引删除失败")
		return err
	}
	logrus.Info("索引删除成功")
	return nil
}
func (a FullTextModel) Create() error {
	indexResponse, err := global.ESClient.Index().Index(a.Index()).BodyJson(a).Do(context.Background())
	if err != nil {
		logrus.Error(err.Error())
		return err
	}
	a.ID = indexResponse.Id
	return nil
}

// 是否存在该文章
func (a FullTextModel) ISExistData() bool {
	res, err := global.ESClient.Search(a.Index()).Query(elastic.NewMatchQuery("keyword", a.Title)).
		Size(1).Do(context.Background())
	if err != nil {
		logrus.Error(err.Error())
		return false
	}
	if res.Hits.TotalHits.Value > 0 {
		return true
	}
	return false
}
func (a FullTextModel) GetDataByID(id string) error {
	res, err := global.ESClient.Get().Index(a.Index()).Id(id).Do(context.Background())
	if err != nil {
		logrus.Error(err.Error())
		return err
	}
	err = json.Unmarshal(res.Source, &a)
	if err != nil {
		return err
	}

	return nil
}
