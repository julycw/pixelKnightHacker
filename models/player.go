package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Player struct {
	Id         int64
	Uid        int
	Sid        string
	Domain     string
	Name       string
	CreateOn   time.Time
	LoginOn    time.Time
	LoginTimes int
}

func AddPlayer(obj Player) (int64, error) {
	orm := GetOrm()
	obj.CreateOn = time.Now()
	id, err := orm.Insert(&obj)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func GetPlayer(id int) (*Player, error) {
	orm := GetOrm()
	player := new(Player)
	if err := orm.QueryTable("player").Filter("Id", id).RelatedSel().One(player); err != nil {
		return nil, err
	}

	return player, nil
}

func GetAllPlayers(cond *orm.Condition, pageIndex, pageSize int, order ...string) (*[]Player, int64, error) {
	orm := GetOrm()
	var players *[]Player = new([]Player)
	total, err := orm.QueryTable("player").SetCond(cond).Count()
	if err != nil {
		return nil, 0, err
	}
	_, err = orm.QueryTable("player").SetCond(cond).Limit(pageSize, (pageIndex-1)*pageSize).OrderBy(order...).All(players)
	if err != nil {
		return nil, 0, err
	}
	return players, total, nil
}

func UpdatePlayer(id int, update *Player) (*Player, error) {
	if obj, err := GetPlayer(id); err == nil {
		if update.Sid != "" {
			obj.Sid = update.Sid
		}

		orm := GetOrm()
		_, err := orm.Update(obj)
		if err != nil {
			return nil, err
		}
		return obj, nil
	} else {
		return nil, err
	}
}

func DeletePlayer(id int) error {
	orm := GetOrm()
	_, err := orm.QueryTable("player").Filter("Id", id).Delete()
	if err != nil {
		return err
	}
	return nil
}
