package api // import "github.com/dayvillefire/tenders/api"

import (
	"net/http"

	"github.com/dayvillefire/tenders/common"
	"github.com/dayvillefire/tenders/models"
	"github.com/gin-gonic/gin"
)

func init() {
	common.ApiMap["department"] = func(r *gin.RouterGroup) {
		r.PUT("", apiDepartmentSave)
		r.GET("/:id", apiDepartmentGet)
		r.DELETE("/:id", apiDepartmentDelete)
	}
	common.ApiMap["departments"] = func(r *gin.RouterGroup) {
		r.GET("", apiDepartmentList)
		r.GET("/:query", apiDepartmentList)
	}
}

func apiDepartmentSave(c *gin.Context) {
	var obj models.Department
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

func apiDepartmentDelete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	var obj models.Department
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

func apiDepartmentGet(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	var obj models.Department
	err := common.DB.Find(&obj, id)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}
	c.JSON(http.StatusOK, obj)
}

func apiDepartmentList(c *gin.Context) {
	query := c.Param("query")
	pageNumber := c.Query("page")
	perPage := c.Query("perpage")
	var obj models.Departments
	var err error
	if query == "" {
		err = common.DB.All(&obj)
	} else {
		query := common.DB.Where("name LIKE CONCAT('%', ?, '%')", query)
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
