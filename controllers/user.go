package controllers

import (
	"encoding/json"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/showntop/ttkeeper/models"
	"github.com/showntop/weapon/crypto/haes"
)

type UserController struct {
	Handler
}

var (
	UserC *UserController = new(UserController)
)

func (c *UserController) GetAll(ctx *gin.Context) {
	var limit int64 = 10
	var offset int64
	// limit: 10 (default is 10)
	if l, ok := ctx.GetQuery("limit"); ok {
		if v, err := strconv.ParseInt(l, 10, 64); err == nil {
			limit = v
		}
	}
	// offset: 0 (default is 0)
	if o, ok := ctx.GetQuery("offset"); ok {
		if v, err := strconv.ParseInt(o, 10, 64); err == nil {
			offset = v
		}
	}

	l, err := models.GetAllUser(offset, limit)
	if err != nil {
		ctx.AbortWithStatusJSON(200, makeResult(CODE_DB_ERROR, err.Error(), nil))
	}
	ctx.JSON(200, makeResult(CODE_OK, "ok", l))
}

func (c *UserController) Post(ctx *gin.Context) {
	var v models.User
	dc := json.NewDecoder(ctx.Request.Body)
	if err := dc.Decode(&v); err != nil {
		ctx.AbortWithStatusJSON(200, makeResult(CODE_PARAM_ERROR, err.Error(), nil))
		return
	}
	if len(v.Username) <= 0 || len(v.Password) <= 0 || v.OrgunitID == 0 {
		ctx.AbortWithStatusJSON(200, makeResult(CODE_PARAM_ERROR, "username,password,orgunitid cannot empty.", nil))
		return
	}
	v.Password, _ = haes.Encrypt(v.Password)
	if user, err := models.AddUser(&v); err == nil {
		ctx.JSON(201, makeResult(CODE_OK, "ok", user))
	} else {
		ctx.JSON(200, makeResult(CODE_DB_ERROR, "db error", nil))
	}
}

func (c *UserController) Grant(ctx *gin.Context) {
	var v models.UserRole
	dc := json.NewDecoder(ctx.Request.Body)
	if err := dc.Decode(&v); err != nil {
		ctx.AbortWithStatusJSON(200, makeResult(CODE_PARAM_ERROR, err.Error(), nil))
		return
	}
	if v.RoleID <= 0 || v.UserID <= 0 {
		ctx.AbortWithStatusJSON(200, makeResult(CODE_PARAM_ERROR, " cannot empty.", nil))
		return
	}
	if model, err := models.AddUserRole(&v); err == nil {
		ctx.JSON(201, makeResult(CODE_OK, "ok", model))
	} else {
		ctx.JSON(200, makeResult(CODE_DB_ERROR, "db error", nil))
	}
}

func (c *UserController) GetAllRoles(ctx *gin.Context) {
	idStr, _ := ctx.GetQuery("user_id")
	id, _ := strconv.ParseInt(idStr, 0, 64)

	v, err := models.GetAllUserRole(id, 0, 0)
	if err != nil {
		ctx.JSON(200, makeResult(CODE_DB_ERROR, err.Error(), nil))
		return
	}
	ctx.JSON(200, makeResult(CODE_OK, "ok", v))
}

func (c *UserController) GetAllPermissions(ctx *gin.Context) {
	id := ctx.GetInt64("USER_ID")
	var containerId int64 = -1
	if ci, ok := ctx.GetQuery("container_id"); ok {
		containerId, _ = strconv.ParseInt(ci, 10, 64)
	}

	v, err := models.GetUserPermissions(id, containerId)
	if err != nil {
		ctx.JSON(200, makeResult(CODE_DB_ERROR, err.Error(), nil))
		return
	}
	ctx.JSON(200, makeResult(CODE_OK, "ok", v))
}
