package v1

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/saltbo/gopkg/ginutil"
	_ "github.com/saltbo/gopkg/httputil"
	"github.com/spf13/viper"
)

type ConfigResource struct {
}

func NewConfigResource() *ConfigResource {
	return &ConfigResource{}
}

func (rs *ConfigResource) Register(router *gin.RouterGroup) {
	router.GET("/configs/:key", rs.find)
	router.PUT("/configs/:key", rs.update)
	router.DELETE("/configs/:key", rs.delete)
}

// find godoc
// @Tags v1/Configs
// @Summary 获取配置项
// @Description 根据键名获取配置项
// @Accept json
// @Produce json
// @Param key path string true "键名"
// @Success 200 {object} httputil.JSONResponse
// @Failure 400 {object} httputil.JSONResponse
// @Failure 500 {object} httputil.JSONResponse
// @Router /v1/configs/{key} [get]
func (rs *ConfigResource) find(c *gin.Context) {
	key := c.Param("key")
	if !viper.IsSet(key) {
		ginutil.JSONBadRequest(c, fmt.Errorf("key %s not exist", key))
		return
	}

	ginutil.JSONData(c, viper.Get(key))
}

// update godoc
// @Tags v1/Configs
// @Summary 修改配置项
// @Description 根据键名修改配置项
// @Accept json
// @Produce json
// @Param key path string true "键名"
// @Param body body map[string]interface{} true "参数"
// @Success 200 {object} httputil.JSONResponse
// @Failure 400 {object} httputil.JSONResponse
// @Failure 500 {object} httputil.JSONResponse
// @Router /v1/configs/{key} [put]
func (rs *ConfigResource) update(c *gin.Context) {
	p := make(map[string]interface{}, 0)
	if err := c.ShouldBind(&p); err != nil {
		ginutil.JSONBadRequest(c, err)
		return
	}

	prefix := c.Param("key")
	for key, value := range p {
		viper.Set(fmt.Sprintf("%s.%s", prefix, key), value)
	}
	if err := viper.WriteConfigAs("config.yml"); err != nil {
		ginutil.JSONServerError(c, err)
		return
	}

	ginutil.JSON(c)
}

// delete godoc
// @Tags v1/Configs
// @Summary 删除配置项
// @Description 根据键名删除配置项
// @Accept json
// @Produce json
// @Param key path string true "键名"
// @Success 200 {object} httputil.JSONResponse
// @Failure 400 {object} httputil.JSONResponse
// @Failure 500 {object} httputil.JSONResponse
// @Router /v1/configs/{key} [delete]
func (rs *ConfigResource) delete(c *gin.Context) {
	key := c.Param("key")
	if !viper.IsSet(key) {
		ginutil.JSONBadRequest(c, fmt.Errorf("key %s not exist", key))
		return
	}

	viper.Set(key, nil)
	if err := viper.SafeWriteConfig(); err != nil {
		ginutil.JSONServerError(c, err)
		return
	}

	ginutil.JSON(c)
}
