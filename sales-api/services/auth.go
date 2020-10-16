package services

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/fgarciareyna/SALES_FACUNDO_GARCIA/sales-api/config"
	"github.com/fgarciareyna/SALES_FACUNDO_GARCIA/sales-api/model"
	"time"
)

func ValidateUser(user model.User) (string, bool) {
	// Validate user and pass
	// Get user role
	return "admin", true
}

func GenerateToken(user model.User) (string, error) {
	token := jwt.New(jwt.GetSigningMethod("HS256"))

	token.Claims = jwt.MapClaims{
		"Name": user.Name,
		"Role": user.Role,
		"exp":  time.Now().Add(time.Hour).Unix(),
	}

	tokenString, err := token.SignedString([]byte(config.Secret))
	if err != nil {
		fmt.Printf("[auth_error] %s\n", err.Error())
		return "", err
	}

	return tokenString, nil
}
