package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Log struct {
	Id       int64
	PlayerId int64
	Action   string
	CreateOn time.Time
}

func AddLog(obj Log) (int64, error) {
	orm := GetOrm()
	obj.CreateOn = time.Now()
	id, err := orm.Insert(&obj)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func GetLog(id int64) (*Log, error) {
	orm := GetOrm()
	log := new(Log)
	if err := orm.QueryTable("log").Filter("Id", id).RelatedSel().One(log); err != nil {
		return nil, err
	}

	return log, nil
}

func GetAllLogs(cond *orm.Condition, pageIndex, pageSize int, order ...string) (*[]Log, int64, error) {
	orm := GetOrm()
	var logs *[]Log = new([]Log)
	total, err := orm.QueryTable("log").SetCond(cond).Count()
	if err != nil {
		return nil, 0, err
	}
	_, err = orm.QueryTable("log").SetCond(cond).Limit(pageSize, (pageIndex-1)*pageSize).OrderBy(order...).All(logs)
	if err != nil {
		return nil, 0, err
	}
	return logs, total, nil
}

func DeleteLog(id int64) error {
	orm := GetOrm()
	_, err := orm.QueryTable("log").Filter("Id", id).Delete()
	if err != nil {
		return err
	}
	return nil
}
