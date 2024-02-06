package api

import (
	"github.com/BeeOntime/api/docs"
	"github.com/BeeOntime/api/handlers"
	"github.com/BeeOntime/config"
	"github.com/BeeOntime/storage"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	// swaggerFiles "github.com/swaggo/files"
)

func SetUpAPI(cfg config.Config, strg storage.StorageI) *gin.Engine {
	docs.SwaggerInfo.Title = "Bee on Time API"
	docs.SwaggerInfo.Version = "1.0"
	// docs.SwaggerInfo.Host = cfg.ServiceHost + cfg.HTTPPort
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	h := handlers.NewHandler(&handlers.Handler{
		Cfg:     cfg,
		Storage: strg,
	})

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowCredentials = true
	corsConfig.AllowHeaders = []string{"*"}
	corsConfig.AllowBrowserExtensions = true
	corsConfig.AllowMethods = []string{"*"}
	r.Use(cors.New(corsConfig))
	api := r.Group("/v1")

	api.GET("/ping", h.Ping)

	// //staffs
	api.POST("/staff", h.CreateStaff)
	api.GET("/staffs", h.GetStaffs)
	api.DELETE("/staff/:id", h.DeleteStaff)
	api.PUT("/staff", h.UpdateStaff)

	//entry points
	api.POST("/staff/entry", h.CreateStaffEntry)
	api.GET("/staff/entry", h.GetStaffEntries)
	api.DELETE("/staff/entry/:id", h.DeleteStaffEntry)
	api.PUT("/staff/entry", h.UpdateStaffEntry)

	//leaves
	api.POST("/staff/leave", h.CreateStaffLeave)
	api.GET("/staff/leave", h.GetStaffLeaves)

	url := ginSwagger.URL("swagger/doc.json")
	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	// r.Any("/api/*any", h.AuthMiddleware(cfg), proxyMiddleware(r, &h), h.Proxy)
	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r

}
