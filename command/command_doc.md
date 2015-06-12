# API action表
## game.aspx get
### request
	uid:2221700
	serverid:1
	ptime:1430738048
	isadult:1
	sign:cdb4f92a4f333744abfa15cc0af90c7d
	sid:jpmfau3finc4toqws3ayu12z
	areaid:1001
### response
	游戏主界面的网页

## main.ashx post

### 基础request
	sid:xxxxxx		//sid，每次登陆会改变
	t:xxx			//target?用来分辨动作
	userid:24795	//用户id，不变


### t = 1001 获取角色信息？
#### request 
	略
#### response
	{"ret":0,"data":{"username":"七月汪汪","sex":1,"figure":99000002,"jb":10593079,"xz":199,"yjjc":12270,"blhj":1,"zrhj":2,"hhhj":0,"hlhj":6,"mazeboxmax":3,"lhsp":6153,"newmsgcount":0,"blue_vip_level":0,"isqh":false,"lshd":0,"ishz":1,"sortid":0}}

### t = 1007 获取背包物品列表
#### request
	略
#### response
	{"ret":0,"data":{"list":[{"id":3851067,"itemid":1000002,"isnew":0,"value":0},{"id":3943340,"itemid":1000002,"isnew":0,"value":0},{"id":3943349,"itemid":1000001,"isnew":0,"value":0},{"id":3943358,"itemid":1000001,"isnew":0,"value":0},{"id":3943365,"itemid":1000001,"isnew":0,"value":0},{"id":3943370,"itemid":1000002,"isnew":0,"value":0},{"id":3943376,"itemid":1000002,"isnew":0,"value":0},{"id":3943378,"itemid":1000001,"isnew":0,"value":0},{"id":3943386,"itemid":1000001,"isnew":0,"value":0},{"id":3943391,"itemid":1000001,"isnew":0,"value":0},{"id":3961221,"itemid":2000006,"isnew":0,"value":0},{"id":3962585,"itemid":8000003,"isnew":1,"value":0}],"packagelv":13}}


### t = 1009 使用道具
#### request
	id:3962585  //物品id
#### response
	{"ret":0,"data":{"rewards":[{"type":23,"count":44,"rewardid":19000097}]}}
#### 备注
	type:6 	经验
	type:23 金币
	count 	数量

### t = 1013 获取版本号
#### request
	略
#### response
	{"ret":0,"data":{"version":"2.2.10"}}


### t = 1302 获取每日祈祷级别
#### request
	略
#### response
	{"ret":0,"data":{"lv":17,"time":0}}
	


### t = 1303 获取每日炼金级别
#### request
	略
#### response
	{"ret":0,"data":{"lv":17,"time":0}}
	

### t = 1305 每日祈祷
#### request
	id:11001
	sid:ylzxw2plsv55kcnicbp24cxq
	t:1305
	userid:24795
#### response
	{"ret":0,"data":{"rewards":[{"type":6,"rewardid":21000001,"count":30600}]}}


### t = 1305 每日炼金
#### request
	id:12001
	sid:ylzxw2plsv55kcnicbp24cxq
	t:1305
	userid:24795
#### response
	{"ret":0,"data":{"rewards":[{"type":23,"count":34000,"rewardid":21000004}]}}


### t = 1801 查看商会任务
#### request
	略
#### response
	{"ret":0,"data":{"list":[{"id":34105,"orderid":1000013,"state":1},{"id":34106,"orderid":1000016,"state":1},{"id":34107,"orderid":1000011,"state":1}],"rewardid":100004,"orderiscomplete":1,"time":6279}}


### t = 2003 获取矿场信息
#### request
	略
#### response
	{"ret":0,"data":{"time":110,"szg":28,"szglv":7,"wjxl":4,"monstercount":0,"monsterlv":48,"yanmoLv":1,"time1":0,"szgmax":45,"slszg":4,"lg":10,"hslg":0,"isreset":false,"orebag1":0,"orebag2":0,"orebag3":0}}


### t = 2004 攻击魔像
#### request
	x:8								//坐标x
	y:99							//坐标y
	t:2004
	time:10715950.8309298679232597
	sid:xxxzpi3mxxz21wmzquznlzlc
	userid:24795
	cmd:exploit
#### response
	{"ret":0,"data":{"x":8,"y":99,"type":2,"rewards":[{"type":35,"count":1,"rewardid":25000061}],"report":{"leftMG":{"name":"七月汪汪","merList":[99000002,19000036,19000006,18000013,18000002,18000042,18000022,18000022,19000017,19000017,18000030,18000023,18000026,17000004,18000002,18000009,18000027,18000041,18000010,11000032],"ATK":513891,"maxATK":513891,"skills":[1100009,1100032,1200030,1200041,1200026,1900015,1200053,1100006,1200026,1200028,1100018,1100008,1100013],"skillsLv":1,"firstStrike":31,"defense":44,"dodge":9,"wz":0},"rightMG":{"name":"魔铁巨像lv43","merList":[20000002],"ATK":640000,"maxATK":640000,"skills":[1200035],"skillsLv":1,"firstStrike":10,"defense":45,"dodge":30,"wz":0},"skillList":[1100009,1000001,1100032,1000001,1100008,1000001,1200028],"dodgeList":[false,false,false,false,false,false,false],"damageValue":[[0,-183405],[-120597,0],[0,-225412],[-104943,0],[0,-156574],[-94069,0],[0,-130634]],"xly":[],"skip":1},"iswin":true,"nodeList":[{"x":12,"y":99,"type":3},{"x":8,"y":95,"type":3},{"x":11,"y":98,"type":3},{"x":9,"y":96,"type":3},{"x":10,"y":97,"type":41}]}}
