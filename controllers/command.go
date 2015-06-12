package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/julycw/pixelKnightHacker/command"
	"github.com/julycw/pixelKnightHacker/models"
	"log"
	"strconv"
	"time"
)

type CommandController struct {
	beego.Controller
}

func (c *CommandController) responseJsonFailed(msg string) {
	c.Data["json"] = models.NewFailed(msg)
	c.ServeJson()
	c.StopRun()
}

func (c *CommandController) responseJsonSuccess(msg string) {
	c.Data["json"] = models.NewSuccess(msg)
	c.ServeJson()
	c.StopRun()
}

func (c *CommandController) Prepare() {
	if c.IsAjax() {
		if c.GetSession("db_id") == nil || c.GetSession("sid") == nil || c.GetSession("userid") == nil {
			c.responseJsonFailed("登录...超时!")
		}

		if action := c.GetString("action"); action != "" {
			log := models.Log{
				PlayerId: c.GetSession("db_id").(int64),
				Action:   action,
			}
			models.AddLog(log)
		}

	} else {
		c.Redirect("login", 302)
	}
}

// @router /command/mineQueue [post]
func (c *CommandController) MineQueue() {
	sid := c.GetSession("sid").(string)
	userid := c.GetSession("userid").(string)
	serverAddr := c.GetSession("serverAddr").(string)
	mineType := c.GetString("mineType", "1")
	var mineId string
	switch mineType {
	case "1":
		mineId = "10001"
	case "2":
		mineId = "10004"
	case "3":
		mineId = "10007"
	case "4":
		mineId = "10010"
	default:
		mineId = "10001"
	}

	rewardCount := 0
	addCount := 0
	//首先获取矿队列表
	result, err := command.GetMineQueue(serverAddr, sid, userid)
	if err != nil {
		c.responseJsonFailed(err.Error())
	} else {
		if result.Ret == 0 {
			for _, v := range result.Data {
				if v.Sid != 0 {
					if v.Time == 0 {
						//收货
						result, err := command.GetMineQueueReward(serverAddr, sid, userid, strconv.Itoa(v.Id))
						if err == nil && result.Ret == 0 {
							rewardCount++
							//收完了必然有空位, 添加队列
							time.Sleep(500 * time.Millisecond)
							result, err := command.AddMineQueue(serverAddr, sid, userid, mineId)
							if err != nil {
								c.responseJsonFailed(err.Error())
							} else {
								if result.Ret == 0 {
									addCount++
								} else {
									c.responseJsonFailed(result.Message)
								}
							}
						}
					}
				} else {
					result, err := command.AddMineQueue(serverAddr, sid, userid, mineId)
					if err != nil {
						c.responseJsonFailed(err.Error())
					} else {
						if result.Ret == 0 {
							addCount++
						} else {
							c.responseJsonFailed(result.Message)
						}
					}
				}
			}
		} else {
			c.responseJsonFailed(result.Message)
		}
	}

	c.responseJsonSuccess(fmt.Sprintf("添加%d个挖矿队伍, 获得%d个队伍成果!", addCount, rewardCount))
}

