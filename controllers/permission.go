package controllers

import (
	"encoding/json"
	// "errors"
	"strconv"
	// "strings"

	"github.com/gin-gonic/gin"
	"github.com/showntop/ttkeeper/models"
)

//  PermissionController operations for Permission
type PermissionController struct {
	Handler
}

var (
	PerC *PermissionController = new(PermissionController)
)

func (c *PermissionController) Post(ctx *gin.Context) {
	var v models.Permission
	dc := json.NewDecoder(ctx.Request.Body)
	if err := dc.Decode(&v); err != nil {
		ctx.AbortWithStatusJSON(200, makeResult(CODE_PARAM_ERROR, err.Error(), nil))
		return
	}
	if v.Action <= 0 || v.ResourceId <= 0 || v.RoleID <= 0 {
		ctx.AbortWithStatusJSON(200, makeResult(CODE_PARAM_ERROR, "role_id, action resource_id cannot empty.", nil))
		return
	}
	m, err := models.AddPermission(&v)
	if err != nil {
		ctx.JSON(200, makeResult(CODE_DB_ERROR, err.Error(), nil))
		return
	}

	ctx.JSON(201, makeResult(CODE_OK, "ok", m))
}
func (c *PermissionController) GetAll(ctx *gin.Context) {
	var limit int64 = 10
	var offset int64
	var roleID int64 = -1
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

	if o, ok := ctx.GetQuery("role_id"); ok {
		if v, err := strconv.ParseInt(o, 10, 64); err == nil {
			roleID = v
		}
	}

	l, err := models.GetAllPermission(roleID, offset, limit)
	if err != nil {
		ctx.AbortWithStatusJSON(200, makeResult(CODE_DB_ERROR, err.Error(), nil))
	}
	ctx.JSON(200, makeResult(CODE_OK, "ok", l))
}