#### 备注
	rewards.type 35 魔像


### t = 2007 捡取矿物
#### request 
	x:10
	y:99
	t:2007
	time:21773510.5510969925671816
	sid:xxxzpi3mxxz21wmzquznlzlc
	userid:24795
	cmd:charge
#### response
	略


### t = 2016 获取挖矿队列列表
#### request
	略
#### response
	{"ret":0,"data":[{"id":4692,"sid":10002,"time":5764},{"id":6823,"sid":0},{"id":4693,"sid":0}]}


### t = 2017 添加挖矿队列
#### request
	id:10001(~10012)
	sid:jtlz21qhzyuulg0nxmfomnsv
	t:2017
	userid:24795
#### response
	略


### t = 2018 收取挖矿队成果
#### request
	id:4693
	sid:u3iudktskhvjkvbe5v4rcua4
	t:2018
	userid:24795
#### response
	{"ret":0,"data":{"rewards":[{"type":9,"count":12,"rewardid":40000003},{"type":8,"count":32,"rewardid":40000001}]}}


### t = 4003 查看好友列表
#### request
	略
#### response
	{"ret":0,"data":{"list":[{"id":4523,"sex":0,"figure":19000007,"username":"银魂开发部","battlevalue":2150000,"mlv":0,"canAttack":true,"cangift":false,"areaid":1001,"ishx":0,"ftype":0,"tjcount":0,"qycount":0,"v1":0,"yellow_vip_level":0,"qqname":""},{"id":22453,"sex":1,"figure":99000001,"username":"腻夏苏伦莫立","battlevalue":718202,"mlv":64,"canAttack":true,"cangift":false,"areaid":1001,"ishx":1,"ftype":1,"tjcount":95,"qycount":0,"v1":0,"yellow_vip_level":0,"qqname":"腻夏苏伦莫立"},{"id":12184,"sex":1,"figure":99000001,"username":"哈啊","battlevalue":598050,"mlv":55,"canAttack":true,"cangift":false,"areaid":1001,"ishx":1,"ftype":1,"tjcount":89,"qycount":0,"v1":0,"yellow_vip_level":0,"qqname":"哈啊"},{"id":5382,"sex":1,"figure":99000002,"username":"玩个蛋蛋","battlevalue":551917,"mlv":39,"canAttack":true,"cangift":false,"areaid":1001,"ishx":1,"ftype":1,"tjcount":84,"qycount":0,"v1":0,"yellow_vip_level":0,"qqname":"玩个蛋蛋"},{"id":6836,"sex":1,"figure":99000003,"username":"HZH","battlevalue":537529,"mlv":39,"canAttack":true,"cangift":false,"areaid":1001,"ishx":1,"ftype":1,"tjcount":96,"qycount":0,"v1":0,"yellow_vip_level":0,"qqname":"HZH"},{"id":23442,"sex":1,"figure":99000001,"username":"俄那麦科","battlevalue":528006,"mlv":38,"canAttack":true,"cangift":false,"areaid":1001,"ishx":1,"ftype":1,"tjcount":108,"qycount":0,"v1":0,"yellow_vip_level":0,"qqname":"俄那麦科"},{"id":19705,"sex":1,"figure":18000002,"username":"瑟密兰内","battlevalue":496771,"mlv":40,"canAttack":true,"cangift":true,"areaid":1001,"ishx":1,"ftype":1,"tjcount":91,"qycount":0,"v1":0,"yellow_vip_level":0,"qqname":"瑟密兰内"},{"id":15001,"sex":1,"figure":99000003,"username":"名柴匹良","battlevalue":490502,"mlv":39,"canAttack":true,"cangift":false,"areaid":1001,"ishx":1,"ftype":1,"tjcount":88,"qycount":0,"v1":0,"yellow_vip_level":0,"qqname":"名柴匹良"},{"id":10782,"sex":1,"figure":99000002,"username":"何破龄难","battlevalue":480858,"mlv":39,"canAttack":true,"cangift":true,"areaid":1001,"ishx":1,"ftype":1,"tjcount":80,"qycount":0,"v1":0,"yellow_vip_level":0,"qqname":"何破龄难"},{"id":19930,"sex":1,"figure":99000002,"username":"诡丝","battlevalue":478041,"mlv":35,"canAttack":true,"cangift":true,"areaid":1001,"ishx":1,"ftype":1,"tjcount":96,"qycount":0,"v1":0,"yellow_vip_level":0,"qqname":"诡丝"},{"id":22009,"sex":1,"figure":99000001,"username":"较黄的圣骑士","battlevalue":473936,"mlv":38,"canAttack":true,"cangift":false,"areaid":1001,"ishx":1,"ftype":1,"tjcount":112,"qycount":0,"v1":0,"yellow_vip_level":0,"qqname":"较黄的圣骑士"},{"id":10671,"sex":2,"figure":99000005,"username":"寿杰奈茨麦叙","battlevalue":464025,"mlv":47,"canAttack":true,"cangift":true,"areaid":1001,"ishx":1,"ftype":1,"tjcount":81,"qycount":0,"v1":0,"yellow_vip_level":0,"qqname":"寿杰奈茨麦叙"},{"id":17274,"sex":1,"figure":99000001,"username":"厄烈保沙罗索","battlevalue":458467,"mlv":41,"canAttack":true,"cangift":true,"areaid":1001,"ishx":1,"ftype":1,"tjcount":82,"qycount":0,"v1":0,"yellow_vip_level":0,"qqname":"厄烈保沙罗索"},{"id":11529,"sex":1,"figure":99000001,"username":"洛各山芙","battlevalue":449116,"mlv":43,"canAttack":true,"cangift":false,"areaid":1001,"ishx":1,"ftype":1,"tjcount":95,"qycount":0,"v1":0,"yellow_vip_level":0,"qqname":"洛各山芙"},{"id":7899,"sex":1,"figure":99000001,"username":"连蒙佐生","battlevalue":434622,"mlv":39,"canAttack":true,"cangift":true,"areaid":1001,"ishx":1,"ftype":1,"tjcount":88,"qycount":0,"v1":0,"yellow_vip_level":0,"qqname":"连蒙佐生"},{"id":6066,"sex":1,"figure":19000033,"username":"丁佛","battlevalue":433482,"mlv":41,"canAttack":true,"cangift":true,"areaid":1001,"ishx":1,"ftype":1,"tjcount":132,"qycount":0,"v1":0,"yellow_vip_level":0,"qqname":"丁佛"},{"id":5700,"sex":1,"figure":99000002,"username":"孤高的橘子","battlevalue":425469,"mlv":37,"canAttack":true,"cangift":true,"areaid":1001,"ishx":1,"ftype":1,"tjcount":102,"qycount":0,"v1":0,"yellow_vip_level":0,"qqname":"孤高的橘子"},{"id":19946,"sex":1,"figure":99000001,"username":"PanDa.WIll","battlevalue":419123,"mlv":41,"canAttack":true,"cangift":true,"areaid":1001,"ishx":1,"ftype":1,"tjcount":113,"qycount":0,"v1":0,"yellow_vip_level":0,"qqname":"PanDa.WIll"},{"id":10453,"sex":2,"figure":99000006,"username":"世罗切托芬","battlevalue":408906,"mlv":30,"canAttack":true,"cangift":true,"areaid":1001,"ishx":1,"ftype":1,"tjcount":75,"qycount":0,"v1":0,"yellow_vip_level":0,"qqname":"世罗切托芬"},{"id":7346,"sex":1,"figure":99000002,"username":"關七人事","battlevalue":406820,"mlv":41,"canAttack":true,"cangift":true,"areaid":1001,"ishx":1,"ftype":1,"tjcount":100,"qycount":0,"v1":0,"yellow_vip_level":0,"qqname":"關七人事"},{"id":33468,"sex":2,"figure":99000004,"username":"泡只小豚豚","battlevalue":401797,"mlv":50,"canAttack":true,"cangift":true,"areaid":1001,"ishx":1,"ftype":1,"tjcount":93,"qycount":0,"v1":0,"yellow_vip_level":0,"qqname":"泡只小豚豚"},{"id":10467,"sex":1,"figure":99000001,"username":"格凯书柔","battlevalue":400445,"mlv":42,"canAttack":true,"cangift":true,"areaid":1001,"ishx":1,"ftype":1,"tjcount":101,"qycount":0,"v1":0,"yellow_vip_level":0,"qqname":"格凯书柔"},{"id":7394,"sex":2,"figure":18000013,"username":"卡到荼蘼","battlevalue":393994,"mlv":33,"canAttack":true,"cangift":false,"areaid":1001,"ishx":1,"ftype":1,"tjcount":85,"qycount":0,"v1":0,"yellow_vip_level":0,"qqname":"卡到荼蘼"},{"id":22781,"sex":1,"figure":99000003,"username":"恩雷达杰","battlevalue":390985,"mlv":29,"canAttack":true,"cangift":false,"areaid":1001,"ishx":1,"ftype":1,"tjcount":125,"qycount":0,"v1":0,"yellow_vip_level":0,"qqname":"恩雷达杰"},{"id":13931,"sex":2,"figure":19000008,"username":"射命丸文","battlevalue":390502,"mlv":35,"canAttack":true,"cangift":false,"areaid":1001,"ishx":1,"ftype":1,"tjcount":73,"qycount":0,"v1":0,"yellow_vip_level":0,"qqname":"射命丸文"},{"id":7417,"sex":1,"figure":99000002,"username":"死宅医生","battlevalue":379781,"mlv":39,"canAttack":true,"cangift":false,"areaid":1001,"ishx":1,"ftype":1,"tjcount":85,"qycount":0,"v1":0,"yellow_vip_level":0,"qqname":"死宅医生"},{"id":8650,"sex":1,"figure":19000008,"username":"诺琳肯武博尼","battlevalue":376367,"mlv":36,"canAttack":true,"cangift":true,"areaid":1001,"ishx":1,"ftype":1,"tjcount":78,"qycount":0,"v1":0,"yellow_vip_level":0,"qqname":"诺琳肯武博尼"},{"id":6664,"sex":1,"figure":99000002,"username":"普尔胡巴","battlevalue":355879,"mlv":38,"canAttack":true,"cangift":true,"areaid":1001,"ishx":1,"ftype":1,"tjcount":67,"qycount":0,"v1":0,"yellow_vip_level":0,"qqname":"普尔胡巴"},{"id":12390,"sex":1,"figure":19000008,"username":"马夏坦布译扎","battlevalue":339888,"mlv":36,"canAttack":true,"cangift":true,"areaid":1001,"ishx":1,"ftype":1,"tjcount":85,"qycount":0,"v1":0,"yellow_vip_level":0,"qqname":"马夏坦布译扎"},{"id":11938,"sex":1,"figure":99000002,"username":"迪亚","battlevalue":335063,"mlv":20,"canAttack":true,"cangift":true,"areaid":1001,"ishx":1,"ftype":1,"tjcount":119,"qycount":0,"v1":0,"yellow_vip_level":0,"qqname":"迪亚"},{"id":6845,"sex":1,"figure":99000002,"username":"暗珍高谢科生","battlevalue":334257,"mlv":35,"canAttack":true,"cangift":true,"areaid":1001,"ishx":1,"ftype":1,"tjcount":90,"qycount":0,"v1":0,"yellow_vip_level":0,"qqname":"暗珍高谢科生"},{"id":22220,"sex":1,"figure":99000001,"username":"阿桑莫比斯","battlevalue":320253,"mlv":35,"canAttack":true,"cangift":true,"areaid":1001,"ishx":1,"ftype":1,"tjcount":71,"qycount":0,"v1":0,"yellow_vip_level":0,"qqname":"阿桑莫比斯"},{"id":22437,"sex":2,"figure":99000004,"username":"星海蓝","battlevalue":304330,"mlv":30,"canAttack":true,"cangift":true,"areaid":1001,"ishx":1,"ftype":1,"tjcount":83,"qycount":0,"v1":0,"yellow_vip_level":0,"qqname":"星海蓝"},{"id":23490,"sex":2,"figure":99000005,"username":"Deepdawn","battlevalue":284564,"mlv":32,"canAttack":true,"cangift":true,"areaid":1001,"ishx":1,"ftype":1,"tjcount":74,"qycount":0,"v1":0,"yellow_vip_level":0,"qqname":"Deepdawn"},{"id":11659,"sex":1,"figure":99000001,"username":"厄柯韦里屋","battlevalue":271372,"mlv":30,"canAttack":true,"cangift":true,"areaid":1001,"ishx":1,"ftype":1,"tjcount":92,"qycount":0,"v1":0,"yellow_vip_level":0,"qqname":"厄柯韦里屋"},{"id":10596,"sex":2,"figure":19000008,"username":"密马汤乔","battlevalue":235372,"mlv":29,"canAttack":true,"cangift":true,"areaid":1001,"ishx":1,"ftype":1,"tjcount":73,"qycount":0,"v1":0,"yellow_vip_level":0,"qqname":"密马汤乔"},{"id":22340,"sex":1,"figure":99000002,"username":"隆施黛尼","battlevalue":233087,"mlv":32,"canAttack":true,"cangift":false,"areaid":1001,"ishx":1,"ftype":1,"tjcount":81,"qycount":0,"v1":0,"yellow_vip_level":0,"qqname":"隆施黛尼"},{"id":10625,"sex":1,"figure":99000003,"username":"庞和芬列帕奥","battlevalue":229696,"mlv":22,"canAttack":true,"cangift":true,"areaid":1001,"ishx":0,"ftype":1,"tjcount":89,"qycount":0,"v1":0,"yellow_vip_level":0,"qqname":"庞和芬列帕奥"},{"id":23838,"sex":1,"figure":99000001,"username":"契坦切屋尼","battlevalue":224721,"mlv":22,"canAttack":true,"cangift":true,"areaid":1001,"ishx":0,"ftype":1,"tjcount":55,"qycount":0,"v1":0,"yellow_vip_level":0,"qqname":"契坦切屋尼"},{"id":7613,"sex":1,"figure":99000003,"username":"夫魯夫魯","battlevalue":218930,"mlv":27,"canAttack":true,"cangift":true,"areaid":1001,"ishx":1,"ftype":1,"tjcount":56,"qycount":0,"v1":0,"yellow_vip_level":0,"qqname":"夫魯夫魯"},{"id":11044,"sex":1,"figure":99000001,"username":"山詹戈胥福鲁","battlevalue":199587,"mlv":9,"canAttack":true,"cangift":true,"areaid":1001,"ishx":0,"ftype":1,"tjcount":60,"qycount":0,"v1":0,"yellow_vip_level":0,"qqname":"山詹戈胥福鲁"},{"id":27621,"sex":1,"figure":99000001,"username":"雷利","battlevalue":167504,"mlv":30,"canAttack":true,"cangift":true,"areaid":1001,"ishx":1,"ftype":1,"tjcount":80,"qycount":0,"v1":0,"yellow_vip_level":0,"qqname":"雷利"}],"isenabled":1}}


