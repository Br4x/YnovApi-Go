package handlers

import (
	"github.com/gin-gonic/gin"
	"ynov_immo/models"
)

func init() {
	groupApi.GET("real-estate", realEstateAll)
	groupApi.GET("real-estate/:id", realEstateOne)
	groupApi.POST("real-estate", realEstateCreate)
	groupApi.PATCH("real-estate", realEstateUpdate)
	groupApi.DELETE("real-estate/:id", realEstateDelete)
}

//All
func realEstateAll(c *gin.Context) {
	mdl := models.RealEstate{}
	query := &models.PaginationQuery{}
	err := c.ShouldBindQuery(query)
	if handleError(c, err) {
		return
	}
	list, total, err := mdl.All(query)
	if handleError(c, err) {
		return
	}
	jsonPagination(c, list, total, query)
}

//One
func realEstateOne(c *gin.Context) {
	var mdl models.RealEstate
	id, err := parseParamID(c)
	if handleError(c, err) {
		return
	}
	mdl.Id = id
	data, err := mdl.One()
	if handleError(c, err) {
		return
	}
	jsonData(c, data)
}

//Create
func realEstateCreate(c *gin.Context) {
	var mdl models.RealEstate
	err := c.ShouldBind(&mdl)
	if handleError(c, err) {
		return
	}
	err = mdl.Create()
	if handleError(c, err) {
		return
	}
	jsonData(c, mdl)
}

//Update
func realEstateUpdate(c *gin.Context) {
	var mdl models.RealEstate
	err := c.ShouldBind(&mdl)
	if handleError(c, err) {
		return
	}
	err = mdl.Update()
	if handleError(c, err) {
		return
	}
	jsonSuccess(c)
}

//Delete
func realEstateDelete(c *gin.Context) {
	var mdl models.RealEstate
	id, err := parseParamID(c)
	if handleError(c, err) {
		return
	}
	mdl.Id = id
	err = mdl.Delete()
	if handleError(c, err) {
		return
	}
	jsonSuccess(c)
}
