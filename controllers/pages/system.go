package pages

import (
	"html/template"

	"github.com/astaxie/beego"
	"github.com/showntop/weapon/jwt"
)

type SystemController struct {
	beego.Controller
}

func (c *SystemController) Get() {

	// get user connection information from cookie.
	cookie, _ := c.Ctx.Request.Cookie("Authorization")
	jclaim, err := jwt.ParseJwt(cookie.Value)
	if err != nil {
		return
	}

	// get url of the id number.
	url := "./views/hauth/theme/default/sysconfig.tpl"
	h, err := template.ParseFiles(url)
	if err != nil {
		return
	}
	h.Execute(c.Ctx.ResponseWriter, jclaim.UserId)

	// key := sha1.GenSha1Key(id, jclaim.UserId, url)

	// if !groupcache.FileIsExist(key) {
	// 	groupcache.RegisterStaticFile(key, url)
	// }

	// tpl, err := groupcache.GetStaticFile(key)
	// if err != nil {
	// 	logs.Error(err)
	// 	hret.Error(c.Ctx.ResponseWriter, 404, i18n.PageNotFound(c.Ctx.Request))
	// 	return
	// }
	// c.Ctx.ResponseWriter.Write(tpl)
	// c.Data["Website"] = "beego.me"
	// c.Data["Email"] = "astaxie@gmail.com"
	// c.TplName = "index.tpl"
}
