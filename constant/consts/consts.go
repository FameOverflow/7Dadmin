package consts

import (
	"7D-admin/utils/systemUtils"
	"fmt"
	"path/filepath"
)

const (
	StartGame   = 0
	StopGame   = 0
	// ClearScreenCmd 检查目前所有的screen作业，并删除已经无法使用的screen作业
	ClearScreenCmd = "screen -wipe "
	UpdateGameVersion = "updateGameVersion"
	UpdateGameMod     = "updateGameMod"
	UPDATE_GAME = "UPDATE_GAME"
	TURN = 1
	OFF  = 0
)

var HomePath string
var GamePath string
func init() {
	home, err := systemUtils.Home()
	if err != nil {
		panic("Home path error: " + err.Error())
	}
	HomePath = home
	fmt.Println("home path: " + HomePath)

	GamePath = filepath.Join(home, "/root/7DaysToDieServer")
}
