package flag

import (
	"gvb_server/models"
)

func EsCreateIndex() {
	models.ArticleModel{}.CreatIndex()
	models.FullTextModel{}.CreatIndex()
}
