package v1

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/saltbo/gopkg/ginutil"
	_ "github.com/saltbo/gopkg/httputil"

	"github.com/zplat-core/apiserver/api/bind"
	"github.com/zplat-core/apiserver/dao"
)

type SubsystemResource struct {
	dSubsystem *dao.Subsystem
}

func NewSubsystemResource() *SubsystemResource {
	return &SubsystemResource{
		dSubsystem: dao.NewSubsystem(),
	}
}

func (rs *SubsystemResource) Register(router *gin.RouterGroup) {
	router.GET("/subsystems", rs.findAll)
	router.POST("/subsystems", rs.create)
	router.PUT("/subsystems/:id", rs.update)
	router.DELETE("/subsystems/:id", rs.delete)
}

// findAll godoc
// @Tags v1/Subsystem
// @Summary 获取所有子系统
// @Description 获取所有子系统
// @Accept json
// @Produce json
// @Success 200 {object} httputil.JSONResponse{data=gin.H{list=[]model.Subsystem,total=int64}}
// @Failure 400 {object} httputil.JSONResponse
// @Failure 500 {object} httputil.JSONResponse
// @Router /v1/subsystems [get]
func (rs *SubsystemResource) findAll(c *gin.Context) {
	p := new(bind.QuerySubsystem)
	if err := c.BindQuery(p); err != nil {
		ginutil.JSONBadRequest(c, err)
		return
	}

	query := dao.NewQuery()
	query.WithPage(p.PageNo, p.PageSize)
	if p.Name != "" {
		query.WithEq("name", p.Name)
	}

	list, total, err := rs.dSubsystem.FindAll(query)
	if err != nil {
		ginutil.JSONServerError(c, err)
		return
	}

	ginutil.JSONList(c, list, total)
}

// create godoc
// @Tags v1/Subsystem
// @Summary 添加子系统
// @Description 添加子系统
// @Accept json
// @Produce json
// @Param body body bind.BodySubsystem true "参数"
// @Success 200 {object} httputil.JSONResponse
// @Failure 400 {object} httputil.JSONResponse
// @Failure 500 {object} httputil.JSONResponse
// @Router /v1/subsystems [post]
func (rs *SubsystemResource) create(c *gin.Context) {
	p := new(bind.BodySubsystem)
	if err := c.ShouldBindJSON(p); err != nil {
		ginutil.JSONBadRequest(c, err)
		return
	}

	data := p.Model()
	if err := rs.dSubsystem.Add(data); err != nil {
		ginutil.JSONBadRequest(c, err)
		return
	}

	ginutil.JSONData(c, data)
}

// update godoc
// @Tags v1/Subsystem
// @Summary 修改子系统
// @Description 修改子系统
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Param body body bind.BodySubsystem true "参数"
// @Success 200 {object} httputil.JSONResponse
// @Failure 400 {object} httputil.JSONResponse
// @Failure 500 {object} httputil.JSONResponse
// @Router /v1/subsystems/{id} [put]
func (rs *SubsystemResource) update(c *gin.Context) {
	p := new(bind.BodySubsystem)
	if err := c.ShouldBindJSON(p); err != nil {
		ginutil.JSONBadRequest(c, err)
		return
	}

	sid, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	if err := rs.dSubsystem.Update(sid, p.Model()); err != nil {
		ginutil.JSONBadRequest(c, err)
		return
	}

	ginutil.JSON(c)
}

// delete godoc
// @Tags v1/Subsystem
// @Summary 删除子系统
// @Description 删除子系统
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} httputil.JSONResponse
// @Failure 400 {object} httputil.JSONResponse
// @Failure 500 {object} httputil.JSONResponse
// @Router /v1/subsystems/{id} [delete]
func (rs *SubsystemResource) delete(c *gin.Context) {
	sid, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		ginutil.JSONBadRequest(c, fmt.Errorf("invalid id: %s", err))
		return
	}

	if err := rs.dSubsystem.Delete(sid); err != nil {
		ginutil.JSONBadRequest(c, err)
		return
	}

	ginutil.JSON(c)
}
