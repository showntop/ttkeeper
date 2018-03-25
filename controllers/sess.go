package controllers

import (
	"fmt"
	"net/http"

	"github.com/astaxie/beego"
	"github.com/showntop/ttkeeper/models"
	"github.com/showntop/weapon/crypto/haes"
	"github.com/showntop/weapon/hret"
	"github.com/showntop/weapon/jwt"
)

// SessController operations for Sess
type SessController struct {
	beego.Controller
}

// URLMapping ...
func (c *SessController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Sess
// @Param	body		body 	models.Sess	true		"body for Sess content"
// @Success 201 {object} models.Sess
// @Failure 403 body is empty
// @router / [post]
func (c *SessController) Post() {
	c.Ctx.Request.ParseForm()

	userName := c.Ctx.Request.FormValue("username")
	userPasswd := c.Ctx.Request.FormValue("password")
	psd, err := haes.Encrypt(userPasswd)
	if err != nil {
		hret.Error(c.Ctx.ResponseWriter, 300, err.Error(), 10)
		return
	}

	user, err := models.GetUserByUsername(userName)
	if err != nil {
		hret.Error(c.Ctx.ResponseWriter, 300, err.Error(), 10)
		return
	}

	if psd == user.Password {
		token, _ := jwt.GenToken(fmt.Sprintf("%d", user.ID), fmt.Sprintf("%s", user.OrgunitID), "", 86400)
		cookie := http.Cookie{Name: "Authorization", Value: token, Path: "/", MaxAge: 86400}
		http.SetCookie(c.Ctx.ResponseWriter, &cookie)
		hret.Success(c.Ctx.ResponseWriter, nil)
	} else {
		hret.Error(c.Ctx.ResponseWriter, 300, "rmsg", 10)
	}
}

// Delete ...
// @Title Delete
// @Description delete the Sess
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *SessController) Delete() {

}
