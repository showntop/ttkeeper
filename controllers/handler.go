package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/showntop/ttkeeper/models"
	"github.com/showntop/weapon/jwt"
)

const (
	CODE_OK          = 10000
	CODE_PARAM_ERROR = 20001
	CODE_AUTH_ERROR  = 20002
	CODE_DB_ERROR    = 30001
)

type result struct {
	Retcode int         `json:"retcode"`
	Retmsg  string      `json:"retmsg"`
	Data    interface{} `json:"data"`
}

func makeResult(code int, msg string, data interface{}) *result {
	if strings.Contains(msg, "Duplicate entry") {
		msg = "this entry has existed."
	}
	return &result{code, msg, data}
}

func (r *result) ToJson() []byte {
	b, err := json.Marshal(r)
	if err != nil {
		return []byte(`{retcode: 20000, retmsg:` + err.Error() + `}`)
	}
	return b
}

type Handler struct {
	Ctx *gin.Context
}

// var Parse = func(ctx *gin.Context) {
// 	var v map[string]interface{}
// 	dc := json.NewDecoder(ctx.Request.Body)
// 	err := dc.Decode(&v)
// 	if err != nil {
// 		ctx.AbortWithStatusJSON(200, result{20000, "request json format wrong.", nil})
// 		return
// 	}
// 	ctx.Set("params", v)
// }

var Authenticate = func(ctx *gin.Context) {
	var err error
	cok, err := ctx.Request.Cookie("Authorization")
	if err != nil {
		ctx.AbortWithStatusJSON(200, result{20000, err.Error(), nil})
		return
	}
	jclaim, err := jwt.ParseJwt(cok.Value)
	if err != nil {
		ctx.AbortWithStatusJSON(200, result{20000, err.Error(), nil})
		return
	}
	userID, _ := strconv.ParseInt(jclaim.UserId, 10, 64)
	ctx.Set("USER_ID", userID)
}

var Permit = func(ctx *gin.Context) {

	var userID int64
	if id, ok := ctx.Get("USER_ID"); ok {
		userID, _ = id.(int64)
	} else {
		ctx.AbortWithStatusJSON(200, result{20000, "has no permissions", nil})
	}

	pcode := fmt.Sprintf("%s %s", ctx.Request.Method, ctx.Request.URL.EscapedPath())
	permited := models.GrantedPermission(userID, models.RESOURCE_TYPE_API, pcode)
	fmt.Println("permited:", permited)
	if !permited {
		ctx.AbortWithStatusJSON(200, result{20000, "has no permissions", nil})
	}
}
