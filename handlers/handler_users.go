package handlers

import (
	"ynov_immo/models"

	"github.com/gin-gonic/gin"
)

// to add auth, just insert jwtMiddleware like so ("user",jwtMiddleware, userAll)

func init() {
	groupApi.GET("user", userAll)
	groupApi.GET("user/:id", userOne)
	groupApi.POST("user", userCreate)
	groupApi.PATCH("user", userUpdate)
	groupApi.DELETE("user/:id", userDelete)
}

//All
func userAll(c *gin.Context) {
	mdl := models.User{}
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
func userOne(c *gin.Context) {
	var mdl models.User
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
func userCreate(c *gin.Context) {
	var mdl models.User
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
func userUpdate(c *gin.Context) {
	var mdl models.User
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
func userDelete(c *gin.Context) {
	var mdl models.User
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
