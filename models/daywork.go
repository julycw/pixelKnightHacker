package models

type DayworkReward struct {
	Basic
	Data DayworkRewardData `json:"data"`
}

type DayworkRewardData struct {
	Rewards []Reward `json:"rewards"`
}
