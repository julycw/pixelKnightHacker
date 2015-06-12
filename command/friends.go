package command

import (
	"encoding/json"
	"github.com/julycw/pixelKnightHacker/models"
)

func AddFriend(serverUrl, sid, userid, name string) (*models.Basic, error) {
	res, err := send(serverUrl, sid, userid, "4009", map[string]string{"name": name})
	if err != nil {
		return nil, err
	}
	v := new(models.Basic)
	if err := json.Unmarshal(res, v); err != nil {
		return nil, err
	}
	return v, nil
}

func GetGift(serverUrl, sid, userid, id string) (*models.FriendReward, error) {
	res, err := send(serverUrl, sid, userid, "4006", map[string]string{"id": id})
	if err != nil {
		return nil, err
	}
	v := new(models.FriendReward)
	if err := json.Unmarshal(res, v); err != nil {
		return nil, err
	}
	return v, nil
}

func SendGift(serverUrl, sid, userid, id, gifttype string) (*models.FriendReward, error) {
	res, err := send(serverUrl, sid, userid, "4004", map[string]string{"id": id, "gifttype": gifttype})
	if err != nil {
		return nil, err
	}
	v := new(models.FriendReward)
	if err := json.Unmarshal(res, v); err != nil {
		return nil, err
	}
	return v, nil
}

func GetFriendList(serverUrl, sid, userid string) (*models.FriendList, error) {
	res, err := send(serverUrl, sid, userid, "4003", nil)
	if err != nil {
		return nil, err
	}
	v := new(models.FriendList)
	if err := json.Unmarshal(res, v); err != nil {
		return nil, err
	}
	return v, nil
}

func GetFriendMessage(serverUrl, sid, userid string) (*models.FriendMessageList, error) {
	res, err := send(serverUrl, sid, userid, "4005", nil)
	if err != nil {
		return nil, err
	}
	v := new(models.FriendMessageList)
	if err := json.Unmarshal(res, v); err != nil {
		return nil, err
	}
	return v, nil
}
