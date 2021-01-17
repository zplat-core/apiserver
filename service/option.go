package service

import (
	"github.com/saltbo/gopkg/gormutil"

	"github.com/zplat-core/apiserver/model"
)

const (
	OptEmailAct   = "EMAIL_ACT"
	OptInvitation = "INVITATION"
)

type Option struct {
}

func NewOption() *Option {
	return &Option{}
}

func (o *Option) Get(key string) (string, error) {
	opt := new(model.Option)
	if err := gormutil.DB().First(opt, "key=?", key).Error; err != nil {
		return "", err
	}

	return opt.Value, nil
}

func (o *Option) Set(key, value string) error {
	opt := &model.Option{Key: key, Value: key}
	if err := gormutil.DB().Save(opt).Error; err != nil {
		return err
	}

	return nil
}
