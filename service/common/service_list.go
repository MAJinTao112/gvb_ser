package common

import (
	"gorm.io/gorm"
	"gvb_server/global"
	"gvb_server/models"
)

type Option struct {
	models.PageInfo
	Debug bool
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
	query := DB.Where(model)

	//DB = global.DB.Session(&gorm.Session{Logger: global.MysqlLog})
	//DB.Model(model).Count(&count)
	count = query.Select("id").Find(&list).RowsAffected
	query = DB.Where(model)
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
