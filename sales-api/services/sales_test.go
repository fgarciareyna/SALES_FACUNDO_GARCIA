package services

import (
	"errors"
	"github.com/fgarciareyna/SALES_FACUNDO_GARCIA/sales-api/config"
	"github.com/fgarciareyna/SALES_FACUNDO_GARCIA/sales-api/db"
	"github.com/fgarciareyna/SALES_FACUNDO_GARCIA/sales-api/model"
	"reflect"
	"testing"
)

type okDB struct {
}

func (okDB) SaveTicket(ticket model.Ticket, limit uint) (model.Ticket, error) {
	return ticket, nil
}

type limitDB struct {
}

func (limitDB) SaveTicket(ticket model.Ticket, limit uint) (model.Ticket, error) {
	return model.Ticket{}, db.ErrorLimitExceeded
}

type errDB struct {
}

func (errDB) SaveTicket(ticket model.Ticket, limit uint) (model.Ticket, error) {
	return model.Ticket{}, errors.New("test db error")
}

func TestRegisterSale(t *testing.T) {
	config.Production = false

	ticket := model.Ticket{}

	cases := []struct {
		name           string
		database       db.Database
		expectedResult model.Ticket
		expectedErr    error
	}{
		{
			name:           "Register ok",
			database:       okDB{},
			expectedResult: ticket,
			expectedErr:    nil,
		},
		{
			name:           "Too many tickets",
			database:       limitDB{},
			expectedResult: ticket,
			expectedErr:    ErrorNoTicketsLeft,
		},
		{
			name:           "Register error",
			database:       errDB{},
			expectedResult: ticket,
			expectedErr:    errors.New("test db error"),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			database = c.database
			result, err := RegisterSale(ticket)
			if !reflect.DeepEqual(err, c.expectedErr) {
				t.Errorf("Result: %v - Expected %v", err, c.expectedErr)
			}
			if !reflect.DeepEqual(result, c.expectedResult) {
				t.Errorf("Result: %v - Expected %v", result, c.expectedResult)
			}
		})
	}
}
