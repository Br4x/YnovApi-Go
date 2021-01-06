package handlers

import (
	"github.com/gin-gonic/gin"
	"ynov_immo/models"
)

func init() {
	groupApi.GET("real-estate-propals-buyer-indication", realEstatePropalsBuyerIndicationAll)
	groupApi.GET("real-estate-propals-buyer-indication/:id", realEstatePropalsBuyerIndicationOne)
	groupApi.POST("real-estate-propals-buyer-indication", realEstatePropalsBuyerIndicationCreate)
	groupApi.PATCH("real-estate-propals-buyer-indication", realEstatePropalsBuyerIndicationUpdate)
	groupApi.DELETE("real-estate-propals-buyer-indication/:id", realEstatePropalsBuyerIndicationDelete)
}

//All
func realEstatePropalsBuyerIndicationAll(c *gin.Context) {
	mdl := models.RealEstatePropalsBuyerIndication{}
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
func realEstatePropalsBuyerIndicationOne(c *gin.Context) {
	var mdl models.RealEstatePropalsBuyerIndication
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
func realEstatePropalsBuyerIndicationCreate(c *gin.Context) {
	var mdl models.RealEstatePropalsBuyerIndication
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
func realEstatePropalsBuyerIndicationUpdate(c *gin.Context) {
	var mdl models.RealEstatePropalsBuyerIndication
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
func realEstatePropalsBuyerIndicationDelete(c *gin.Context) {
	var mdl models.RealEstatePropalsBuyerIndication
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
