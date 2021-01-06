package handlers

import (
	"github.com/gin-gonic/gin"
	"ynov_immo/models"
)

func init() {
	groupApi.GET("cities-district", citiesDistrictAll)
	groupApi.GET("cities-district/:id", citiesDistrictOne)
	groupApi.POST("cities-district", citiesDistrictCreate)
	groupApi.PATCH("cities-district", citiesDistrictUpdate)
	groupApi.DELETE("cities-district/:id", citiesDistrictDelete)
}

//All
func citiesDistrictAll(c *gin.Context) {
	mdl := models.CitiesDistrict{}
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
func citiesDistrictOne(c *gin.Context) {
	var mdl models.CitiesDistrict
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
func citiesDistrictCreate(c *gin.Context) {
	var mdl models.CitiesDistrict
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
func citiesDistrictUpdate(c *gin.Context) {
	var mdl models.CitiesDistrict
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
func citiesDistrictDelete(c *gin.Context) {
	var mdl models.CitiesDistrict
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
