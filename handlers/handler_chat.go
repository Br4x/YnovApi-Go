package handlers

import (
	"github.com/gin-gonic/gin"
	"ynov_immo/models"
)

func init() {
	groupApi.GET("chat", chatAll)
	groupApi.GET("chat/:id", chatOne)
	groupApi.POST("chat", chatCreate)
	groupApi.PATCH("chat", chatUpdate)
	groupApi.DELETE("chat/:id", chatDelete)
}

//All
func chatAll(c *gin.Context) {
	mdl := models.Chat{}
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
func chatOne(c *gin.Context) {
	var mdl models.Chat
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
func chatCreate(c *gin.Context) {
	var mdl models.Chat
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
func chatUpdate(c *gin.Context) {
	var mdl models.Chat
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
func chatDelete(c *gin.Context) {
	var mdl models.Chat
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
