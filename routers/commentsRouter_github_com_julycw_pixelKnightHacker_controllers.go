package routers

import (
	"github.com/astaxie/beego"
)

func init() {
	
	beego.GlobalControllerRouter["github.com/julycw/pixelKnightHacker/controllers:CommandController"] = append(beego.GlobalControllerRouter["github.com/julycw/pixelKnightHacker/controllers:CommandController"],
		beego.ControllerComments{
			"MineQueue",
			`/command/mineQueue`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/julycw/pixelKnightHacker/controllers:CommandController"] = append(beego.GlobalControllerRouter["github.com/julycw/pixelKnightHacker/controllers:CommandController"],
		beego.ControllerComments{
			"Jjc",
			`/command/jjc`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/julycw/pixelKnightHacker/controllers:CommandController"] = append(beego.GlobalControllerRouter["github.com/julycw/pixelKnightHacker/controllers:CommandController"],
		beego.ControllerComments{
			"Explore",
			`/command/explore`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/julycw/pixelKnightHacker/controllers:CommandController"] = append(beego.GlobalControllerRouter["github.com/julycw/pixelKnightHacker/controllers:CommandController"],
		beego.ControllerComments{
			"Friends",
			`/command/friends`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/julycw/pixelKnightHacker/controllers:CommandController"] = append(beego.GlobalControllerRouter["github.com/julycw/pixelKnightHacker/controllers:CommandController"],
		beego.ControllerComments{
			"Daywork",
			`/command/daywork`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/julycw/pixelKnightHacker/controllers:CommandController"] = append(beego.GlobalControllerRouter["github.com/julycw/pixelKnightHacker/controllers:CommandController"],
		beego.ControllerComments{
			"HeartBeat",
			`/command/heartBeat`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/julycw/pixelKnightHacker/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/julycw/pixelKnightHacker/controllers:UserController"],
		beego.ControllerComments{
			"Get",
			`/login`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/julycw/pixelKnightHacker/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/julycw/pixelKnightHacker/controllers:UserController"],
		beego.ControllerComments{
			"Logout",
			`/logout`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/julycw/pixelKnightHacker/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/julycw/pixelKnightHacker/controllers:UserController"],
		beego.ControllerComments{
			"Post",
			`/login`,
			[]string{"post"},
			nil})

}
