package command

import (
	"encoding/json"
	"github.com/julycw/pixelKnightHacker/models"
)

func GetMonthCardReward1(serverUrl, sid, userid string) (*models.DayworkReward, error) {
	res, err := send(serverUrl, sid, userid, "190201", map[string]string{"id": "7"})
	if err != nil {
		return nil, err
	}
	v := new(models.DayworkReward)
	if err := json.Unmarshal(res, v); err != nil {
		return nil, err
	}
	return v, nil
}

func GetMonthCardReward2(serverUrl, sid, userid string) (*models.DayworkReward, error) {
	res, err := send(serverUrl, sid, userid, "190201", map[string]string{"id": "8"})
	if err != nil {
		return nil, err
	}
	v := new(models.DayworkReward)
	if err := json.Unmarshal(res, v); err != nil {
		return nil, err
	}
	return v, nil
}

func GetDaySignReward(serverUrl, sid, userid string) (*models.DayworkReward, error) {
	res, err := send(serverUrl, sid, userid, "190102", nil)
	if err != nil {
		return nil, err
	}
	v := new(models.DayworkReward)
	if err := json.Unmarshal(res, v); err != nil {
		return nil, err
	}
	return v, nil
}

func GetDayExp(serverUrl, sid, userid string) (*models.DayworkReward, error) {
	res, err := send(serverUrl, sid, userid, "1304", map[string]string{"id": "11001"})
	if err != nil {
		return nil, err
	}
	v := new(models.DayworkReward)
	if err := json.Unmarshal(res, v); err != nil {
		return nil, err
	}
	return v, nil
}

func GetDayMoney(serverUrl, sid, userid string) (*models.DayworkReward, error) {
	res, err := send(serverUrl, sid, userid, "1305", map[string]string{"id": "12001"})
	if err != nil {
		return nil, err
	}
	v := new(models.DayworkReward)
	if err := json.Unmarshal(res, v); err != nil {
		return nil, err
	}
	return v, nil
}
