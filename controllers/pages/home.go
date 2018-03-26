package pages

import (
	"github.com/astaxie/beego"
	"github.com/showntop/weapon/jwt"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Get() {

	var err error

	cok, err := c.Ctx.Request.Cookie("Authorization")
	if err != nil {
		beego.Warn(err)
		c.Ctx.Redirect(302, "/login")
		return
	}
	// beego.Info(cok)
	jclaim, err := jwt.ParseJwt(cok.Value)
	if err != nil {
		c.Ctx.Redirect(302, "/login")
		return
	}

	var url = "hauth/theme/default/index.tpl"

	c.Data["username"] = jclaim.UserId
	c.TplName = url
}
