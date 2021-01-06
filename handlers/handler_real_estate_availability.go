package handlers

import (
	"github.com/gin-gonic/gin"
	"ynov_immo/models"
)

func init() {
	groupApi.GET("real-estate-availability", realEstateAvailabilityAll)
	groupApi.GET("real-estate-availability/:id", realEstateAvailabilityOne)
	groupApi.POST("real-estate-availability", realEstateAvailabilityCreate)
	groupApi.PATCH("real-estate-availability", realEstateAvailabilityUpdate)
	groupApi.DELETE("real-estate-availability/:id", realEstateAvailabilityDelete)
}

//All
func realEstateAvailabilityAll(c *gin.Context) {
	mdl := models.RealEstateAvailability{}
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
func realEstateAvailabilityOne(c *gin.Context) {
	var mdl models.RealEstateAvailability
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
func realEstateAvailabilityCreate(c *gin.Context) {
	var mdl models.RealEstateAvailability
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
func realEstateAvailabilityUpdate(c *gin.Context) {
	var mdl models.RealEstateAvailability
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
func realEstateAvailabilityDelete(c *gin.Context) {
	var mdl models.RealEstateAvailability
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
