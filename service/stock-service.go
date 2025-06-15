package service

import (
	"fmt"
	"stock-app/db"
	"stock-app/model"
)

func GetAllStock(limit int, page int) ([]model.Stock, error) {
	return db.GetAllStock(limit, page)
}

func GetStock(ticker string) (model.Stock, error) {
	return db.GetStockByTicker(ticker)
}

func GetQueriedStocks(search string, sortingType string, ascending bool, limit int, page int) ([]model.Stock, error) {
	if !validateSort(sortingType) {
		return nil, fmt.Errorf("invalid sorting type")
	}
	return db.GetFilteredSortedStocks(search, sortingType, ascending, limit, page)
}

func validateSort(sortingType string) bool {
	switch sortingType {
	case "TICKER", "BROKERAGE", "TIME":
		return true
	}
	return false
}