### t = 4004 送礼
#### request
	gifttype:3
	id:12184
	sid:l1n3tyxnxj2yzxurcmhf1kng
	t:4004
	userid:24795
#### response
	略


### t = 4005 查看社交信息（界面右侧）
#### request
	略
#### response
	{"ret":0,"data":[{"senderid":11529,"sendername":"洛各山芙","sendersex":1,"senderfigure":99000001,"id":1010861,"gifttype":2,"state":0,"type":2,"isnew":1},{"senderid":12184,"sendername":"哈啊","sendersex":1,"senderfigure":99000001,"id":1010033,"gifttype":3,"state":0,"type":2,"isnew":1},{"senderid":4523,"sendername":"银魂开发部","sendersex":0,"senderfigure":19000007,"id":1003779,"gifttype":5,"state":0,"type":2,"isnew":0},{"senderid":4523,"sendername":"银魂开发部","sendersex":0,"senderfigure":19000007,"id":1003552,"gifttype":5,"state":0,"type":2,"isnew":0},{"senderid":22453,"sendername":"腻夏苏伦莫立","sendersex":1,"senderfigure":99000001,"id":1001733,"gifttype":3,"state":0,"type":2,"isnew":0},{"senderid":4523,"sendername":"银魂开发部","sendersex":0,"senderfigure":19000007,"id":998799,"gifttype":5,"state":1,"type":2,"isnew":0},{"senderid":4523,"sendername":"银魂开发部","sendersex":0,"senderfigure":19000007,"id":997428,"gifttype":5,"state":1,"type":2,"isnew":0},{"senderid":13931,"sendername":"射命丸文","sendersex":2,"senderfigure":99000005,"id":996562,"gifttype":3,"state":1,"type":2,"isnew":0},{"senderid":11938,"sendername":"迪亚","sendersex":1,"senderfigure":99000002,"id":995414,"gifttype":0,"state":0,"type":1,"isnew":0},{"senderid":23442,"sendername":"俄那麦科","sendersex":1,"senderfigure":99000001,"id":980351,"gifttype":3,"state":1,"type":2,"isnew":0},{"senderid":5382,"sendername":"玩个蛋蛋","sendersex":1,"senderfigure":99000002,"id":974464,"gifttype":3,"state":1,"type":2,"isnew":0},{"senderid":22009,"sendername":"较黄的圣骑士","sendersex":1,"senderfigure":99000001,"id":971158,"gifttype":3,"state":1,"type":2,"isnew":0},{"senderid":4523,"sendername":"银魂开发部","sendersex":0,"senderfigure":19000007,"id":968094,"gifttype":5,"state":1,"type":2,"isnew":0},{"senderid":22340,"sendername":"隆施黛尼","sendersex":1,"senderfigure":99000002,"id":965610,"gifttype":3,"state":1,"type":2,"isnew":0},{"senderid":7394,"sendername":"卡到荼蘼","sendersex":2,"senderfigure":99000004,"id":965229,"gifttype":3,"state":1,"type":2,"isnew":0},{"senderid":4523,"sendername":"银魂开发部","sendersex":0,"senderfigure":19000007,"id":963092,"gifttype":5,"state":1,"type":2,"isnew":0},{"senderid":11529,"sendername":"洛各山芙","sendersex":1,"senderfigure":99000001,"id":960672,"gifttype":2,"state":1,"type":2,"isnew":0},{"senderid":12184,"sendername":"哈啊","sendersex":1,"senderfigure":99000001,"id":960594,"gifttype":3,"state":1,"type":2,"isnew":0},{"senderid":22453,"sendername":"腻夏苏伦莫立","sendersex":1,"senderfigure":99000001,"id":953803,"gifttype":3,"state":1,"type":2,"isnew":0},{"senderid":7417,"sendername":"死宅医生","sendersex":1,"senderfigure":99000002,"id":953552,"gifttype":3,"state":1,"type":2,"isnew":0},{"senderid":12184,"sendername":"哈啊","sendersex":1,"senderfigure":99000001,"id":947239,"gifttype":3,"state":1,"type":2,"isnew":0},{"senderid":13931,"sendername":"射命丸文","sendersex":2,"senderfigure":99000005,"id":944342,"gifttype":3,"state":1,"type":2,"isnew":0},{"senderid":10671,"sendername":"寿杰奈茨麦叙","sendersex":2,"senderfigure":99000005,"id":930340,"gifttype":0,"state":0,"type":1,"isnew":0},{"senderid":22009,"sendername":"较黄的圣骑士","sendersex":1,"senderfigure":99000001,"id":919877,"gifttype":2,"state":1,"type":2,"isnew":0},{"senderid":6664,"sendername":"普尔胡巴","sendersex":1,"senderfigure":99000002,"id":918616,"gifttype":3,"state":1,"type":2,"isnew":0},{"senderid":5382,"sendername":"玩个蛋蛋","sendersex":1,"senderfigure":99000002,"id":916697,"gifttype":3,"state":1,"type":2,"isnew":0},{"senderid":22340,"sendername":"隆施黛尼","sendersex":1,"senderfigure":99000002,"id":916493,"gifttype":3,"state":1,"type":2,"isnew":0},{"senderid":4523,"sendername":"银魂开发部","sendersex":0,"senderfigure":19000007,"id":910094,"gifttype":5,"state":1,"type":2,"isnew":0},{"senderid":4523,"sendername":"银魂开发部","sendersex":0,"senderfigure":19000007,"id":909141,"gifttype":5,"state":1,"type":2,"isnew":0},{"senderid":7417,"sendername":"死宅医生","sendersex":1,"senderfigure":99000002,"id":903496,"gifttype":3,"state":1,"type":2,"isnew":0},{"senderid":23442,"sendername":"俄那麦科","sendersex":1,"senderfigure":99000001,"id":896957,"gifttype":3,"state":1,"type":2,"isnew":0},{"senderid":22453,"sendername":"腻夏苏伦莫立","sendersex":1,"senderfigure":99000001,"id":894226,"gifttype":3,"state":1,"type":2,"isnew":0},{"senderid":6836,"sendername":"HZH","sendersex":1,"senderfigure":99000003,"id":887758,"gifttype":3,"state":1,"type":2,"isnew":0},{"senderid":15001,"sendername":"名柴匹良","sendersex":1,"senderfigure":99000003,"id":878050,"gifttype":0,"state":0,"type":1,"isnew":0},{"senderid":22340,"sendername":"隆施黛尼","sendersex":1,"senderfigure":99000002,"id":873375,"gifttype":3,"state":1,"type":2,"isnew":0},{"senderid":6664,"sendername":"普尔胡巴","sendersex":1,"senderfigure":99000002,"id":871332,"gifttype":3,"state":1,"type":2,"isnew":0},{"senderid":22009,"sendername":"较黄的圣骑士","sendersex":1,"senderfigure":99000001,"id":866864,"gifttype":3,"state":1,"type":2,"isnew":0},{"senderid":12184,"sendername":"哈啊","sendersex":1,"senderfigure":99000001,"id":865283,"gifttype":3,"state":1,"type":2,"isnew":0},{"senderid":22781,"sendername":"恩雷达杰","sendersex":1,"senderfigure":99000003,"id":863710,"gifttype":0,"state":0,"type":1,"isnew":0}]}
