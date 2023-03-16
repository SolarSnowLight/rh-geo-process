package handler

import (
	utilContext "geo-process/pkg/handler/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary RegionGetAll
// @Tags API для получения географических данных
// @Description Получение списка всех пользователей находящихся в системе
// @ID region-get-all
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Токен доступа для текущего пользователя" example(Bearer access_token)
// @Success 200 {object} adminModel.UsersResponseModel "data"
// @Failure 400,404 {object} httpModel.ResponseMessage
// @Failure 500 {object} httpModel.ResponseMessage
// @Failure default {object} httpModel.ResponseMessage
// @Router /region/get/all [post]
func (h *Handler) RegionGetAll(c *gin.Context) {
	data, err := h.services.Geo.GetRegionAll()
	if err != nil {
		utilContext.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, data)
}
