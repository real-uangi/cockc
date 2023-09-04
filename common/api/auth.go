// Package api
// @author uangi
// @date 2023/8/16 9:33
package api

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/real-uangi/cockc/common/plog"
	"github.com/real-uangi/cockc/common/rdb"
	"github.com/real-uangi/cockc/common/recoverUtil"
	"github.com/real-uangi/cockc/common/response"
	"github.com/real-uangi/cockc/constants"
	"net/http"
	"strings"
)

var logger = plog.New("api")

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
		c.JSON(http.StatusUnauthorized, response.UnAuthorized("410 Unauthorized"))
	}
	return v.(AuthInfo)
}

func IsAuthed(c *gin.Context) bool {
	return c.GetBool(constants.AuthResultContext)
}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer recoverUtil.Gin().Catch(c)
		if isNoAuth(c.FullPath()) {
			return
		}
		token := c.GetHeader(constants.AuthHeader)
		if len(token) == 0 {
			c.JSON(http.StatusUnauthorized, response.UnAuthorized("401 Unauthorized"))
			c.Abort()
			return
		}
		bs, err := rdb.GetClient().Get(context.Background(), formatKey(token)).Bytes()
		logger.Panic(err)
		var info AuthInfo
		err = json.Unmarshal(bs, &info)
		logger.Panic(err)
		c.Set(constants.AuthInfoContext, info)
		c.Set(constants.AuthResultContext, true)
	}
}

func isNoAuth(uri string) bool {
	return strings.Contains(uri, constants.NoTokenKeyword)
}

func formatKey(token string) string {
	return constants.AuthRedisKey + token
}
