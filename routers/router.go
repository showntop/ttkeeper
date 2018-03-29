package routers

import (
	"github.com/gin-gonic/gin"
	. "github.com/showntop/ttkeeper/controllers"
	// "github.com/showntop/ttkeeper/controllers/pages"
)

func init() {

	router := gin.Default()
	router.Use(gin.Logger())

	router.GET("/", func(ctx *gin.Context) {
	})
	// router.GET("/login", &pages.LoginController{})
	// router.GET("/home", &pages.HomeController{})
	// router.GET("/system", &pages.SystemController{})
	// router.GET("/orgunit", &pages.OrgunitController{})

	// beego.Router("/api/sessions", &SessController{})
	// router.Use(Parse)

	router.POST("/v1/ss", func(ctx *gin.Context) {
		ssc := new(SessController)
		ssc.Ctx = ctx
		ssc.Post()
	})

	router.Use(Authenticate)
	router.DELETE("/v1/ss", func(ctx *gin.Context) {
		ssc := SessController{}
		ssc.Ctx = ctx
		ssc.Delete()
	})

	v1 := router.Group("/v1")
	v1.Use(Permit)
	{
		v1.POST("/u", UserC.Post)
		v1.GET("/u", UserC.GetAll)

		v1.POST("/r", RoleC.Post)
		v1.GET("/r", RoleC.GetAll)

		v1.POST("/og", OrgC.Post)
		v1.GET("/og", OrgC.GetAll)

		v1.POST("/rs", ResC.Post)
		v1.GET("/rs", ResC.GetAll)

		v1.GET("/u:user_id/p", func(ctx *gin.Context) {

		})

		v1.POST("/p", func(ctx *gin.Context) {

		})
		v1.GET("/p", func(ctx *gin.Context) {

		})

	}

	router.Run()
}
