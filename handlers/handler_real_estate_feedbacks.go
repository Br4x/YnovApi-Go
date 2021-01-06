package handlers

import (
	"github.com/gin-gonic/gin"
	"ynov_immo/models"
)

func init() {
	groupApi.GET("real-estate-feedback", realEstateFeedbackAll)
	groupApi.GET("real-estate-feedback/:id", realEstateFeedbackOne)
	groupApi.POST("real-estate-feedback", realEstateFeedbackCreate)
	groupApi.PATCH("real-estate-feedback", realEstateFeedbackUpdate)
	groupApi.DELETE("real-estate-feedback/:id", realEstateFeedbackDelete)
}

//All
func realEstateFeedbackAll(c *gin.Context) {
	mdl := models.RealEstateFeedback{}
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
func realEstateFeedbackOne(c *gin.Context) {
	var mdl models.RealEstateFeedback
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
func realEstateFeedbackCreate(c *gin.Context) {
	var mdl models.RealEstateFeedback
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
func realEstateFeedbackUpdate(c *gin.Context) {
	var mdl models.RealEstateFeedback
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
func realEstateFeedbackDelete(c *gin.Context) {
	var mdl models.RealEstateFeedback
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
