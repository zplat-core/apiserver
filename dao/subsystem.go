package dao

import (
	"github.com/saltbo/gopkg/gormutil"

	"github.com/zplat-core/apiserver/model"
)

type Subsystem struct {
}

func NewSubsystem() *Subsystem {
	return &Subsystem{}
}

func (s *Subsystem) FindAll(query *Query) (rets []model.Subsystem, total int64, err error) {
	sn := gormutil.DB().Where(query.SQL(), query.Params...)
	sn.Model(model.Subsystem{}).Count(&total)
	rets = make([]model.Subsystem, 0)
	err = sn.Offset(query.Offset).Limit(query.Limit).Find(&rets).Error
	return
}

func (s *Subsystem) Add(subsystem *model.Subsystem) error {
	return gormutil.DB().Create(subsystem).Error
}

func (s *Subsystem) Update(id int64, subsystem *model.Subsystem) error {
	return gormutil.DB().Model(subsystem).Where("id=?", id).Update(subsystem).Error
}

func (s *Subsystem) Delete(id int64) error {
	return gormutil.DB().Delete(&model.Subsystem{}, "id=?", id).Error
}
