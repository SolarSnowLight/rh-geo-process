package handler

import (
	utilContext "geo-process/pkg/handler/util"
	model "geo-process/pkg/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary RegionGetAll
// @Tags API для получения географических данных
// @Description Получение списка всех регионов, определённых в БД
// @ID region-get-all
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Токен доступа для текущего пользователя" example(Bearer access_token)
// @Success 200 {object} model.RegionModel "data"
// @Failure 400,404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Failure default {object} model.ResponseMessage
// @Router /region/get/all [get]
func (h *Handler) RegionGetAll(c *gin.Context) {
	// Заглушка
	_ = &model.CityDB{}

	data, err := h.services.Geo.GetRegionAll()
	if err != nil {
		utilContext.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, data)
}

// @Summary CityGetAllByRegion
// @Tags API для получения географических данных
// @Description Получение списка всех городов определённого региона
// @ID city-get-all-by-region
// @Accept  json
// @Produce  json
// @Param 	Authorization header string true "Токен доступа для текущего пользователя" example(Bearer access_token)
// @Param 	region_id  	  path      	int  true  "Region ID"
// @Success 200 		  {object} 		model.RegionModel "data"
// @Failure 400,404 	  {object} 		model.ResponseMessage
// @Failure 500 		  {object} 		model.ResponseMessage
// @Failure default 	  {object} 		model.ResponseMessage
// @Router /city/get/all/{region_id} [get]
func (h *Handler) CityGetAllByRegion(c *gin.Context) {
	region, err := strconv.Atoi(c.Param("region_id"))
	if err != nil {
		utilContext.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	data, err := h.services.Geo.GetCitiesByRegion(region)
	if err != nil {
		utilContext.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, data)
}

// @Summary CityGetAll
// @Tags API для получения географических данных
// @Description Получение списка всех городов РФ
// @ID city-get-all
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Токен доступа для текущего пользователя" example(Bearer access_token)
// @Success 200 {object} model.RegionModel "data"
// @Failure 400,404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Failure default {object} model.ResponseMessage
// @Router /city/get/all [get]
func (h *Handler) CityGetAll(c *gin.Context) {
	data, err := h.services.Geo.GetCities()
	if err != nil {
		utilContext.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, data)
}
