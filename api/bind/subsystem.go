package bind

import (
	"github.com/zplat-core/apiserver/model"
)

type QuerySubsystem struct {
	QueryPage

	Name string `form:"name"`
}

type BodySubsystem struct {
	Name    string `json:"name"  binding:"required"`
	Intro   string `json:"intro"  binding:"required"`
	Address string `json:"address"  binding:"required"`
}

func (p *BodySubsystem) Model() *model.Subsystem {
	return &model.Subsystem{
		Name:    p.Name,
		Intro:   p.Intro,
		Address: p.Address,
	}
}