#### 备注
	state=0 代表可以领取

### t = 4006 领取好友礼物
#### request
	id:1010861
	sid:ylzxw2plsv55kcnicbp24cxq
	t:4006
	userid:24795
#### response
	{"ret":0,"data":{"msg":"一乐拉面*1,","rewards":[{"type":4,"id":4,"count":1,"rewardid":1000006}]}}


### t = 4008 查询当前冒险战利品数量
#### request
	略
#### response
	{"ret":0,"data":{"count":3}}


### t = 4009 添加好友
#### request
	userid:24795
	sid:jtlz21qhzyuulg0nxmfomnsv
	t:4009
	name:Deepdawn (用户名)
#### response
	略


### t = 5002 获取冒险相关信息
#### request 
	略
#### response
	{"ret":0,"data":{"bzd":42674,"event1iscomplete":1,"event2iscomplete":1,"event3iscomplete":1,"event1time":1200,"event2time":1200,"event3time":1200,"boxcount":0,"boxtime":300,"isoutfirstbox":1,"jb":20608504,"mlist":[{"id":329458,"mid":99000002,"exp":31649254,"wakeup":4,"equiplv":6,"isplaying":1,"isleader":1,"lv":39,"battlevalue":66500,"tplv":0,"isbj":0,"islock":0},{"id":876849,"mid":19000036,"exp":14766913,"wakeup":4,"equiplv":12,"isplaying":1,"isleader":0,"lv":29,"battlevalue":39020,"tplv":0,"isbj":0,"islock":0},{"id":919444,"mid":19000006,"exp":4165256,"wakeup":3,"equiplv":0,"isplaying":1,"isleader":0,"lv":40,"battlevalue":12031,"tplv":0,"isbj":0,"islock":0},{"id":654998,"mid":18000042,"exp":18367258,"wakeup":4,"equiplv":8,"isplaying":1,"isleader":0,"lv":32,"battlevalue":29696,"tplv":0,"isbj":0,"islock":0},{"id":616352,"mid":18000002,"exp":19366328,"wakeup":3,"equiplv":15,"isplaying":1,"isleader":0,"lv":58,"battlevalue":19107,"tplv":0,"isbj":0,"islock":0},{"id":810612,"mid":18000002,"exp":10402207,"wakeup":3,"equiplv":0,"isplaying":1,"isleader":0,"lv":50,"battlevalue":11372,"tplv":0,"isbj":0,"islock":0},{"id":616242,"mid":18000013,"exp":19372180,"wakeup":3,"equiplv":17,"isplaying":1,"isleader":0,"lv":58,"battlevalue":19784,"tplv":0,"isbj":0,"islock":0},{"id":780547,"mid":18000030,"exp":11496315,"wakeup":3,"equiplv":0,"isplaying":1,"isleader":0,"lv":51,"battlevalue":11396,"tplv":0,"isbj":0,"islock":0},{"id":446334,"mid":19000017,"exp":22257686,"wakeup":3,"equiplv":5,"isplaying":1,"isleader":0,"lv":60,"battlevalue":14894,"tplv":0,"isbj":0,"islock":0},{"id":502627,"mid":19000017,"exp":21892084,"wakeup":3,"equiplv":5,"isplaying":1,"isleader":0,"lv":60,"battlevalue":14894,"tplv":0,"isbj":0,"islock":0},{"id":330029,"mid":18000022,"exp":22691999,"wakeup":3,"equiplv":5,"isplaying":1,"isleader":0,"lv":60,"battlevalue":12602,"tplv":0,"isbj":0,"islock":0},{"id":726758,"mid":18000022,"exp":17669205,"wakeup":3,"equiplv":0,"isplaying":1,"isleader":0,"lv":57,"battlevalue":11570,"tplv":0,"isbj":0,"islock":0},{"id":810510,"mid":17000004,"exp":11523301,"wakeup":3,"equiplv":9,"isplaying":1,"isleader":0,"lv":51,"battlevalue":15324,"tplv":0,"isbj":0,"islock":0},{"id":810728,"mid":18000027,"exp":11455199,"wakeup":4,"equiplv":12,"isplaying":1,"isleader":0,"lv":25,"battlevalue":31925,"tplv":0,"isbj":0,"islock":0},{"id":847430,"mid":18000041,"exp":10594711,"wakeup":3,"equiplv":6,"isplaying":1,"isleader":0,"lv":50,"battlevalue":12379,"tplv":0,"isbj":0,"islock":0},{"id":446232,"mid":18000010,"exp":21245539,"wakeup":3,"equiplv":18,"isplaying":1,"isleader":0,"lv":59,"battlevalue":22073,"tplv":0,"isbj":0,"islock":0},{"id":751105,"mid":11000032,"exp":10903591,"wakeup":3,"equiplv":0,"isplaying":1,"isleader":0,"lv":50,"battlevalue":10207,"tplv":0,"isbj":0,"islock":0},{"id":877064,"mid":18000007,"exp":3624166,"wakeup":3,"equiplv":0,"isplaying":1,"isleader":0,"lv":38,"battlevalue":9643,"tplv":0,"isbj":0,"islock":0},{"id":963044,"mid":18000017,"exp":3624152,"wakeup":3,"equiplv":0,"isplaying":1,"isleader":0,"lv":38,"battlevalue":10429,"tplv":0,"isbj":0,"islock":0},{"id":936447,"mid":18000053,"exp":3684241,"wakeup":3,"equiplv":0,"isplaying":1,"isleader":0,"lv":38,"battlevalue":10397,"tplv":0,"isbj":0,"islock":0}]}}


