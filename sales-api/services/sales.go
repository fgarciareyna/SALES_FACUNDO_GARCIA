package services

import (
	"errors"
	"fmt"
	"github.com/fgarciareyna/SALES_FACUNDO_GARCIA/sales-api/config"
	"github.com/fgarciareyna/SALES_FACUNDO_GARCIA/sales-api/db"
	"github.com/fgarciareyna/SALES_FACUNDO_GARCIA/sales-api/model"
)

var (
	database = db.DB
	ErrorNoTicketsLeft = errors.New("there are no available tickets left")
)

func RegisterSale(ticket model.Ticket) (model.Ticket, error) {
	result, err := database.SaveTicket(ticket, config.AvailableTickets)
	if err != nil {
		if err == db.ErrorLimitExceeded {
			return model.Ticket{}, ErrorNoTicketsLeft
		}
		fmt.Printf("[db_error] %s\n", err.Error())
		return model.Ticket{}, err
	}

	return result, nil
}
