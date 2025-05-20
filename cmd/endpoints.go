package cmd

import (
	"fmt"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	_ "github.com/HenriqueOtsuka/goexpert-challenge-1/docs"
	"github.com/HenriqueOtsuka/goexpert-challenge-1/internal/controllers/app"
	"github.com/HenriqueOtsuka/goexpert-challenge-1/internal/controllers/cep"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title CEP API
// @version 1
// @description This is the CEP API
// @contact.name   Henrique Otsuka // @contact.url    https://www.github.com/HenriqueOtsuka
// @tag.name App
func SetupRouter() *gin.Engine {
	router := gin.New()
	middleware := NewMiddleware()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("[GIN] %s \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	router.Use(gin.Recovery())
	router.Use(gzip.Gzip(gzip.DefaultCompression))
	router.Use(middleware.Errors)
	router.Use(middleware.Options)
	router.NoRoute(middleware.NotFound)
	router.NoMethod(middleware.MethodNotAllowed)

	router.GET("appstatus", app.HandleGetAppStatus)
	router.GET("cep/:cep", cep.HandleGetCepTemperature)

	return router
}
