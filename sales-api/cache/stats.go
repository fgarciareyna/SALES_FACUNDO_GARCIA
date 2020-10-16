package cache

import (
	"github.com/fgarciareyna/SALES_FACUNDO_GARCIA/sales-api/model"
	"sync"
)

var (
	statsCache  = make(map[string]int)
	statsMux    sync.RWMutex
)

func GetValues() []model.Stat {
	statsMux.RLock()

	result := make([]model.Stat, 0)
	for k, v := range statsCache {
		result = append(result, model.Stat{Country: k, Total: v})
	}
	statsMux.RUnlock()
	return result
}

func SetValues(stats []model.Stat) {
	statsMux.Lock()
	defer statsMux.Unlock()

	statsCache = make(map[string]int)
	for _, t := range stats {
		statsCache[t.Country] = t.Total
	}
}
