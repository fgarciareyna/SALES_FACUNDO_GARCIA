package db

import (
	"errors"
	"github.com/fgarciareyna/SALES_FACUNDO_GARCIA/sales-api/config"
	"github.com/fgarciareyna/SALES_FACUNDO_GARCIA/sales-api/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
	ErrorLimitExceeded = errors.New("limit exceeded")
)

type Postgres struct {
}

func InitDB() {
	if !config.Production {
		return
	}
	var err error
	db, err = gorm.Open(postgres.Open(config.DbConnection), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(&model.Ticket{}); err != nil {
		panic(err)
	}
}

func (Postgres) SaveTicket(ticket model.Ticket, limit uint) (model.Ticket, error) {
	if db == nil {
		InitDB()
	}
	tx := db.Begin()
	if result := tx.Save(&ticket); result.Error != nil {
		tx.Rollback()
		return ticket, result.Error
	}

	if ticket.ID >= limit {
		tx.Rollback()
		return ticket, ErrorLimitExceeded
	}

	tx.Commit()

	return ticket, nil
}

func GetTotalByCountry() ([]model.Stat, error) {
	if db == nil {
		InitDB()
	}
	var stats []model.Stat
	if result := db.Model(&model.Ticket{}).Select("country, count(id) as total").Group("country").Scan(&stats); result.Error != nil {
		return nil, result.Error
	}
	return stats, nil
}
