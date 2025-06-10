package db

import (
	"database/sql"
	"fmt"
	"stock-app/model"
	"strconv"
)

func CreateStockTable() error {
	query := `
    CREATE TABLE IF NOT EXISTS stocks (
        ticker       TEXT,
        company      TEXT,
        action       TEXT,
        brokerage    TEXT,
        rating_from  TEXT,
        rating_to    TEXT,
        target_from  FLOAT,
        target_to    FLOAT,
        time         TIMESTAMPTZ,
        PRIMARY KEY (ticker)
    )`

	_, err := DB.Exec(query)
	if err != nil {
		return fmt.Errorf("error creating table: %w", err)
	}
	return nil
}

func SaveStock(stock model.Stock) error {
	query := `
        INSERT INTO stocks (
            ticker, company, action, brokerage,
            rating_from, rating_to, target_from, target_to, time
        ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		ON CONFLICT (ticker) DO UPDATE SET
			company = EXCLUDED.company,
			action = EXCLUDED.action,
			brokerage = EXCLUDED.brokerage,
			rating_from = EXCLUDED.rating_from,
			rating_to = EXCLUDED.rating_to,
			target_from = EXCLUDED.target_from,
			target_to = EXCLUDED.target_to,
			time = EXCLUDED.time;
    `
	_, err := DB.Exec(query,
		stock.Ticker,
		stock.Company,
		stock.Action,
		stock.Brokerage,
		stock.RatingFrom,
		stock.RatingTo,
		stock.TargetFrom,
		stock.TargetTo,
		stock.Time,
	)

	if err != nil {
		return fmt.Errorf("failed to insert stock: %w", err)
	}
	return nil
}

func SaveAllStock(stocks []model.Stock) error {
	for _, item := range stocks {
		err := SaveStock(item)
		if err != nil {
			return err
		}
	}
	return nil
}

func parseRows(rows *sql.Rows) ([]model.Stock, error) {
	var stocks []model.Stock

	for rows.Next() {
		var stock model.Stock
		err := rows.Scan(
			&stock.Ticker,
			&stock.Company,
			&stock.Action,
			&stock.Brokerage,
			&stock.RatingFrom,
			&stock.RatingTo,
			&stock.TargetFrom,
			&stock.TargetTo,
			&stock.Time,
		)

		if err != nil {
			return nil, fmt.Errorf("row scan failed: %w", err)
		}
		stocks = append(stocks, stock)
	}
	return stocks, nil
}

func queryStocks(query string) ([]model.Stock, error) {
	rows, err := DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()

	stocks, err := parseRows(rows)
	if err != nil {
		return nil, fmt.Errorf("parse failed: %w", err)
	}

	return stocks, nil
}

func GetAllStock() ([]model.Stock, error) {
	query := `
        SELECT ticker, company, action, brokerage,
               rating_from, rating_to, target_from, target_to, time
        FROM stocks
        ORDER BY ticker
    `
	return queryStocks(query)
}

func GetRecent(days int) ([]model.Stock, error) {
	query := `
        SELECT ticker, company, action, brokerage,
               rating_from, rating_to, target_from, target_to, time
		FROM stocks
		WHERE time >= NOW() - INTERVAL '` + strconv.Itoa(days) + ` days'
		ORDER BY time DESC
    `
	return queryStocks(query)
}

func GetStockByTicker(ticker string) (model.Stock, error) {
	query := `
        SELECT ticker, company, action, brokerage,
               rating_from, rating_to, target_from, target_to, time
		FROM stocks
		WHERE ticker = '` + ticker + `'`
	stock, err := queryStocks(query)
	if len(stock) == 0 {
		return model.Stock{}, fmt.Errorf("couldnt find stock")
	}
	return stock[0], err
}

func GetFilteredSortedStocks(search string, sortingType string, ascending bool) ([]model.Stock, error) {
	query := `
        SELECT ticker, company, action, brokerage,
               rating_from, rating_to, target_from, target_to, time
		FROM stocks
		`
	if search != "" {
		query += `WHERE
  			ticker LIKE('%` + search + `%') or
			company LIKE('%` + search + `%') or
			brokerage LIKE('%` + search + `%')
		`
	}
	query += ` ORDER BY ` + sortingType

	if ascending {
		query += ` ASC`
	} else {
		query += ` DESC`
	}

	stocks, err := queryStocks(query)
	if len(stocks) == 0 {
		return nil, fmt.Errorf("couldnt find stock")
	}
	return stocks, err
}
