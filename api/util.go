package api

import (
	"log"
	"strconv"
	"time"

	"github.com/dayvillefire/tenders/common"
	"github.com/dayvillefire/tenders/config"
	"github.com/dayvillefire/tenders/models"
	"github.com/gin-gonic/gin"
)

func contextRequestHeader(c *gin.Context, key string) string {
	if config.Config.Debug {
		log.Printf("contextRequestHeader: %#v", c.Request.Header[key])
	}
	if values, _ := c.Request.Header[key]; len(values) > 0 {
		return values[0]
	}
	return ""
}

// MustQueryInt produces an integer from a gin Context query variable, and
// if one isn't present (or it isn't a number), it produces a zero
func MustQueryInt(c *gin.Context, key string) int {
	q := c.Query(key)
	if q == "" {
		return 0
	}
	i, _ := strconv.Atoi(key)
	return i
}

func toDate(in string) (time.Time, error) {
	return time.Parse("2006-01-02", in)
}

func mustUserLookup(username string) (models.User, error) {
	user := models.User{}
	query := common.DB.Where("name = ?", username).Limit(1)
	err := query.All(&user)
	return user, err
}
