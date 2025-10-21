package main

import (
	"context"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"strconv"
	"time"
)

func EsConnect() *elastic.Client {
	sniffOpt := elastic.SetSniff(false)
	host := "http://127.0.0.1:9200"
	c, err := elastic.NewClient(
		elastic.SetURL(host),
		sniffOpt,
		elastic.SetBasicAuth("", ""),
	)
	if err != nil {
		logrus.Fatalf("es连接失败 %s", err.Error())
	}
	return c
}

// 插入数据到已存在的索引，并打印实际文档ID
func InsertData() {
	client := EsConnect()
	indexName := "users"

	users := []map[string]interface{}{
		{"username": "Alice", "age": 25, "created_at": time.Now()},
		{"username": "Bob", "age": 30, "created_at": time.Now()},
		{"username": "Charlie", "age": 28, "created_at": time.Now()},
	}

	for i, user := range users {
		// 使用 strconv.Itoa 将数字正确转换为字符串作为 ID
		docID := strconv.Itoa(i + 1)
		resp, err := client.Index().
			Index(indexName).
			Id(docID).
			BodyJson(user).
			Do(context.Background())
		if err != nil {
			logrus.Errorf("插入数据失败: %s", err)
		} else {
			logrus.Infof("插入数据成功: %+v, 文档ID: %s", user, resp.Id)
		}
	}
}

func main() {
	InsertData()
}
