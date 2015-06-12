package models

type ExploreInfo struct {
	Basic
	Data ExploreInfoData `json:"data"`
}

type ExploreInfoData struct {
	Bzd int `json:"bzd"`
}

type FoodList struct {
	Basic
	Data []FoodItem `json:"data"`
}

type FoodItem struct {
	Id     int `json:"id"`
	FoodId int `json:"foodid"`
	Count  int `json:"count"`
}

type EatFootResult struct {
	Basic
	Data EatFootResultData `json:"data"`
}

type EatFootResultData struct {
	Rewards []Reward `json:"rewards"`
}

type RewardCount struct {
	Basic
	Data RewardCountData `json:"data"`
}

type RewardCountData struct {
	Count int `json:"count"`
}

type ExploreRewardsList struct {
	Basic
	Data ExploreRewardsListData `json:"data"`
}

type ExploreRewardsListData struct {
	Rewards []Reward `json:"rewards"`
}
