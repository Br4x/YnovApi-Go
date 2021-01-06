package handlers

import (
	"github.com/gin-gonic/gin"
	"ynov_immo/models"
)

func init() {
	groupApi.GET("real-estate-propal", realEstatePropalAll)
	groupApi.GET("real-estate-propal/:id", realEstatePropalOne)
	groupApi.POST("real-estate-propal", realEstatePropalCreate)
	groupApi.PATCH("real-estate-propal", realEstatePropalUpdate)
	groupApi.DELETE("real-estate-propal/:id", realEstatePropalDelete)
}

//All
func realEstatePropalAll(c *gin.Context) {
	mdl := models.RealEstatePropal{}
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
func realEstatePropalOne(c *gin.Context) {
	var mdl models.RealEstatePropal
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
func realEstatePropalCreate(c *gin.Context) {
	var mdl models.RealEstatePropal
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
func realEstatePropalUpdate(c *gin.Context) {
	var mdl models.RealEstatePropal
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
func realEstatePropalDelete(c *gin.Context) {
	var mdl models.RealEstatePropal
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
