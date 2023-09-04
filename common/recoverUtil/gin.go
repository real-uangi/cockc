// Package recoverUtil
// @author uangi
// @date 2023/9/4 16:33
package recoverUtil

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/real-uangi/cockc/common/api"
	"net/http"
)

type GinRecover struct {
	serial int64
}

func Gin() *GinRecover {
	return &GinRecover{serial: 0}
}

func (g *GinRecover) Catch(c *gin.Context) {
	ev := recover()
	if ev != nil {
		logger.Error(fmt.Sprint(ev))
		c.JSON(http.StatusInternalServerError, api.Fail(ev.(error).Error(), nil))
	}
}
