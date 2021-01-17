package model

import "time"

type Subsystem struct {
	Id           int64      `json:"id"`
	Name         string     `json:"name" gorm:"size:32;not null"`
	Intro        string     `json:"intro"`
	Address      string     `json:"address"`
	AccessKey    string     `json:"access_key"`
	AccessSecret string     `json:"access_secret"`
	Deleted      *time.Time `json:"-" gorm:"column:deleted_at"`
	Created      time.Time  `json:"created" gorm:"column:created_at;not null"`
	Updated      time.Time  `json:"updated" gorm:"column:updated_at;not null"`
}
