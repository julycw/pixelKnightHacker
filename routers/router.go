package routers

import (
	"github.com/astaxie/beego"
	"github.com/julycw/pixelKnightHacker/controllers"
)

func init() {
	beego.Router("/", &controllers.IndexController{})

	beego.Include(&controllers.UserController{})
	beego.Include(&controllers.CommandController{})
}
