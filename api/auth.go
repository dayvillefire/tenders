package api

import (
	"github.com/dayvillefire/tenders/auth"
	"github.com/dayvillefire/tenders/common"
	"github.com/gin-gonic/gin"
)

func init() {
	common.ApiMap["auth"] = func(r *gin.RouterGroup) {
		r.GET("/refresh_token", auth.GetAuthMiddleware().RefreshHandler)
	}
}
