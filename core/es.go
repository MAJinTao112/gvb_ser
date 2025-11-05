package core

import (
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"gvb_server/global"
	"log"
)

func EsConnect() *elastic.Client {
	var err error
	sniffOpt := elastic.SetSniff(false)
	host := global.Config.Es.Url()
	c, err := elastic.NewClient(elastic.SetURL(host), sniffOpt, elastic.SetBasicAuth("", ""))
	if err != nil {
		logrus.Fatalf("connect elastic err:%v", err)
	}
	log.Println("connect elastic success")
	return c
}
