package api // import "github.com/dayvillefire/tenders/api"

import (
	"net/http"

	"github.com/dayvillefire/tenders/common"
	"github.com/dayvillefire/tenders/models"
	"github.com/gin-gonic/gin"
)

func init() {
	common.ApiMap["event"] = func(r *gin.RouterGroup) {
		r.PUT("", apiEventSave)
		r.GET("/:id", apiEventGet)
		r.GET("/:id/commitments", apiEventGet)
		r.DELETE("/:id", apiEventDelete)
	}
	common.ApiMap["events"] = func(r *gin.RouterGroup) {
		r.GET("", apiEventList)
		r.GET("/range/:from/:to", apiEventList)
	}
}

func apiEventSave(c *gin.Context) {
	var obj models.Event
	if c.BindJSON(&obj) != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	_, err := common.DB.ValidateAndSave(&obj)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, obj.ID)
}

func apiEventDelete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	var obj models.Event
	err := common.DB.Find(&obj, id)
	if err != nil {
		c.AbortWithError(http.StatusGone, err)
		return
	}
	err = common.DB.Destroy(&obj)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, true)
}

func apiEventGet(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	var obj models.Event
	err := common.DB.Find(&obj, id)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}
	c.JSON(http.StatusOK, obj)
}

func apiEventList(c *gin.Context) {
	fromDtRaw := c.Param("from")
	toDtRaw := c.Param("to")
	fromDt, _ := toDate(fromDtRaw)
	toDt, _ := toDate(toDtRaw)

	pageNumber := c.Query("page")
	perPage := c.Query("perpage")
	var obj models.Events
	var err error
	if fromDtRaw == "" || toDtRaw == "" {
		err = common.DB.All(&obj)
	} else {
		query := common.DB.Where("dateof BETWEEN ? AND ?", fromDt, toDt)
		if pageNumber != "" && perPage != "" {
			query = query.Paginate(MustQueryInt(c, "page"), MustQueryInt(c, "perpage"))
		}
		err = query.All(&obj)
	}
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}
	c.JSON(http.StatusOK, obj)
}

func apiEventCommitments(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	var evt models.Event
	err := common.DB.Find(&evt, id)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	var obj models.Commitments
	query := common.DB.Where("events.id = ?", id).LeftJoin("events", "events.id = commitments.event")
	err = query.All(&obj)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, obj)
}
