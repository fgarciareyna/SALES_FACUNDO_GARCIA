package controllers

import (
	"github.com/fgarciareyna/SALES_FACUNDO_GARCIA/sales-api/model"
	"github.com/fgarciareyna/SALES_FACUNDO_GARCIA/sales-api/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	var user model.User
	if err := c.BindJSON(&user); err != nil {
		c.String(http.StatusBadRequest, "Invalid user info")
		return
	}

	role, isValid := services.ValidateUser(user)
	if !isValid {
		c.String(http.StatusNotFound, "User not found")
		return
	}

	user.Role = role

	token, err := services.GenerateToken(user)
		if err != nil {
			c.String(http.StatusInternalServerError, "Unexpected error logging user")
			return
	}

	c.JSON(200, gin.H{"token": token})
}
