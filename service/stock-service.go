package service

import (
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
	return db.GetFilteredSortedStocks(search, sortingType, ascending, limit, page)
}
