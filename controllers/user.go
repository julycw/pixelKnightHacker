package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/julycw/pixelKnightHacker/command"
	"github.com/julycw/pixelKnightHacker/models"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

var (
	sidExp    = regexp.MustCompile(`sid=(.*?)&`)
	uidExp    = regexp.MustCompile(`LoadGame\(0,(.*?)\);`)
	domainExp = regexp.MustCompile(`http://(.*?)/`)
)

type UserController struct {
	beego.Controller
}

// @router /login [get]
func (c *UserController) Get() {
	c.Data["message"] = c.GetSession("message")
	c.DelSession("message")
	c.TplNames = "login.tpl"
}

// @router /logout [get]
func (c *UserController) Logout() {
	if c.GetSession("db_id") != nil {
		id, _ := c.GetSession("db_id").(int64)
		log := models.Log{
			PlayerId: id,
			Action:   "logout",
		}
		models.AddLog(log)
		c.DelSession("sid")
		c.DelSession("userid")
		c.DelSession("serverAddr")
		c.DelSession("db_id")
	}
	c.Redirect("/login", 302)
}

// @router /login [post]
func (c *UserController) Post() {
	var userId string
	var sId string
	var serverAddr string
	var domain string
	var userName string
	var db_id int64

	userId = c.GetString("userid")
	sId = c.GetString("sid")

	if url := c.GetString("url"); url != "" {
		if matchs := domainExp.FindStringSubmatch(url); matchs != nil {
			domain = matchs[1]
		}
		if matchs := sidExp.FindStringSubmatch(url); matchs != nil {
			sId = matchs[1]
			// xxx/service/main.ashx?uid=2221700&serverid=1&ptime=1431053658&isadult=1&sign=c8fc39151f567d58063670f58910fa57&sid=rarqbuf4ntlsl2lzxq4wvu5h&areaid=1001
			if resp, err := http.Get(url); err == nil {
				body, _ := ioutil.ReadAll(resp.Body)
				resp.Body.Close()

				if matchs := uidExp.FindStringSubmatch(string(body)); matchs != nil {
					userId = matchs[1]
				}
			}
		}
	}
	serverAddr = "http://" + domain + "/service/main.ashx"
	msg, ok := func(serverAddr, userId, sId string) (string, bool) {
		if userId == "" {
			return "您输入的Url地址不正确", false
		}
		if sId == "" {
			return "您输入的Url地址不正确", false
		}

		if info, err := command.GetUserInfo(serverAddr, sId, userId); err != nil {
			return "网络故障: " + err.Error(), false
		} else {
			if info.Ret == 0 {
				userName = info.Data.UserName
				c.SetSession("username", userName)
			} else {
				return info.Message, false
			}
		}
		return "", true
	}(serverAddr, userId, sId)

	if ok {
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

		c.SetSession("sid", sId)
		c.SetSession("userid", userId)
		c.SetSession("serverAddr", serverAddr)
		c.SetSession("username", userName)
		c.SetSession("db_id", db_id)
		c.Redirect(fmt.Sprintf("/?u=%s&s=%s&d=%s", userId, sId, domain), 302)
	} else {
		c.SetSession("message", msg)
		c.Redirect("login", 302)
	}
}
