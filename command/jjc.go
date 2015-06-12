package command

import (
	"encoding/json"
	"github.com/julycw/pixelKnightHacker/models"
)

func GetJjcList(serverUrl, sid, userid string) (*models.JjcHeroList, error) {
	res, err := send(serverUrl, sid, userid, "6001", nil)
	if err != nil {
		return nil, err
	}
	v := new(models.JjcHeroList)
	if err := json.Unmarshal(res, v); err != nil {
		return nil, err
	}
	return v, nil
}

func GetJjcReward(serverUrl, sid, userid, typeid string) (*models.JjcReward, error) {
	res, err := send(serverUrl, sid, userid, "6006", map[string]string{
		"type": typeid,
	})
	if err != nil {
		return nil, err
	}
	v := new(models.JjcReward)
	if err := json.Unmarshal(res, v); err != nil {
		return nil, err
	}
	return v, nil
}

func PKWithHero(serverUrl, sid, userid, heroid string) (*models.PKResult, error) {
	res, err := send(serverUrl, sid, userid, "6007", map[string]string{
		"id": heroid,
	})
	if err != nil {
		return nil, err
	}
	v := new(models.PKResult)
	if err := json.Unmarshal(res, v); err != nil {
		return nil, err
	}
	return v, nil
}
