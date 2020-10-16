package services

import (
	"github.com/fgarciareyna/SALES_FACUNDO_GARCIA/sales-api/cache"
	"github.com/fgarciareyna/SALES_FACUNDO_GARCIA/sales-api/model"
)

func GetStats() []model.Stat {
	return cache.GetValues()
}
