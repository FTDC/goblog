package main

import (
	_ "myapp/routers"
	"github.com/astaxie/beego"
	"myapp/models"
)

func main() {

	models.RegisterDB()
	beego.Run()
}

