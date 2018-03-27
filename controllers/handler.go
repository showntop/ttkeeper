package controllers

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/showntop/weapon/jwt"
)

const (
	CODE_OK         = 10000
	CODE_AUTH_ERROR = 20000
	CODE_DB_ERROR   = 20000
)

type result struct {
	Retcode int         `json:"retcode"`
	Retmsg  string      `json:"retmsg"`
	Data    interface{} `json:"data"`
}

func makeResult(code int, msg string, data interface{}) *result {
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

var Parse = func(ctx *gin.Context) {
	var v map[string]interface{}
	dc := json.NewDecoder(ctx.Request.Body)
	err := dc.Decode(&v)
	if err != nil {
		ctx.AbortWithStatusJSON(200, result{20000, "request json format wrong.", nil})
		return
	}
	ctx.Set("params", v)
}

var Authenticate = func(ctx *gin.Context) {
	var err error
	cok, err := ctx.Request.Cookie("Authorization")
	if err != nil {
		ctx.AbortWithStatusJSON(200, result{20000, err.Error(), nil})
		return
	}

	_, err = jwt.ParseJwt(cok.Value)
	if err != nil {
		ctx.AbortWithStatusJSON(200, result{20000, err.Error(), nil})
		return
	}
	ctx.Set("user", "value")
}

var Permit = func(ctx *gin.Context) {
	var err error
	cok, err := ctx.Request.Cookie("Authorization")
	if err != nil {
		ctx.AbortWithStatusJSON(200, result{20000, err.Error(), nil})
		return
	}

	_, err = jwt.ParseJwt(cok.Value)
	if err != nil {
		ctx.AbortWithStatusJSON(200, result{20000, err.Error(), nil})
		return
	}
}
