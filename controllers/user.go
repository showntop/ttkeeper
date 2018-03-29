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

func (c *UserController) GetPermissions(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, _ := strconv.ParseInt(idStr, 0, 64)

	containerIdStr := ctx.Request.URL.Query().Get("container_id")
	containerId, _ := strconv.ParseInt(containerIdStr, 0, 64)
	v, err := models.GetUserPermissions(id, containerId)
	if err != nil {
		ctx.JSON(200, makeResult(CODE_DB_ERROR, err.Error(), nil))
		return
	}
	ctx.JSON(200, makeResult(CODE_OK, "ok", v))
}

// // GetOne ...
// // @Title Get One
// // @Description get User by id
// // @Param	id		path 	string	true		"The key for staticblock"
// // @Success 200 {object} models.User
// // @Failure 403 :id is empty
// // @router /:id [get]
// func (c *UserController) GetOne() {
// 	idStr := c.Ctx.Input.Param(":id")
// 	id, _ := strconv.ParseInt(idStr, 0, 64)
// 	v, err := models.GetUserById(id)
// 	if err != nil {
// 		c.Data["json"] = err.Error()
// 	} else {
// 		c.Data["json"] = v
// 	}
// 	c.ServeJSON()
// }

// // Put ...
// // @Title Put
// // @Description update the User
// // @Param	id		path 	string	true		"The id you want to update"
// // @Param	body		body 	models.User	true		"body for User content"
// // @Success 200 {object} models.User
// // @Failure 403 :id is not int
// // @router /:id [put]
// func (c *UserController) Put() {
// 	idStr := c.Ctx.Input.Param(":id")
// 	id, _ := strconv.ParseInt(idStr, 0, 64)
// 	v := models.User{Id: id}
// 	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
// 	if err := models.UpdateUserById(&v); err == nil {
// 		c.Data["json"] = "OK"
// 	} else {
// 		c.Data["json"] = err.Error()
// 	}
// 	c.ServeJSON()
// }

// // Delete ...
// // @Title Delete
// // @Description delete the User
// // @Param	id		path 	string	true		"The id you want to delete"
// // @Success 200 {string} delete success!
// // @Failure 403 id is empty
// // @router /:id [delete]
// func (c *UserController) Delete() {
// 	idStr := c.Ctx.Input.Param(":id")
// 	id, _ := strconv.ParseInt(idStr, 0, 64)
// 	if err := models.DeleteUser(id); err == nil {
// 		c.Data["json"] = "OK"
// 	} else {
// 		c.Data["json"] = err.Error()
// 	}
// 	c.ServeJSON()
// }
