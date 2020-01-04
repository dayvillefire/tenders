package api // import "github.com/dayvillefire/tenders/api"

import (
	"net/http"

	"github.com/dayvillefire/tenders/common"
	"github.com/dayvillefire/tenders/config"
	"github.com/gin-gonic/gin"
)

func init() {
	common.ApiMap["config"] = func(r *gin.RouterGroup) {
		r.GET("/", apiUIConfig)
	}
}

func apiUIConfig(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"debug": config.Config.Debug,
	})
}
