package models

import "time"

type UserCollectModel struct {
	UserID    uint      `gorm:"primary_key"`
	UserModel UserModel `gorm:"foreignKey:UserID"`
	ArticleID uint      `gorm:"primary_key"`
	CreatedAt time.Time
}
