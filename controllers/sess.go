package controllers

import (
	"fmt"
	"net/http"

	"github.com/showntop/ttkeeper/models"
	"github.com/showntop/weapon/crypto/haes"
	"github.com/showntop/weapon/jwt"
)

// SessController operations for Sess
type SessController struct {
	Handler
}

func (c *SessController) Post() {

	var username, password string
	params := c.Ctx.GetStringMap("params")
	if u, ok := params["username"].(string); ok {
		username = u
	} else {
		c.Ctx.JSON(http.StatusOK, makeResult(CODE_AUTH_ERROR, "username validate error", nil))
		return
	}

	if u, ok := params["password"].(string); ok {
		password = u
	} else {
		c.Ctx.JSON(http.StatusOK, makeResult(CODE_AUTH_ERROR, "password validate error", nil))
		return
	}

	psd, err := haes.Encrypt(password)
	if err != nil {
		c.Ctx.JSON(http.StatusOK, makeResult(CODE_AUTH_ERROR, err.Error(), nil))
		return
	}

	user, err := models.GetUserByUsername(username)
	if err != nil {
		c.Ctx.JSON(http.StatusOK, makeResult(CODE_AUTH_ERROR, err.Error(), nil))
		return
	}
	fmt.Printf("user: %+v\r\n", user)
	if psd == user.Password {
		token, _ := jwt.GenToken(fmt.Sprintf("%d", user.ID), fmt.Sprintf("%d", 1), fmt.Sprintf("%d", user.OrgunitID), 86400)
		cookie := http.Cookie{Name: "Authorization", Value: token, Path: "/", MaxAge: 86400}
		http.SetCookie(c.Ctx.Writer, &cookie)
		c.Ctx.JSON(http.StatusOK, makeResult(CODE_OK, "ok", map[string]interface{}{"token": token, "expire": 86400}))
	} else {
		c.Ctx.JSON(http.StatusOK, makeResult(CODE_AUTH_ERROR, "password wrong.", nil))
	}
}

func (c *SessController) Delete() {
	c.Ctx.JSON(http.StatusOK, makeResult(CODE_OK, "user has logouted.", nil))
}
