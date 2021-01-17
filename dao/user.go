package dao

import (
	"fmt"

	"github.com/saltbo/gopkg/gormutil"

	"github.com/zplat-core/apiserver/model"
)

type User struct {
}

func NewUser() *User {
	return &User{}
}

func (u *User) FindAll(query *Query) (list []model.UserFormats, total int64, err error) {
	ut := model.User{}.TableName()
	pt := model.UserProfile{}.TableName()
	joinQuery := fmt.Sprintf("left join %s on %s.ux = %s.ux", pt, pt, ut)
	sn := gormutil.DB().Table(ut)
	if len(query.Params) > 0 {
		sn = sn.Where(query.SQL(), query.Params...)
	}
	sn.Count(&total)
	err = sn.Offset(query.Offset).Limit(query.Limit).Select("*").Joins(joinQuery).Find(&list).Error
	for idx, item := range list {
		list[idx] = *item.Format()
	}
	return
}
