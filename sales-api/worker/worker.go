package worker

import (
	"fmt"
	"github.com/fgarciareyna/SALES_FACUNDO_GARCIA/sales-api/cache"
	"github.com/fgarciareyna/SALES_FACUNDO_GARCIA/sales-api/db"
	"time"
)

var (
	ticker = time.NewTicker(time.Minute)
)

func StartWorker() {
	go func() {
		for {
			<- ticker.C
			UpdateCache()
		}
	}()
}

func UpdateCache() {
	totals, err := db.GetTotalByCountry()
	if err != nil {
		fmt.Printf("[db_error] %s\n", err.Error())
		return
	}
	cache.SetValues(totals)
}
