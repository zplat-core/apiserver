package model

import (
	"time"
)

type UserInvitation struct {
	Id      int64     `json:"id"`
	Ux      string    `json:"ux" gorm:"size:32;unique_index;not null"`
	SubUx   string    `json:"sub_ux" gorm:"size:32;unique_index;not null"`
	Created time.Time `json:"created" gorm:"column:created_at;not null"`
	Updated time.Time `json:"updated" gorm:"column:updated_at;not null"`
}

func (UserInvitation) TableName() string {
	return "mu_user_invitation"
}
