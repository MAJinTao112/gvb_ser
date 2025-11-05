package models

import "time"

type UserCollectModel struct {
	UserID    uint      `gorm:"primary_key"`
	UserModel UserModel `gorm:"foreignKey:UserID"`
	ArticleID string    `gorm:"size:32;primary_key"`
	CreatedAt time.Time
}
