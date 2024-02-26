package main

import (
	"7D-admin/bootstrap"
	"7D-admin/config/global"
	"7D-admin/router"
	"fmt"
)

func main() {
	bootstrap.Init()
	app := router.NewRoute()
	err := app.Run(":" + global.Config.Port)
	if err != nil {
		fmt.Println("启动失败！！！", err)
	}
}