package main

import (
	"CloudRestaurant/controller"
	"CloudRestaurant/tool"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := tool.ParseConfig("./config/app.json")
	if err != nil {
		panic(err.Error())
	}
	app := gin.Default()
	regiestRouter(app)
	app.Run(cfg.AppHost + ":" + cfg.AppPort)
}

//路由设置
func regiestRouter(router *gin.Engine) {
	new(controller.HelloController).Router(router)
	new(controller.MemberController).Router(router)
}
