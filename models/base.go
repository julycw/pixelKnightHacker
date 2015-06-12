package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"net/url"
	"time"
)

var (
	dbName     string
	dbHost     string
	dbPort     string
	dbUser     string
	dbPassword string
)

type Basic struct {
	Ret     int    `json:"ret"`
	Message string `json:"msg"`
	Time    string `json:"time"`
}

func New(ret int, msg string) *Basic {
	return &Basic{
		Ret:     ret,
		Message: msg,
		Time:    time.Now().Format("2006-01-02 15:04:05"),
	}
}

type Reward struct {
	Id       int `json:"id"`
	Type     int `json:"type"`
	Count    int `json:"count"`
	RewardId int `json:"rewardid"`
}

func NewSuccess(msg string) *Basic {
	return New(0, msg)
}

func NewFailed(msg string) *Basic {
	return New(1, msg)
}

func GetOrm() orm.Ormer {
	o := orm.NewOrm()
	o.Using("default")
	return o
}

func init() {
	dbUser = beego.AppConfig.String("dbuser")
	dbPassword = beego.AppConfig.String("dbpass")
	dbHost = beego.AppConfig.String("dbhost")
	dbPort = beego.AppConfig.String("dbport")
	dbName = beego.AppConfig.String("dbname")
	orm.DefaultTimeLoc = time.Local
	// 需要在init中注册定义的model
	orm.RegisterModel(
		new(Player),
		new(Log),
	)
	maxIdle := 10 //(可选)  设置最大空闲连接
	maxConn := 50 //(可选)  设置最大数据库连接 (go >= 1.2)

	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		dbUser,
		dbPassword,
		dbHost,
		dbPort,
		dbName,
	) + "&loc=" + url.QueryEscape("Local")

	orm.RegisterDataBase("default", "mysql",
		connStr,
		maxIdle,
		maxConn,
	)
}
