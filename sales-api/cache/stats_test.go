package cache

import (
	"github.com/fgarciareyna/SALES_FACUNDO_GARCIA/sales-api/model"
	"reflect"
	"testing"
)

func TestGetValues(t *testing.T) {
	cases := []struct {
		name        string
		cacheValues map[string]int
		expected    []model.Stat
	}{
		{
			name:        "Empty cache",
			cacheValues: make(map[string]int),
			expected:    make([]model.Stat, 0),
		},
		{
			name: "Some values in cache",
			cacheValues: map[string]int{
				"ar": 8,
				"br": 15,
			},
			expected: []model.Stat{
				{
					Country: "ar",
					Total:   8,
				}, {
					Country: "br",
					Total:   15,
				},
			},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			statsCache = c.cacheValues
			if result := GetValues(); !reflect.DeepEqual(result, c.expected) {
				t.Errorf("Result: %v - Expected %v", result, c.expected)
			}
		})
	}
}

func TestSetValues(t *testing.T) {
	cases := []struct {
		name        string
		cacheValues map[string]int
		values      []model.Stat
		expected    map[string]int
	}{
		{
			name:        "Set values in clean cache",
			cacheValues: make(map[string]int),
			values: []model.Stat{
				{
					Country: "ar",
					Total:   8,
				}, {
					Country: "br",
					Total:   15,
				},
			},
			expected: map[string]int{
				"ar": 8,
				"br": 15,
			},
		},
		{
			name: "Update values in cache",
			cacheValues: map[string]int{
				"ar": 8,
				"br": 15,
			},
			values: []model.Stat{
				{
					Country: "co",
					Total:   6,
				}, {
					Country: "ar",
					Total:   2,
				},
			},
			expected: map[string]int{
				"co": 6,
				"ar": 2,
			},
		},
		{
			name: "Clear cache",
			cacheValues: map[string]int{
				"ar": 8,
				"br": 15,
			},
			values: make([]model.Stat, 0),
			expected: make(map[string]int),
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			statsCache = c.cacheValues
			SetValues(c.values)
			if !reflect.DeepEqual(statsCache, c.expected) {
				t.Errorf("Result: %v - Expected %v", statsCache, c.expected)
			}
		})
	}
}
