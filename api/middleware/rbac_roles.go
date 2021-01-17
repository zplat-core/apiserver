package middleware

import (
	"github.com/storyicon/grbac"
	"github.com/storyicon/grbac/pkg/meta"
)

var defaultRules = grbac.Rules{
	{
		Resource: &meta.Resource{
			Host:   "*",
			Path:   "**",
			Method: "*",
		},
		Permission: &meta.Permission{
			AuthorizedRoles: []string{"admin", "member"},
		},
	},
	{
		Resource: &meta.Resource{
			Host:   "*",
			Path:   "/api/moreu/tokens",
			Method: "POST",
		},
		Permission: &meta.Permission{
			AllowAnyone: true,
		},
	},
	{
		Resource: &meta.Resource{
			Host:   "*",
			Path:   "/api/moreu/users",
			Method: "GET",
		},
		Permission: &meta.Permission{
			AuthorizedRoles: []string{"admin"},
		},
	},
}
