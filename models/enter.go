package models

import "time"

type MODEL struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"-"`
}

//type MODEL struct {
//	ID        uint      `gorm:"primaryKey" json:"id"`
//	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
//	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"-"`
//}

type PageInfo struct {
	Page  int    `form:"page"`
	Key   string `form:"key"`
	Limit int    `form:"limit"`
	Sort  string `form:"sort"`
}
type RemoveRequest struct {
	IDList []uint `json:"id_list"`
}
type ESIDRequest struct {
	ID string `json:"id" form:"id" uri:"id"`
}
type ESIDListRequest struct {
	IDList []string `json:"id_list"`
}
