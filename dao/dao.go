package dao

import (
	"github.com/saltbo/gopkg/gormutil"

	"github.com/zplat-core/apiserver/model"
)

func Init(driver, dsn string) {
	gormutil.Init(gormutil.Config{
		Driver: driver,
		DSN:    dsn,
	}, true)
	gormutil.AutoMigrate(model.Tables())
}