// @router /command/jjc [post]
func (c *CommandController) Jjc() {
	var chance int
	var winCount int
	var newFriend int
	var rewardResult string
	sid := c.GetSession("sid").(string)
	userid := c.GetSession("userid").(string)
	serverAddr := c.GetSession("serverAddr").(string)
	// 获取竞技场对手列表
	list, err := command.GetJjcList(serverAddr, sid, userid)
	if err != nil {
		log.Println(err.Error())
	} else {
		if list.Ret == 0 {
			safeReturn := 20
			chance = 10 - list.Data.ArenaCount
			winCount = 0
			func() {
				for {
					for i, v := range list.Data.HeroList {
						// 机会用完退出
						if chance <= 0 {
							return
						}
						// 判断是否已经Win过
						if v.IsWin == 0 {
							// 加好友
							if v.UserId > 0 {
								if ret, err := command.AddFriend(serverAddr, sid, userid, v.UserName); err == nil && ret.Ret == 0 {
									newFriend++
								}
							}
							//进行PK
							if result, err := command.PKWithHero(serverAddr, sid, userid, strconv.Itoa(v.Id)); err != nil {
								log.Println(err.Error())
								return
							} else {
								chance--
								if result.Data.IsWin {
									list.Data.HeroList[i].IsWin = 2 //2代表这次打赢了,且胜利已经被记录
									winCount++
								}
							}
							time.Sleep(250 * time.Millisecond)
						} else if v.IsWin == 1 {
							winCount++
						}
						// 完成9胜后退出
						if winCount == 9 {
							return
						}
					}
					if safeReturn <= 0 {
						return
					}
					safeReturn--
				}
			}()
			if winCount >= 1 && list.Data.Box1IsGet == 0 {
				result, err := command.GetJjcReward(serverAddr, sid, userid, "1")
				if err == nil && result.Ret == 0 {
					rewardResult += "1胜宝箱领取成功!"
				}
			}
			if winCount >= 4 && list.Data.Box2IsGet == 0 {
				result, err := command.GetJjcReward(serverAddr, sid, userid, "4")
				if err == nil && result.Ret == 0 {
					rewardResult += "4胜宝箱领取成功!"
				}
			}
			if winCount >= 9 && list.Data.Box3IsGet == 0 {
				result, err := command.GetJjcReward(serverAddr, sid, userid, "9")
				if err == nil && result.Ret == 0 {
					rewardResult += "9胜宝箱领取成功!"
				}
			}

			if rewardResult == "" {
				rewardResult = "无!"
			}
		} else {
			c.responseJsonFailed(list.Message)
		}
	}
	c.responseJsonSuccess(fmt.Sprintf("剩余%d次机会, 获得%d胜, 奖励情况:%s 添加%d个新的好友!", chance, winCount, rewardResult, newFriend))
}

// @router /command/explore [post]
func (c *CommandController) Explore() {
	sid := c.GetSession("sid").(string)
	userid := c.GetSession("userid").(string)
	serverAddr := c.GetSession("serverAddr").(string)

	getRewardCount := 0
	getExp := 0
	getMoney := 0
	//获取战利品
	if rewardCount, err := command.GetRewardCount(serverAddr, sid, userid); err == nil {
		if rewardCount.Ret == 0 {
			if rewards, err := command.GetRewards(serverAddr, sid, userid); err == nil {
				getRewardCount = len(rewards.Data.Rewards)
				if rwd, err := command.UseAllExpItem(serverAddr, sid, userid); err == nil && rwd.Data.Rewards != nil && len(rwd.Data.Rewards) > 0 {
					getExp = rwd.Data.Rewards[0].Count
				}
				if rwd, err := command.UseAllMoneyItem(serverAddr, sid, userid); err == nil && rwd.Data.Rewards != nil && len(rwd.Data.Rewards) > 0 {
					getMoney = rwd.Data.Rewards[0].Count
				}
			} else {
				c.responseJsonFailed(err.Error())
			}
		} else {
			c.responseJsonFailed(rewardCount.Message)
		}
	} else {
		c.responseJsonFailed(err.Error())
	}

	responseMessage := ""
	responseMessage += fmt.Sprintf("收获%d个战利品, 获得%d金币, %d经验!", getRewardCount, getMoney, getExp)

	c.responseJsonSuccess(responseMessage)
}

// @router /command/friends [post]
func (c *CommandController) Friends() {
	sid := c.GetSession("sid").(string)
	userid := c.GetSession("userid").(string)
	serverAddr := c.GetSession("serverAddr").(string)
	givenCount := 0
	sentCount := 0
	// 送礼
	if friendsList, err := command.GetFriendList(serverAddr, sid, userid); err == nil {
		if friendsList.Ret == 0 {
			for _, v := range friendsList.Data.List {
				if v.CanGift {
					ret, err := command.SendGift(serverAddr, sid, userid, strconv.Itoa(v.Id), "3")
					if err == nil && ret.Ret == 0 {
						sentCount++
					}
				}
			}
		} else {
			c.responseJsonFailed(friendsList.Message)
		}
	} else {
		c.responseJsonFailed(err.Error())
	}

	// 收礼
	if msgList, err := command.GetFriendMessage(serverAddr, sid, userid); err == nil {
		if msgList.Ret == 0 {
			for _, v := range msgList.Data {
				if v.GiftType > 0 && v.State == 0 {
					ret, err := command.GetGift(serverAddr, sid, userid, strconv.Itoa(v.Id))
					if err == nil && ret.Ret == 0 {
						givenCount++
					}
				}
			}
		} else {
			c.responseJsonFailed(msgList.Message)
		}
	} else {
		c.responseJsonFailed(err.Error())
	}

	c.responseJsonSuccess(fmt.Sprintf("送出%d个礼物, 收到%d个礼物!", sentCount, givenCount))
}

