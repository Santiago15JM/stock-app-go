package service

import (
	"math"
	"sort"
	"stock-app/db"
	"stock-app/model"
	"strings"
	"time"
)

func FindBestPick() []model.ScoredStock {
	recent_stocks, _ := db.GetRecent(5)

	stock_pool := filterByRating(recent_stocks)

	scores := scoreStocks(stock_pool)

	sort.Slice(scores, func(i, j int) bool {
		return scores[i].Score > scores[j].Score
	})

	if len(scores) < 5 {
		return scores
	}
	return scores[:5]
}

func filterByRating(stocks []model.Stock) []model.Stock {
	var filtered []model.Stock
	for _, stock := range stocks {
		normRating := strings.ToLower(stock.RatingTo)
		weight, ok := ratingWeights[normRating]

		if ok && weight > 1.0 && stock.TargetFrom < stock.TargetTo {
			filtered = append(filtered, stock)
		}
	}
	return filtered
}

func scoreStocks(stocks []model.Stock) []model.ScoredStock {
	scores := []model.ScoredStock{}
	for _, stock := range stocks {
		improvement := calculateImprovement(stock.TargetFrom, stock.TargetTo)
		normRating := strings.ToLower(stock.RatingTo)
		ratingWeight := ratingWeights[normRating]
		recencyWeight := getRecencyWeight(stock.Time)

		score := ratingWeight * improvement * recencyWeight
		scores = append(scores, model.ScoredStock{Stock: stock, Score: score})
	}
	return scores
}

func calculateImprovement(target_from float32, target_to float32) float32 {
	return (target_to - target_from) / target_from
}

func getRecencyWeight(stockTime time.Time) float32 {
	dur := time.Since(stockTime)
	daysSince := math.Floor(dur.Hours() / 24)
	return float32(1.0 - (0.1 * daysSince))
}
