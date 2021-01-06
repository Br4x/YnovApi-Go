package handlers

import (
	"github.com/gin-gonic/gin"
	"ynov_immo/models"
)

func init() {
	groupApi.GET("chat-comment", chatCommentAll)
	groupApi.GET("chat-comment/:id", chatCommentOne)
	groupApi.POST("chat-comment", chatCommentCreate)
	groupApi.PATCH("chat-comment", chatCommentUpdate)
	groupApi.DELETE("chat-comment/:id", chatCommentDelete)
}

//All
func chatCommentAll(c *gin.Context) {
	mdl := models.ChatComment{}
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
func chatCommentOne(c *gin.Context) {
	var mdl models.ChatComment
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
func chatCommentCreate(c *gin.Context) {
	var mdl models.ChatComment
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
func chatCommentUpdate(c *gin.Context) {
	var mdl models.ChatComment
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
func chatCommentDelete(c *gin.Context) {
	var mdl models.ChatComment
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
