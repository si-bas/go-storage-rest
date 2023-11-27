package model

import (
	"time"

	"gorm.io/gorm"
)

type File struct {
	ID           uint64         `gorm:"primaryKey;autoIncrement" json:"id"`
	ClientID     uint64         `gorm:"<-" json:"-"`
	Code         string         `gorm:"<-;unique" json:"-"`
	OriginalName string         `gorm:"<-" json:"-"`
	Name         string         `gorm:"<-" json:"name"`
	Extension    string         `gorm:"<-" json:"extension"`
	Size         int64          `gorm:"<-" json:"size"`
	Path         string         `gorm:"<-" json:"-"`
	CreatedAt    time.Time      `json:"created_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at"`
}

func (File) TableName() string {
	return "files"
}

type FileCreate struct {
	OriginalName string
	Name         string
	Extension    string
	Size         int64
	Path         string
}