// @router /command/daywork [post]
func (c *CommandController) Daywork() {
	sid := c.GetSession("sid").(string)
	userid := c.GetSession("userid").(string)
	serverAddr := c.GetSession("serverAddr").(string)
	msg := ""
	func() {
		//签到
		if ret, err := command.GetDaySignReward(serverAddr, sid, userid); err == nil {
			if ret.Ret == 0 {
				msg += "签到完成!"
			} else {
				msg += ret.Message + "!"
			}
		} else {
			c.responseJsonFailed(err.Error())
		}
		//炼金
		if ret, err := command.GetDayMoney(serverAddr, sid, userid); err == nil {
			if ret.Ret == 0 && ret.Data.Rewards != nil && len(ret.Data.Rewards) > 0 {
				msg += fmt.Sprintf("炼金获得%d金币!", ret.Data.Rewards[0].Count)
			} else {
				msg += ret.Message + "!"
			}
		} else {
			c.responseJsonFailed(err.Error())
		}
		//祈祷
		if ret, err := command.GetDayExp(serverAddr, sid, userid); err == nil {
			if ret.Ret == 0 && ret.Data.Rewards != nil && len(ret.Data.Rewards) > 0 {
				msg += fmt.Sprintf("祈祷获得%d经验!", ret.Data.Rewards[0].Count)
			} else {
				msg += ret.Message + "!"
			}
		} else {
			c.responseJsonFailed(err.Error())
		}
		//月卡
		if ret, err := command.GetMonthCardReward1(serverAddr, sid, userid); err == nil {
			if ret.Ret == 0 {
				msg += "领取初级月卡成功!"
			} else {
				msg += ret.Message + "!"
			}
		} else {
			c.responseJsonFailed(err.Error())
		}
		if ret, err := command.GetMonthCardReward2(serverAddr, sid, userid); err == nil {
			if ret.Ret == 0 {
				msg += "领取中级月卡成功!"
			} else {
				msg += ret.Message + "!"
			}
		} else {
			c.responseJsonFailed(err.Error())
		}
	}()

	c.responseJsonSuccess(msg)
}

// @router /command/heartBeat [post]
func (c *CommandController) HeartBeat() {
	sid := c.GetSession("sid").(string)
	userid := c.GetSession("userid").(string)
	serverAddr := c.GetSession("serverAddr").(string)

	mineCount := 0
	if result, err := command.GetMineCount(serverAddr, sid, userid); err == nil {
		if result.Ret == 0 {
			mineCount = result.Data.Szg
		} else {
			c.responseJsonFailed(result.Message)
		}
	} else {
		c.responseJsonFailed(err.Error())
	}

	// 刷新一下饥饿值
	command.RefreshHungry(serverAddr, sid, userid)
	// 吃东西
	eatReward := 0
	leftBzd := 0
	if bzd, err := command.GetExploreInfo(serverAddr, sid, userid); err == nil {
		leftBzd = bzd.Data.Bzd
		if leftBzd < 1000 {
			foodList, err := command.GetFoodList(serverAddr, sid, userid)
			if err != nil {
				c.responseJsonFailed(err.Error())
			} else {
				for _, v := range foodList.Data {
					if v.Count > 0 {
						eatResult, err := command.EatFood(serverAddr, sid, userid, strconv.Itoa(v.FoodId)) /* 2是梅子饭团*/
						if err == nil && eatResult.Ret == 0 && eatResult.Data.Rewards != nil && len(eatResult.Data.Rewards) > 0 {
							eatReward += eatResult.Data.Rewards[0].Count
							// 只吃一个好, 吃多伤身体
							break
						}
					}
				}
			}
		}
	} else {
		c.responseJsonFailed(err.Error())
	}
	responseMessage := fmt.Sprintf("冒险体力:%d, 挖矿体力:%d", leftBzd, mineCount)
	if eatReward > 0 {
		responseMessage += fmt.Sprintf(", 吃东西补充%d点体力!", eatReward)
	}

	c.responseJsonSuccess(responseMessage)
}
