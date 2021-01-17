package system

import (
	"github.com/saltbo/gopkg/strutil"

	"github.com/zplat-core/apiserver/service"
)

func InvitationOnly() bool {
	invitation, err := service.NewOption().Get(service.OptInvitation)
	if err != nil {
		return false
	}

	return strutil.BoolFromStr(invitation, false)
}

func EmailActed() bool {
	invitation, err := service.NewOption().Get(service.OptEmailAct)
	if err != nil {
		return false
	}

	return strutil.BoolFromStr(invitation, false)
}
