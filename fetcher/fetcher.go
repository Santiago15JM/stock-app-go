package fetcher

import (
	"encoding/json"
	"errors"
	"github.com/lib/pq"
	"log"
	"net/http"
	"net/url"
	"os"
	"stock-app/db"
	"stock-app/model"
	"strconv"
	"strings"
	"time"
)

func Sync() error {
	log.Println("Started sync")
	next := ""
	for {
		items, nextPage, err := requestPage(next)
		if err != nil {
			return err
		}
		newStocks, err := bulkMap(items)
		if err != nil {
			return err
		}
		err = db.SaveAllStock(newStocks)
		if err != nil {
			var pqErr *pq.Error
			if errors.As(err, &pqErr) && pqErr.Code == "23505" { //If duplicate key
				break
			} else {
				return err
			}
		}

		if nextPage == "" {
			break
		}
		next = nextPage
	}
	return nil
}

func requestPage(next_page string) ([]model.StockApiItem, string, error) {
	var STOCK_API_URL = os.Getenv("STOCK_API_URL")
	var STOCK_API_TOKEN = os.Getenv("STOCK_API_TOKEN")

	req_url, err := url.Parse(STOCK_API_URL)
	if err != nil {
		return nil, "", errors.New("couldnt parse api url")
	}

	if len(next_page) != 0 {
		param := url.Values{}
		param.Add("next_page", next_page)

		req_url.RawQuery = param.Encode()
	}

	req, err := http.NewRequest(http.MethodGet, req_url.String(), nil)
	if err != nil {
		return nil, "", errors.New("couldnt create sync request")
	}

	req.Header.Add("Authorization", "Bearer "+STOCK_API_TOKEN)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, "", errors.New("couldnt make sync request")
	}
	defer res.Body.Close()

	var result model.StockApiResponse
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&result)
	if err != nil {
		return nil, "", errors.New("couldnt decode response")
	}

	return result.Items, result.NextPage, nil
}

func mapStockItem(item model.StockApiItem) (model.Stock, error) {
	targetFrom, err := strconv.ParseFloat(strings.ReplaceAll(strings.TrimPrefix(item.TargetFrom, "$"), ",", ""), 32)
	if err != nil {
		return model.Stock{}, errors.New("couldnt parse targetFrom: " + item.TargetFrom)
	}
	targetTo, err := strconv.ParseFloat(strings.ReplaceAll(strings.TrimPrefix(item.TargetTo, "$"), ",", ""), 32)
	if err != nil {
		return model.Stock{}, errors.New("couldnt parse targetTo: " + item.TargetTo)
	}
	parsedTime, err := time.Parse(time.RFC3339Nano, item.Time)
	if err != nil {
		return model.Stock{}, errors.New("couldnt parse time: " + item.Time)
	}

	stock := model.Stock{
		Ticker:     item.Ticker,
		TargetFrom: float32(targetFrom),
		TargetTo:   float32(targetTo),
		Company:    item.Company,
		Action:     item.Action,
		Brokerage:  item.Brokerage,
		RatingFrom: item.RatingFrom,
		RatingTo:   item.RatingTo,
		Time:       parsedTime,
	}

	return stock, nil
}

func bulkMap(items []model.StockApiItem) ([]model.Stock, error) {
	stocks := []model.Stock{}

	for _, item := range items {
		stock, err := mapStockItem(item)
		if err != nil {
			return nil, err
		}
		stocks = append(stocks, stock)
	}

	return stocks, nil
}
