package controllers

import (
	"encoding/json"
	// "errors"
	"strconv"
	// "strings"

	"github.com/gin-gonic/gin"
	"github.com/showntop/ttkeeper/models"
)

//  RoleController operations for Role
type RoleController struct {
	Handler
}

var (
	RoleC *RoleController = new(RoleController)
)

func (c *RoleController) GetAll(ctx *gin.Context) {
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

	l, err := models.GetAllRole(offset, limit)
	if err != nil {
		ctx.AbortWithStatusJSON(200, makeResult(CODE_DB_ERROR, err.Error(), nil))
	}
	ctx.JSON(200, makeResult(CODE_OK, "ok", l))
}

func (c *RoleController) Post(ctx *gin.Context) {
	var v models.Role
	dc := json.NewDecoder(ctx.Request.Body)
	if err := dc.Decode(&v); err != nil {
		ctx.AbortWithStatusJSON(200, makeResult(CODE_PARAM_ERROR, err.Error(), nil))
		return
	}
	if len(v.Name) <= 0 {
		ctx.AbortWithStatusJSON(200, makeResult(CODE_PARAM_ERROR, "name, cannot empty.", nil))
		return
	}
	if user, err := models.AddRole(&v); err == nil {
		ctx.JSON(201, makeResult(CODE_OK, "ok", user))
	} else {
		ctx.JSON(200, makeResult(CODE_DB_ERROR, "db error", nil))
	}
}

func (c *RoleController) Grant() {
	// var v models.Permission
	// json.Unmarshal(, &v)

	// roleID, _ := strconv.ParseInt(c.Ctx.Input.Param(":id"), 10, 64)
	// role := models.Role{ID: roleID}
	// if err := role.Grant(&v); err == nil {
	// 	c.Ctx.Output.SetStatus(201)
	// 	c.Data["json"] = v
	// } else {
	// 	c.Data["json"] = err.Error()
	// }
	// c.ServeJSON()
}
