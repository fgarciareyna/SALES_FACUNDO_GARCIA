package db

import "github.com/fgarciareyna/SALES_FACUNDO_GARCIA/sales-api/model"

var (
	DB Database = Postgres{}
)

type Database interface {
	SaveTicket(ticket model.Ticket, limit uint) (model.Ticket, error)
}
