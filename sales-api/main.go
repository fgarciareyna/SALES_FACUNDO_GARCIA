package main

import (
	"github.com/fgarciareyna/SALES_FACUNDO_GARCIA/sales-api/router"
	"github.com/fgarciareyna/SALES_FACUNDO_GARCIA/sales-api/worker"
)

func main() {
	worker.StartWorker()
	router.StartRouter()
}
