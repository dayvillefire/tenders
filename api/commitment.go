package api // import "github.com/dayvillefire/tenders/api"

import (
	"net/http"

	"github.com/dayvillefire/tenders/common"
	"github.com/dayvillefire/tenders/models"
	"github.com/gin-gonic/gin"
)

func init() {
	common.ApiMap["commitment"] = func(r *gin.RouterGroup) {
		r.PUT("", apiCommitmentSave)
		r.GET("/:id", apiCommitmentGet)
		r.DELETE("/:id", apiCommitmentDelete)
	}
	common.ApiMap["commitments"] = func(r *gin.RouterGroup) {
		r.GET("", apiCommitmentList)
		r.GET("/:query", apiCommitmentList)
	}
}

func apiCommitmentSave(c *gin.Context) {
	user := c.MustGet(gin.AuthUserKey).(string)

	userObj, err := mustUserLookup(user)

	var obj models.Commitment
	if c.BindJSON(&obj) != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	obj.User = userObj

	_, err = common.DB.ValidateAndSave(&obj)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, obj.ID)
}

func apiCommitmentDelete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	var obj models.Commitment
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

func apiCommitmentGet(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	var obj models.Commitment
	err := common.DB.Find(&obj, id)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}
	c.JSON(http.StatusOK, obj)
}

func apiCommitmentList(c *gin.Context) {
	query := c.Param("query")
	pageNumber := c.Query("page")
	perPage := c.Query("perpage")
	var obj models.Commitments
	var err error
	if query == "" {
		err = common.DB.All(&obj)
	} else {
		query := common.DB.Where("name LIKE CONCAT('%', ?, '%')", query) // TODO: FIXME: XXX
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
