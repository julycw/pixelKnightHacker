package models

type AddMineQueueResult struct {
	Basic
}

type MineQueue struct {
	Basic
	Data []MineQueueDataItem `json:"data"`
}

type MineQueueDataItem struct {
	Id   int `json:"id"`
	Sid  int `json:"sid"`
	Time int `json:"time`
}

type MineQueueReward struct {
	Basic
}

type MineInfo struct {
	Basic
	Data MineInfoData `json:"data"`
}

type MineInfoData struct {
	Szg int `json:"szg"`
}
