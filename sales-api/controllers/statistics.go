package controllers

import (
	"github.com/fgarciareyna/SALES_FACUNDO_GARCIA/sales-api/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Stats(c *gin.Context) {
	c.JSON(http.StatusOK, services.GetStats())
}
