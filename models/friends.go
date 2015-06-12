package models

type FriendReward struct {
	Basic
	Data FriendRewardData `json:"data"`
}

type FriendRewardData struct {
	Rewards []Reward `json:"rewards"`
}

type FriendList struct {
	Basic
	Data FriendListData `json:"data"`
}

type FriendListData struct {
	List []FriendListItemData `json:"list"`
}

type FriendListItemData struct {
	Id       int    `json:"id"`
	UserName string `json:"username"`
	CanGift  bool   `json:"cangift"`
}

type FriendMessageList struct {
	Basic
	Data []FriendMessageItemData `json:"data"`
}

type FriendMessageItemData struct {
	SenderId     int    `json:"senderid"`
	SenderName   string `json:"sendername"`
	SenderFigure int    `json:"senderfigure"`
	Id           int    `json:"id"`
	GiftType     int    `json:"gifttype"`
	State        int    `json:"state"`
	Type         int    `json:"type"`
	IsNew        int    `json:"isnew"`
}
