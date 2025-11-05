package common

import (
	"fmt"
	"gorm.io/gorm"
	"gvb_server/global"
	"gvb_server/models"
)

type Option struct {
	models.PageInfo
	Debug bool
	Likes []string //模糊匹配字段
}

func ComList[T any](model T, option Option) (list []T, count int64, err error) {

	//count := global.DB.Find(&imageList).RowsAffected
	//fmt.Println(count)
	DB := global.DB
	if option.Debug {
		DB = global.DB.Session(&gorm.Session{Logger: global.MysqlLog})
	}
	if option.Sort == "" {
		option.Sort = "created_at asc" // 默认按照时间往前排
	}
	DB = DB.Where(model)
	for index, column := range option.Likes {
		if index == 0 {
			DB.Where(fmt.Sprintf("%s like ?", column), fmt.Sprintf("%%%s%%", option.Key))
		}
		DB.Or(fmt.Sprintf("%s like ?", column), fmt.Sprintf("%%%s%%", option.Key))
	}
	//DB = global.DB.Session(&gorm.Session{Logger: global.MysqlLog})
	//DB.Model(model).Count(&count)
	count = DB.Find(&list).RowsAffected
	query := DB.Where(model)

	offset := (option.Page - 1) * option.Limit
	if offset < 0 {
		offset = 0
	}
	if option.Limit == 0 {
		option.Limit = -1
	}
	err = query.Limit(option.Limit).Offset(offset).Order(option.Sort).Find(&list).Error

	return list, count, err
	//return list, count, err
}
