package controllers

import (
	"encoding/json"
	// "errors"
	"strconv"
	// "strings"

	"github.com/gin-gonic/gin"
	"github.com/showntop/ttkeeper/models"
)

//  ResourceController operations for Resource
type ResourceController struct {
	Handler
}

var (
	ResC *ResourceController = new(ResourceController)
)

func (c *ResourceController) Post(ctx *gin.Context) {
	var v models.Resource
	dc := json.NewDecoder(ctx.Request.Body)
	if err := dc.Decode(&v); err != nil {
		ctx.AbortWithStatusJSON(200, makeResult(CODE_PARAM_ERROR, err.Error(), nil))
		return
	}
	if len(v.Name) <= 0 || v.ParentID == 0 || v.Rtype <= models.RESOURCE_TYPE_INVA || v.Code == "" {
		ctx.AbortWithStatusJSON(200, makeResult(CODE_PARAM_ERROR, "name, parent_id type code cannot empty.", nil))
		return
	}
	m, err := models.AddResource(&v)
	if err != nil {
		ctx.JSON(200, makeResult(CODE_DB_ERROR, err.Error(), nil))
		return
	}

	ctx.JSON(201, makeResult(CODE_OK, "ok", m))
}

func (c *ResourceController) GetAll(ctx *gin.Context) {

	var limit int64 = 10
	var offset int64
	var parentID int64
	var rtype int64
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

	if o, ok := ctx.GetQuery("parent_id"); ok {
		if v, err := strconv.ParseInt(o, 10, 64); err == nil {
			parentID = v
		}
	}

	if o, ok := ctx.GetQuery("type"); ok {
		if v, err := strconv.ParseInt(o, 10, 64); err == nil {
			rtype = v
		}
	}

	l, err := models.GetAllResource(parentID, rtype, offset, limit)
	if err != nil {
		ctx.AbortWithStatusJSON(200, makeResult(CODE_DB_ERROR, err.Error(), nil))
	}
	ctx.JSON(200, makeResult(CODE_OK, "ok", l))
}
