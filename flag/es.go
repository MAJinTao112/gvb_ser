package flag

import (
	"fmt"
	"gvb_server/models"
	"log"
)

func EsCreateIndex() {
	err := models.ArticleModel{}.CreatIndex()
	fmt.Println(err)
	fmt.Println("EsCreateIndex")
	if err != nil {
		log.Fatal(err)
	}
}
