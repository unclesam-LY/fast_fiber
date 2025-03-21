package models

import "time"

type Model struct {
	ID        uint      `gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"-"`
}

type PageInfo struct {
	Page  int    `form:"page" default:"1"`
	Key   string `form:"key"`
	Limit int    `form:"limit" default:"0"`
	Sort  string `form:"sort"`
}

type RemoveRequest struct {
	IDList []uint `json:"id_list"`
}
