package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["username"] = "astaxie@gmail.com"
	c.TplName = "hauth/theme/default/index.tpl"
}

type ErrorResult struct {
	Version   string `json:"version"`
	ErrorCode int    `json:"error_code"`
	ErrorMsg  string `json:"error_msg"`
}

func newErrorResult(code int, msg string) string {
	b, _ := json.Marshal(&ErrorResult{Version: "v1.0", ErrorCode: code, ErrorMsg: msg})
	return string(b)
}
