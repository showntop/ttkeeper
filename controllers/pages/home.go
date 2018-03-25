package pages

import (
	"html/template"

	"github.com/astaxie/beego"
	"github.com/showntop/weapon/jwt"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Get() {

	cok, _ := c.Ctx.Request.Cookie("Authorization")
	jclaim, err := jwt.ParseJwt(cok.Value)
	if err != nil {
		c.Ctx.Redirect(302, "/")
		return
	}

	var url = "./views/hauth/theme/default/index.tpl"

	h, err := template.ParseFiles(url)
	if err != nil {
		return
	}
	h.Execute(c.Ctx.ResponseWriter, jclaim.UserId)

	// c.Data["Website"] = "beego.me"
	// c.Data["Email"] = "astaxie@gmail.com"
	// c.TplName = "index.tpl"
}
