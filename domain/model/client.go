package model

import (
	"time"

	"gorm.io/gorm"
)

type Client struct {
	ID        uint64         `gorm:"primaryKey;autoIncrement" json:"id"`
	Code      string         `gorm:"<-" json:"code"`
	Name      string         `gorm:"<-" json:"name"`
	Key       string         `gorm:"<-" json:"key"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func (Client) TableName() string {
	return "clients"
}

type ClientFilter struct {
	Name string `json:"name"`
}

type ClientFilterParams struct {
	Page  uint              `query:"page" form:"page"`
	Limit uint              `query:"limit" form:"limit"`
	Sort  map[string]string `query:"sort" form:"sort"`
	Name  string            `query:"name" form:"name"`
}

type ClientCreatePayload struct {
	Name string `json:"name"`
}

type ClientExists struct {
	Code string
	Name string
	Key  string
}
