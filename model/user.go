package model

import (
	"strings"
	"time"
)

const (
	RoleAdmin  = "admin"
	RoleMember = "member"
	RoleGuest  = "guest"
)

const (
	StatusInactivated = iota
	StatusActivated
	StatusDisabled
)

var roles = map[string]string{
	RoleAdmin:  "管理员",
	RoleMember: "注册用户",
	RoleGuest:  "游客",
}

var status = map[uint8]string{
	StatusInactivated: "未激活",
	StatusActivated:   "已激活",
	StatusDisabled:    "已禁用",
}

type User struct {
	Id        int64      `json:"id"`
	Ux        string     `json:"ux" gorm:"size:32;unique_index;not null"` // Global unique user ID
	Email     string     `json:"email" gorm:"size:32;unique_index;not null"`
	Username  string     `json:"username" gorm:"size:20;unique_index;not null"`
	Password  string     `json:"-" gorm:"size:32;not null"`
	Status    uint8      `json:"-" gorm:"size:1;not null"`
	StatusTxt string     `json:"status" gorm:"-"`
	Roles     string     `json:"-" gorm:"size:64;not null"`
	RoleTxt   string     `json:"role" gorm:"-"`
	Ticket    string     `json:"ticket" gorm:"size:6;unique_index;not null"`
	Deleted   *time.Time `json:"-" gorm:"column:deleted_at"`
	Created   time.Time  `json:"created" gorm:"column:created_at;not null"`
	Updated   time.Time  `json:"updated" gorm:"column:updated_at;not null"`
}

func (User) TableName() string {
	return "mu_user"
}

func (u *User) Activated() bool {
	return u.Status == StatusActivated
}

func (u *User) RolesSplit() []string {
	return strings.Split(u.Roles, ",")
}

func (u *User) Format() *User {
	u.RoleTxt = roles[u.Roles]
	u.StatusTxt = status[u.Status]
	return u
}

type UserFormats struct {
	Id        int64  `json:"id"`
	Ux        string `json:"ux" gorm:"size:32;unique_index;not null"` // Global unique user ID
	Email     string `json:"email" gorm:"size:32;unique_index;not null"`
	Username  string `json:"username" gorm:"size:20;unique_index;not null"`
	Roles     string `json:"-" gorm:"size:64;not null"`
	RoleName  string `json:"role" gorm:"-"`
	Status    uint8  `json:"-" gorm:"size:1;not null"`
	StatusTxt string `json:"status" gorm:"-"`
	UserProfile
}

func (u *UserFormats) Format() *UserFormats {
	u.RoleName = roles[u.Roles]
	u.StatusTxt = status[u.Status]
	return u
}