### t = 5004 刷新冒险体力
#### request
	略
#### response
	略


### t = 5007 获取冒险战利品
#### request
	略
#### response
	{"ret":0,"data":{"rewards":[{"type":1,"id":8000003,"count":1,"rewardid":18000795}]}}


### t = 6001 获取竞技场对手列表
#### request
	略
#### response
	{"ret":0,"data":{"list":[{"id":3053656,"sex":0,"figure":"11000019","username":"布拉格布兰奇","battlevalue":295000,"lv":0,"iswin":0,"userid":0},{"id":3053657,"sex":0,"figure":"18000007","username":"胡夫凯特","battlevalue":312400,"lv":0,"iswin":0,"userid":0},{"id":3053658,"sex":0,"figure":"99000001","username":"放开我的娑娜","battlevalue":312516,"lv":0,"iswin":0,"userid":11064},{"id":3053659,"sex":0,"figure":"99000005","username":"ko7777","battlevalue":363655,"lv":0,"iswin":0,"userid":13295},{"id":3053660,"sex":0,"figure":"99000001","username":"提督","battlevalue":363706,"lv":0,"iswin":0,"userid":17202},{"id":3053661,"sex":0,"figure":"99000002","username":"荷施莎可瑞莫","battlevalue":363809,"lv":0,"iswin":0,"userid":10404},{"id":3053662,"sex":0,"figure":"99000002","username":"流浪的牧羊人","battlevalue":378255,"lv":0,"iswin":0,"userid":9691},{"id":3053663,"sex":0,"figure":"18000013","username":"马骁尧","battlevalue":378419,"lv":0,"iswin":0,"userid":17904},{"id":3053664,"sex":0,"figure":"15000010","username":"辛迪卡斯基","battlevalue":392400,"lv":0,"iswin":0,"userid":0}],"time":6966,"arenacount":0,"arenabox1isget":0,"arenabox2isget":0,"arenabox3isget":0}}


