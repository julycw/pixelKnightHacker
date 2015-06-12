package command

import (
	"encoding/json"
	"github.com/julycw/pixelKnightHacker/models"
)

func GetUserInfo(serverUrl, sid, userid string) (*models.UserInfo, error) {
	res, err := send(serverUrl, sid, userid, "1001", nil)
	if err != nil {
		return nil, err
	}
	v := new(models.UserInfo)
	if err := json.Unmarshal(res, v); err != nil {
		return nil, err
	}
	return v, nil
}

func GetVersion(serverUrl, sid, userid string) (*models.Version, error) {
	res, err := send(serverUrl, sid, userid, "1013", nil)
	if err != nil {
		return nil, err
	}
	v := new(models.Version)
	if err := json.Unmarshal(res, v); err != nil {
		return nil, err
	}
	return v, nil
}
