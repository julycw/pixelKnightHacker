package command

import (
	"encoding/json"
	"github.com/julycw/pixelKnightHacker/models"
)

func GetExploreInfo(serverUrl, sid, userid string) (*models.ExploreInfo, error) {
	res, err := send(serverUrl, sid, userid, "5002", nil)
	if err != nil {
		return nil, err
	}
	v := new(models.ExploreInfo)
	if err := json.Unmarshal(res, v); err != nil {
		return nil, err
	}
	return v, nil
}

func RefreshHungry(serverUrl, sid, userid string) (*models.Basic, error) {
	res, err := send(serverUrl, sid, userid, "5004", nil)
	if err != nil {
		return nil, err
	}
	v := new(models.Basic)
	if err := json.Unmarshal(res, v); err != nil {
		return nil, err
	}
	return v, nil
}

func GetFoodList(serverUrl, sid, userid string) (*models.FoodList, error) {
	res, err := send(serverUrl, sid, userid, "7001", nil)
	if err != nil {
		return nil, err
	}
	v := new(models.FoodList)
	if err := json.Unmarshal(res, v); err != nil {
		return nil, err
	}
	return v, nil
}

func EatFood(serverUrl, sid, userid, foodid string) (*models.EatFootResult, error) {
	res, err := send(serverUrl, sid, userid, "7002", map[string]string{
		"foodid": foodid,
	})
	if err != nil {
		return nil, err
	}
	v := new(models.EatFootResult)
	if err := json.Unmarshal(res, v); err != nil {
		return nil, err
	}
	return v, nil
}

func GetRewardCount(serverUrl, sid, userid string) (*models.RewardCount, error) {
	res, err := send(serverUrl, sid, userid, "4008", nil)
	if err != nil {
		return nil, err
	}
	v := new(models.RewardCount)
	if err := json.Unmarshal(res, v); err != nil {
		return nil, err
	}
	return v, nil
}

func GetRewards(serverUrl, sid, userid string) (*models.ExploreRewardsList, error) {
	res, err := send(serverUrl, sid, userid, "5007", nil)
	if err != nil {
		return nil, err
	}
	v := new(models.ExploreRewardsList)
	if err := json.Unmarshal(res, v); err != nil {
		return nil, err
	}
	return v, nil
}

func UseAllExpItem(serverUrl, sid, userid string) (*models.ExploreRewardsList, error) {
	res, err := send(serverUrl, sid, userid, "8003", nil)
	if err != nil {
		return nil, err
	}
	v := new(models.ExploreRewardsList)
	if err := json.Unmarshal(res, v); err != nil {
		return nil, err
	}
	return v, nil
}

func UseAllMoneyItem(serverUrl, sid, userid string) (*models.ExploreRewardsList, error) {
	res, err := send(serverUrl, sid, userid, "8004", nil)
	if err != nil {
		return nil, err
	}
	v := new(models.ExploreRewardsList)
	if err := json.Unmarshal(res, v); err != nil {
		return nil, err
	}
	return v, nil
}
