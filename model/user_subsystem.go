package model

import "time"

type UserSubsystem struct {
	Id      int64                  `json:"id"`
	Ux      string                 `json:"ux" gorm:"size:32;unique_index;not null"`
	Opts    map[string]interface{} `json:"location" gorm:"size:32;not null"`
	Deleted *time.Time             `json:"-" gorm:"column:deleted_at"`
	Created time.Time              `json:"created" gorm:"column:created_at;not null"`
	Updated time.Time              `json:"updated" gorm:"column:updated_at;not null"`
}

func (UserSubsystem) TableName() string {
	return "mu_user_subsystem"
}
