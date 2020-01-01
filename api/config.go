package api // import "github.com/dayvillefire/tenders/api"

import (
	"log"
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

func contextRequestHeader(c *gin.Context, key string) string {
	if config.Config.Debug {
		log.Printf("contextRequestHeader: %#v", c.Request.Header[key])
	}
	if values, _ := c.Request.Header[key]; len(values) > 0 {
		return values[0]
	}
	return ""
}
