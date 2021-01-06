package handlers

import (
	"github.com/gin-gonic/gin"
	"ynov_immo/models"
)

func init() {
	groupApi.GET("real-estate-visit", realEstateVisitAll)
	groupApi.GET("real-estate-visit/:id", realEstateVisitOne)
	groupApi.POST("real-estate-visit", realEstateVisitCreate)
	groupApi.PATCH("real-estate-visit", realEstateVisitUpdate)
	groupApi.DELETE("real-estate-visit/:id", realEstateVisitDelete)
}

//All
func realEstateVisitAll(c *gin.Context) {
	mdl := models.RealEstateVisit{}
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
func realEstateVisitOne(c *gin.Context) {
	var mdl models.RealEstateVisit
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
func realEstateVisitCreate(c *gin.Context) {
	var mdl models.RealEstateVisit
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
func realEstateVisitUpdate(c *gin.Context) {
	var mdl models.RealEstateVisit
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
func realEstateVisitDelete(c *gin.Context) {
	var mdl models.RealEstateVisit
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
