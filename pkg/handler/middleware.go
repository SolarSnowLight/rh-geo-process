package handler

import (
	"encoding/json"
	constant "geo-process/pkg/constant"
	utilContext "geo-process/pkg/handler/util"
	model "geo-process/pkg/model"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

/* Middleware для идентификация пользователя через основной сервис */
func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader("Authorization")

	if header == "" {
		utilContext.NewErrorResponse(c, http.StatusUnauthorized, "Пустой заголовок авторизации!")
		return
	}

	client := &http.Client{}

	req, err := http.NewRequest("POST", viper.GetString("main_url")+constant.SERVICE_EXTERNAL_VERIFY, nil)
	if err != nil {
		utilContext.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	req.Header.Add("Authorization", header)
	resp, err := client.Do(req)
	if err != nil {
		utilContext.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			utilContext.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
			return
		}

		var message model.ResponseMessage
		if err := json.Unmarshal(body, &message); err != nil {
			utilContext.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
			return
		}

		utilContext.NewErrorResponse(c, http.StatusUnauthorized, message.Message)
		return
	}
}
