package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/saltbo/gopkg/ginutil"
	"github.com/spf13/viper"

	"github.com/zplat-core/apiserver/api/v1"
	_ "github.com/zplat-core/apiserver/docs"
)

// @title Zplat API
// @version 1.0.0
// @description This is a zplat api-server api.

// @contact.name Zplat Support
// @contact.url https://saltbo.cn
// @contact.email saltbo@foxmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /api
func SetupServerRoutes(ge *gin.Engine) {
	apiRouter := ge.Group("/api/v1")
	apiRouter.Use(func(c *gin.Context) {
		if viper.ConfigFileUsed() == "" {
			ginutil.JSONError(c, http.StatusInternalServerError, fmt.Errorf("system is not initialized"))
			return
		}
	})

	ginutil.SetupResource(apiRouter,
		v1.NewConfigResource(),
		v1.NewTokenResource(),
		v1.NewUserResource(),
	)
}
