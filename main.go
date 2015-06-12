package main

import (
	"github.com/astaxie/beego"
	_ "github.com/julycw/pixelKnightHacker/routers"
)

func main() {
	beego.SessionOn = true
	beego.Run()
}
