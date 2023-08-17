// Package api
// @author uangi
// @date 2023/8/16 9:33
package api

import (
	"github.com/gin-gonic/gin"
	"github.com/real-uangi/cockc/constants"
	"net/http"
)

type UserLevel int

const (
	Admin UserLevel = 0
	User  UserLevel = 1
)

type AuthInfo struct {
	AuthToken string    `json:"authToken"`
	UserName  string    `json:"userName"`
	UserLevel UserLevel `json:"userLevel"`
	UserId    int64     `json:"userId"`
	Account   string    `json:"account"`
	Phone     string    `json:"phone"`
}

func GetUserInfo(c *gin.Context) AuthInfo {
	v, b := c.Get(constants.AuthInfoContext)
	if !b {
		c.JSON(http.StatusUnauthorized, UnAuthorized("410 Unauthorized"))
	}
	return v.(AuthInfo)
}

func IsAuthed(c *gin.Context) bool {
	return c.GetBool(constants.AuthResultContext)
}
