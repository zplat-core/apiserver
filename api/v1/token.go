package v1

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/saltbo/gopkg/ginutil"
	_ "github.com/saltbo/gopkg/httputil"

	"github.com/zplat-core/apiserver/api/middleware"
	"github.com/zplat-core/apiserver/service/system"

	//"github.com/zplat-core/apiserver/internel/pkg/system"

	"github.com/zplat-core/apiserver/api/bind"
	"github.com/zplat-core/apiserver/service"
)

type TokenResource struct {
}

func NewTokenResource() *TokenResource {
	return &TokenResource{}
}

func (rs *TokenResource) Register(router *gin.RouterGroup) {
	router.POST("/tokens", rs.create)
	router.DELETE("/tokens", rs.delete)
}

// create godoc
// @Tags v1/Tokens
// @Summary 登录/密码重置
// @Description 用于账户登录和申请密码重置
// @Accept json
// @Produce json
// @Param body body bind.BodyToken true "参数"
// @Success 200 {object} httputil.JSONResponse
// @Failure 400 {object} httputil.JSONResponse
// @Failure 500 {object} httputil.JSONResponse
// @Router /v1/tokens [post]
func (rs *TokenResource) create(c *gin.Context) {
	p := new(bind.BodyToken)
	if err := c.ShouldBindJSON(p); err != nil {
		ginutil.JSONBadRequest(c, err)
		return
	}

	// issue a signIn token into cookies
	if p.Password != "" {
		user, err := service.UserSignIn(p.Email, p.Password)
		if err != nil {
			ginutil.JSONBadRequest(c, err)
			return
		} else if system.EmailActed() && !user.Activated() {
			ginutil.JSONBadRequest(c, fmt.Errorf("account is not activated"))
			return
		}

		expireSec := 7 * 24 * 3600
		token, err := service.TokenCreate(user.Ux, expireSec, user.RolesSplit()...)
		if err != nil {
			ginutil.JSONServerError(c, err)
			return
		}

		middleware.TokenCookieSet(c, token, expireSec)
		middleware.RoleCookieSet(c, user.Roles, expireSec)
		ginutil.JSON(c)
		return
	}

	user, ok := service.UserEmailExist(p.Email)
	if !ok {
		ginutil.JSONBadRequest(c, fmt.Errorf("email not exist"))
		return
	}

	// issue a short-term token for password reset
	token, err := service.TokenCreate(user.Ux, 300)
	if err != nil {
		ginutil.JSONServerError(c, err)
		return
	}

	resetLink := service.PasswordRestLink(ginutil.GetOrigin(c), p.Email, token)
	if err := service.PasswordResetNotify(p.Email, resetLink); err != nil {
		ginutil.JSONServerError(c, err)
		return
	}

	ginutil.JSON(c)
}

// delete godoc
// @Tags v1/Tokens
// @Summary 退出登录
// @Description 用户状态登出
// @Accept json
// @Produce json
// @Success 200 {object} httputil.JSONResponse
// @Failure 400 {object} httputil.JSONResponse
// @Failure 500 {object} httputil.JSONResponse
// @Router /v1/tokens [delete]
func (rs *TokenResource) delete(c *gin.Context) {
	middleware.TokenCookieSet(c, "", 1)
	middleware.RoleCookieSet(c, "", 1)
	return
}
