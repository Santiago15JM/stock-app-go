package service

import (
	"stock-app/db"
	"stock-app/model"
)

func GetAllStock() ([]model.Stock, error) {
	return db.GetAllStock()
}

func GetStock(ticker string) (model.Stock, error) {
	return db.GetStockByTicker(ticker)
}

func GetQueriedStocks(search string, sortingType string, ascending bool) ([]model.Stock, error) {
	return db.GetFilteredSortedStocks(search, sortingType, ascending)
}
