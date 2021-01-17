package model

import "time"

type Option struct {
	Id      int64      `json:"id"`
	Key     string     `json:"key" gorm:"not null"`
	Value   string     `json:"value" gorm:"not null"`
	Deleted *time.Time `json:"-" gorm:"column:deleted_at"`
	Created time.Time  `json:"created" gorm:"column:created_at;not null"`
	Updated time.Time  `json:"updated" gorm:"column:updated_at;not null"`
}