### t = 6006 领取竞技场奖励
#### request
	userid:24795
	sid:xxxzpi3mxxz21wmzquznlzlc
	t:6006
	type:1(或4、9，代表领取几胜奖励)
#### response
	{"ret":0,"data":{"msg":{"msg":"勇气点数*3,冠军点数*2,","list":[{"type":27,"count":3,"rewardid":1400002},{"type":28,"count":2,"rewardid":1400003}]}}}


### t = 6007 进行竞技场战斗
#### request
	id:3053656 	(对手id)
	sid:xxxzpi3mxxz21wmzquznlzlc
	t:6007
	userid:24795
#### response
	{"ret":0,"data":{"report":{"leftMG":{"name":"七月汪汪","merList":[99000002,19000036,19000006,18000013,18000002,18000042,18000022,18000022,19000017,19000017,18000030,18000023,18000026,17000004,18000002,18000009,18000027,18000041,18000010,11000032],"ATK":342813,"maxATK":342813,"skills":[1100009,1100032,1200041,1200026,1900015,1200053,1100006,1200026,1200028,1100018,1100008,1100013],"skillsLv":1,"firstStrike":31,"defense":44,"dodge":9,"wz":0},"rightMG":{"name":"布拉格布兰奇","merList":[11000019,16000011,17000003,13000015,18000013,13000013,19000001,17000009,12000008,16000010,11000030,11000019,11000023,17000003,11000001,10000004,16000011,16000010,15000011,19000006],"ATK":295000,"maxATK":295000,"skills":[1700001,1200021,2200003,1200041,2200001,2000003,1200020,1700001,2200003,1200002,1200021,1200020,1200018,1200030],"skillsLv":1,"firstStrike":4,"defense":26,"dodge":8,"wz":0},"skillList":[1200030,2000003,1200028],"dodgeList":[false,false,false],"damageValue":[[0,-261189],[-43320,46349],[0,-121184]],"xly":[],"skip":1},"iswin":true}}



