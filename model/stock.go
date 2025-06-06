package model

import "time"

type Stock struct {
    Ticker      string    `json:"ticker"`
    TargetFrom  float32    `json:"target_from"`
    TargetTo    float32    `json:"target_to"`
    Company     string    `json:"company"`
    Action      string    `json:"action"`
    Brokerage   string    `json:"brokerage"`
    RatingFrom  string    `json:"rating_from"`
    RatingTo    string    `json:"rating_to"`
    Time        time.Time `json:"time"`
}
