package api // import "github.com/dayvillefire/tenders/api"

import (
	"net/http"

	"github.com/dayvillefire/tenders/common"
	"github.com/gin-gonic/gin"
)

func init() {
	common.ApiMap["user"] = func(r *gin.RouterGroup) {
		r.GET("/info", apiUserInfo)
	}
}

func apiUserInfo(c *gin.Context) {
	user := c.MustGet(gin.AuthUserKey).(string)

	userObj, err := mustUserLookup(user)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"first_name": userObj.FirstName,
		"last_name":  userObj.LastName,
		"email":      userObj.EmailAddress,
		"shortcode":  userObj.ShortCode,
		"avatar":     userObj.PictureURL,
	})
}
