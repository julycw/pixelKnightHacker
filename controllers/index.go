package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/julycw/pixelKnightHacker/command"
	"github.com/julycw/pixelKnightHacker/models"
	"log"
	"strconv"
	"time"
)

type IndexController struct {
	beego.Controller
}

func (c *IndexController) Prepare() {
	if c.GetSession("sid") == nil || c.GetSession("userid") == nil || c.GetSession("serverAddr") == nil {
		userId := c.GetString("u")
		sId := c.GetString("s")
		domain := c.GetString("d")
		var db_id int64

		if userId == "" || sId == "" || domain == "" {
			c.Redirect("login", 302)
		}

		serverAddr := "http://" + domain + "/service/main.ashx"

		if info, err := command.GetUserInfo(serverAddr, sId, userId); err != nil {
			c.Redirect("login", 302)
		} else {
			if info.Ret == 0 {
				userName := info.Data.UserName
				o := models.GetOrm()
				var result []orm.Params
				num, err := o.Raw("SELECT id FROM player WHERE Uid = ? AND domain = ?", userId, domain).Values(&result)
				if err == nil && num > 0 {
					if _, err := o.Raw("UPDATE player SET login_on = ?, login_times = login_times + 1 WHERE id = ?",
						time.Now(),
						result[0]["id"],
					).Exec(); err != nil {
						log.Println(err.Error())
					} else {
						id, _ := strconv.ParseInt(result[0]["id"].(string), 10, 64)
						db_id = id
					}
				} else {
					uid, _ := strconv.ParseInt(userId, 10, 64)
					player := models.Player{
						Uid:        int(uid),
						Sid:        sId,
						Domain:     domain,
						Name:       userName,
						LoginOn:    time.Now(),
						LoginTimes: 1,
					}
					if id, err := models.AddPlayer(player); err != nil {
						log.Println("Add user:", uid, sId, domain, userName, ", details:", err.Error())
					} else {
						db_id = id
					}
				}

				log := models.Log{
					PlayerId: db_id,
					Action:   "login",
				}
				models.AddLog(log)

				c.SetSession("userid", userId)
				c.SetSession("sid", sId)
				c.SetSession("serverAddr", serverAddr)
				c.SetSession("username", userName)
				c.SetSession("db_id", db_id)
			} else {
				c.Redirect("login", 302)
			}
		}
	}
}

func (c *IndexController) Get() {
	sid := c.GetSession("sid").(string)
	userid := c.GetSession("userid").(string)
	serverAddr := c.GetSession("serverAddr").(string)

	var v string
	if c.GetSession("version") == nil {
		if version, err := command.GetVersion(serverAddr, sid, userid); err != nil {
			c.Redirect("login", 302)
		} else {
			if version.Ret == 0 {
				v = version.Data.Version
			} else {
				c.Redirect("login", 302)
			}
		}
	}

	c.Data["UserName"] = c.GetSession("username")
	c.Data["Version"] = v

	c.TplNames = "index.tpl"
}
