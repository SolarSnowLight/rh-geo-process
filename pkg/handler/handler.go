package handler

import (
	routeConstant "geo-process/pkg/constant"
	"geo-process/pkg/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	// Установка максимального размера тела Multipart
	router.MaxMultipartMemory = 50 << 20
	// router.Static("/public", "./public")

	// Установка CORS-политик
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{viper.GetString("client_url")},
		AllowMethods:     []string{"POST", "GET"},
		AllowHeaders:     []string{"Origin", "Content-type", "Authorization"},
		AllowCredentials: true,
	}))

	// URL: /swagger/index.html
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET(routeConstant.REGION+routeConstant.GET_ALL, h.RegionGetAll)

	return router
}
