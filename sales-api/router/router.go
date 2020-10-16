package router

import (
	"github.com/fgarciareyna/SALES_FACUNDO_GARCIA/sales-api/config"
	"github.com/fgarciareyna/SALES_FACUNDO_GARCIA/sales-api/controllers"
	"github.com/gin-gonic/contrib/jwt"
	"github.com/gin-gonic/gin"
)

func StartRouter() {
	router := gin.New()

	// Health
	router.GET("/ping", controllers.Ping)

	// Login
	router.POST("/login", controllers.Login)

	private := router.Group("/api")
	private.Use(jwt.Auth(config.Secret))

	// Sales
	private.POST("/v1/sales/:country", controllers.RegisterSale)

	// Statistics
	private.GET("/v1/stats", controllers.Stats)

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
