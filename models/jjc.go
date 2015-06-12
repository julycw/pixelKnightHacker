package models

type JjcHeroList struct {
	Basic
	Data HeroListData `json:"data"`
}

type HeroListData struct {
	HeroList   []HeroData `json:"list"`
	Time       int        `json:"time"`
	ArenaCount int        `json:"arenacount"`
	Box1IsGet  int        `json:"arenabox1isget"`
	Box2IsGet  int        `json:"arenabox2isget"`
	Box3IsGet  int        `json:"arenabox3isget"`
}

type HeroData struct {
	Id          int    `json:"id"`
	UserId      int    `json:"userid"`
	UserName    string `json:"username"`
	BattleValue int    `json:"battlevalue"`
	IsWin       int    `json:"iswin"`
}

type PKResult struct {
	Basic
	Data PKResultData `json:"data"`
}

type PKResultData struct {
	IsWin bool `json:"iswin"`
}

type JjcReward struct {
	Basic
}
