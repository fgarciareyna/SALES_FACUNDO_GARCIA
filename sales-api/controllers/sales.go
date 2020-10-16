package controllers

import (
	"github.com/fgarciareyna/SALES_FACUNDO_GARCIA/sales-api/model"
	"github.com/fgarciareyna/SALES_FACUNDO_GARCIA/sales-api/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterSale(c *gin.Context) {
	var ticket model.Ticket
	if err := c.BindJSON(&ticket); err != nil {
		c.String(http.StatusBadRequest, "Invalid ticket info")
		return
	}

	country := c.Param("country")
	ticket.Country = country

	result, err := services.RegisterSale(ticket)
	if err != nil {
		if err == services.ErrorNoTicketsLeft {
			c.String(http.StatusBadRequest, "There are no available tickets left")
			return
		}
		c.String(http.StatusInternalServerError, "Unexpected error registering sale")
	}

	c.JSON(http.StatusOK, result)
}
