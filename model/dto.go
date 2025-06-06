package model

type StockApiItem struct {
	Ticker     string `json:"ticker"`
	TargetFrom string `json:"target_from"`
	TargetTo   string `json:"target_to"`
	Company    string `json:"company"`
	Action     string `json:"action"`
	Brokerage  string `json:"brokerage"`
	RatingFrom string `json:"rating_from"`
	RatingTo   string `json:"rating_to"`
	Time       string `json:"time"`
}

type StockApiResponse struct {
	Items    []StockApiItem `json:"items"`
	NextPage string         `json:"next_page"`
}

type ScoredStock struct {
	Stock Stock
	Score float32
}