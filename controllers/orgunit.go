package controllers

import (
	// "encoding/json"
	// "errors"
	"strconv"
	// "strings"

	"github.com/astaxie/beego"
	"github.com/showntop/ttkeeper/models"
)

//  OrgunitController operations for Orgunit
type OrgunitController struct {
	beego.Controller
}

// URLMapping ...
func (c *OrgunitController) URLMapping() {
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Post", c.Post)
	// c.Mapping("GetOne", c.GetOne)
	// c.Mapping("Put", c.Put)
	// c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Orgunit
// @Param	body		body 	models.Orgunit	true		"body for Orgunit content"
// @Success 201 {int} models.Orgunit
// @Failure 403 body is empty
// @router / [post]
func (c *OrgunitController) Post() {
	c.Ctx.Request.ParseForm()

	var v models.Orgunit
	v.ParentID, _ = strconv.ParseInt(c.Ctx.Request.FormValue("parent_id"), 10, 64)
	v.Name = c.Ctx.Request.FormValue("name")

	if _, err := models.AddOrgunit(&v); err == nil {
		c.Ctx.Output.SetStatus(201)
		c.Data["json"] = v
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// // GetOne ...
// // @Title Get One
// // @Description get Orgunit by id
// // @Param	id		path 	string	true		"The key for staticblock"
// // @Success 200 {object} models.Orgunit
// // @Failure 403 :id is empty
// // @router /:id [get]
// func (c *OrgunitController) GetOne() {
// 	idStr := c.Ctx.Input.Param(":id")
// 	id, _ := strconv.ParseInt(idStr, 0, 64)
// 	v, err := models.GetOrgunitById(id)
// 	if err != nil {
// 		c.Data["json"] = err.Error()
// 	} else {
// 		c.Data["json"] = v
// 	}
// 	c.ServeJSON()
// }

// GetAll ...
// @Title Get All
// @Description get Orgunit
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Orgunit
// @Failure 403
// @router / [get]
func (c *OrgunitController) GetAll() {
	parentID, _ := strconv.ParseInt(c.Ctx.Request.URL.Query().Get("parent_id"), 10, 64)

	l, err := models.GetAllOrgunit(parentID)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// // Put ...
// // @Title Put
// // @Description update the Orgunit
// // @Param	id		path 	string	true		"The id you want to update"
// // @Param	body		body 	models.Orgunit	true		"body for Orgunit content"
// // @Success 200 {object} models.Orgunit
// // @Failure 403 :id is not int
// // @router /:id [put]
// func (c *OrgunitController) Put() {
// 	idStr := c.Ctx.Input.Param(":id")
// 	id, _ := strconv.ParseInt(idStr, 0, 64)
// 	v := models.Orgunit{Id: id}
// 	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
// 	if err := models.UpdateOrgunitById(&v); err == nil {
// 		c.Data["json"] = "OK"
// 	} else {
// 		c.Data["json"] = err.Error()
// 	}
// 	c.ServeJSON()
// }

// // Delete ...
// // @Title Delete
// // @Description delete the Orgunit
// // @Param	id		path 	string	true		"The id you want to delete"
// // @Success 200 {string} delete success!
// // @Failure 403 id is empty
// // @router /:id [delete]
// func (c *OrgunitController) Delete() {
// 	idStr := c.Ctx.Input.Param(":id")
// 	id, _ := strconv.ParseInt(idStr, 0, 64)
// 	if err := models.DeleteOrgunit(id); err == nil {
// 		c.Data["json"] = "OK"
// 	} else {
// 		c.Data["json"] = err.Error()
// 	}
// 	c.ServeJSON()
// }
