package handlers

import (
	"github.com/gin-gonic/gin"
	"ynov_immo/models"
)

func init() {
	groupApi.GET("real-estate-image", realEstateImageAll)
	groupApi.GET("real-estate-image/:id", realEstateImageOne)
	groupApi.POST("real-estate-image", realEstateImageCreate)
	groupApi.PATCH("real-estate-image", realEstateImageUpdate)
	groupApi.DELETE("real-estate-image/:id", realEstateImageDelete)
}

//All
func realEstateImageAll(c *gin.Context) {
	mdl := models.RealEstateImage{}
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
func realEstateImageOne(c *gin.Context) {
	var mdl models.RealEstateImage
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
func realEstateImageCreate(c *gin.Context) {
	var mdl models.RealEstateImage
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
func realEstateImageUpdate(c *gin.Context) {
	var mdl models.RealEstateImage
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
func realEstateImageDelete(c *gin.Context) {
	var mdl models.RealEstateImage
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
