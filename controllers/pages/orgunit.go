package pages

import (
	"github.com/astaxie/beego"
)

type OrgunitController struct {
	beego.Controller
}

func (c *OrgunitController) Get() {
	//auth

	// get url of the id number.
	url := "hauth/org_page.tpl"
	c.TplName = url
}
