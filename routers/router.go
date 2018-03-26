package routers

import (
	"github.com/astaxie/beego"
	. "github.com/showntop/ttkeeper/controllers"
	"github.com/showntop/ttkeeper/controllers/pages"
)

func init() {
	beego.Router("/", &pages.HomeController{})
	beego.Router("/login", &pages.LoginController{})
	beego.Router("/home", &pages.HomeController{})
	beego.Router("/system", &pages.SystemController{})
	beego.Router("/orgunit", &pages.OrgunitController{})

	beego.Router("/api/sessions", &SessController{})
	beego.Router("/api/users", &UserController{})
	beego.Router("/api/u/:id/permissions", &UserController{}, "get:GetPermissions")
	beego.Router("/api/roles", &RoleController{})
	beego.Router("/api/resources", &ResourceController{})
	beego.Router("/api/permissions", &PermissionController{})
	beego.Router("/api/roles/:id/permissions", &RoleController{}, "post:Grant")

}
