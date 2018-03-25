package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/showntop/ttkeeper/models"
)

//  ResourceController operations for Resource
type ResourceController struct {
	beego.Controller
}

// URLMapping ...
func (c *ResourceController) URLMapping() {
	c.Mapping("Post", c.Post)
}

// Post ...
// @Title Post
// @Description create Resource
// @Param	body		body 	models.Resource	true		"body for Resource content"
// @Success 201 {int} models.Resource
// @Failure 403 body is empty
// @router / [post]
func (c *ResourceController) Post() {
	var v models.Resource
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if _, err := models.AddResource(&v); err == nil {
		c.Ctx.Output.SetStatus(201)
		c.Data["json"] = v
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
