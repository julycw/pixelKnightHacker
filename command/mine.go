package command

import (
	"encoding/json"
	"github.com/julycw/pixelKnightHacker/models"
)

func GetMineCount(serverUrl, sid, userid string) (*models.MineInfo, error) {
	res, err := send(serverUrl, sid, userid, "2003", nil)
	if err != nil {
		return nil, err
	}
	v := new(models.MineInfo)
	if err := json.Unmarshal(res, v); err != nil {
		return nil, err
	}
	return v, nil
}

func GetMineQueue(serverUrl, sid, userid string) (*models.MineQueue, error) {
	res, err := send(serverUrl, sid, userid, "2016", nil)
	if err != nil {
		return nil, err
	}
	v := new(models.MineQueue)
	if err := json.Unmarshal(res, v); err != nil {
		return nil, err
	}
	return v, nil
}

func GetMineQueueReward(serverUrl, sid, userid, id string) (*models.MineQueueReward, error) {
	res, err := send(serverUrl, sid, userid, "2018", map[string]string{
		"id": id,
	})
	if err != nil {
		return nil, err
	}
	v := new(models.MineQueueReward)
	if err := json.Unmarshal(res, v); err != nil {
		return nil, err
	}
	return v, nil
}

// id: 10001~10012
func AddMineQueue(serverUrl, sid, userid, id string) (*models.AddMineQueueResult, error) {
	res, err := send(serverUrl, sid, userid, "2017", map[string]string{
		"id": id,
	})
	if err != nil {
		return nil, err
	}
	v := new(models.AddMineQueueResult)
	if err := json.Unmarshal(res, v); err != nil {
		return nil, err
	}
	return v, nil
}