### t = 7001 获取食物列表
#### request
	略
#### response
	{"ret":0,"data":[{"id":94528,"foodid":1,"count":2},{"id":94423,"foodid":2,"count":26},{"id":94509,"foodid":3,"count":8},{"id":94529,"foodid":4,"count":10},{"id":111420,"foodid":5,"count":11},{"id":111475,"foodid":6,"count":0},{"id":114049,"foodid":7,"count":2},{"id":111372,"foodid":8,"count":12},{"id":94422,"foodid":9,"count":10},{"id":117616,"foodid":10,"count":1},{"id":94455,"foodid":11,"count":17},{"id":115652,"foodid":12,"count":0}]}


### t = 7002 吃食物
#### request
	userid:24795
	sid:u3iudktskhvjkvbe5v4rcua4
	t:7002
	foodid:7
#### response
	{"ret":0,"data":{"rewards":[{"type":24,"count":3600,"rewardid":1200008}]}}


### t = 8003 一键消费经验道具
#### request
	略
#### response
	略
	

### t = 8004 一键消费金币道具
#### request
	略
#### response
	略


### t = 190101 获取签到列表
#### request
	略
#### response
	{"ret":0,"data":{"count":0,"isget":false,"month":5,"praylv":1,"alchemistlv":1,"ispray":false,"isalchemist":false,"date":"155","list":[44000233,44000234,44000235,44000236,44000237,44000238,44000239,44000240,44000241,44000242,44000243,44000244,44000245,44000246,44000247,44000248,44000249,44000250,44000251,44000252]}}


### t = 190102 获取今日签到奖励
#### request
	略
#### response
	{"ret":0,"data":{"rewards":[{"type":4,"id":11,"count":6,"rewardid":44000233}]}}


### t = 190201 领取月卡
#### request
	id:7 (或8。7为初级月卡，8为高级月卡)
	sid:ylzxw2plsv55kcnicbp24cxq
	t:190201
	userid:24795
#### response
	{"ret":0,"data":{"rewards":[{"type":99,"id":1000003,"count":3,"rewardid":45000013}]}}


### t = 190202 查看月卡领取情况
#### request
	略
#### response
	{"ret":0,"data":{"days":[14,17,0,28],"isget":[true,true,false]}}