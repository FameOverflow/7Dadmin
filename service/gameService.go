package service

import (
	"7D-admin/config"
	"7D-admin/constant/consts"
	"7D-admin/model"
	"7D-admin/utils/shellUtils"
	"log"
	"strings"
	"sync"
	"time"
)

var launchLock = sync.Mutex{}

type GameService struct {
	lock sync.Mutex
	c    HomeService

	logRecord LogRecordService
}

func (g *GameService) LaunchGame(worldName string) {
	launchLock.Lock()
	defer func() {
		launchLock.Unlock()
		if r := recover(); r != nil {
		}
	}()
	dir_config := config.GetDirConfig()
	InstallDir := dir_config.InstallDir
	serverconfigDir := dir_config.SeverconfigDir
	var startCmd = ""

	startCmd = "cd " + InstallDir + "screen -d -m -S \"" + worldName + "\"  ./startserver.sh"
	if serverconfigDir != "" {
		startCmd += " -configfile=" + serverconfigDir
	}
	startCmd += "  ;"

	log.Println("正在启动游戏", "command: ", startCmd)
	_, err := shellUtils.Shell(startCmd)
	if err != nil {
		log.Panicln("启动 "+worldName+" error,", err)
	}
}

func (g *GameService) GetGameStatus(worldName string) bool {
	cmd := " ps -ef | grep -v grep | grep -v tail |grep '" + worldName + "'|grep " + " |sed -n '1P'|awk '{print $2}' "
	result, err := shellUtils.Shell(cmd)
	if err != nil {
		return false
	}
	res := strings.Split(result, "\n")[0]
	log.Println("查询世界状态", cmd, result, res, res != "")
	return res != ""
}

func (g *GameService) ShutdownGame(worldName string) {
	if !g.GetGameStatus(worldName) {
		return
	}
	shell := "screen -S \"" + worldName + "\" -p 0 -X stuff $'\003'"
	log.Println("正在shutdown世界", "worldName: ", worldName, "command: ", shell)
	_, err := shellUtils.Shell(shell)
	if err != nil {
		log.Println("shut down " + worldName + " error: " + err.Error())
		log.Println("shutdown 失败，将强制杀掉世界")
	}
}

func ClearScreen() bool {
	result, err := shellUtils.Shell(consts.ClearScreenCmd)
	if err != nil {
		return false
	}
	res := strings.Split(result, "\n")[0]
	return res != ""
}
func (g *GameService)StartGame(worldName string) {
	g.ShutdownGame(worldName)
	g.LaunchGame(worldName)
	ClearScreen()
}
func (g *GameService)KillGame(worldName string) {
	if !g.GetGameStatus(worldName) {
		return
	}
	cmd := " ps -ef | grep -v grep | grep -v tail |grep '" + worldName + "' |sed -n '1P'|awk '{print $2}' |xargs kill -9"
	log.Println("正在kill世界", "worldName: ", worldName, "command: ", cmd)
	_, err := shellUtils.Shell(cmd)
	if err != nil {
		// TODO 强制杀掉
		log.Println("kill "+worldName+" error: ", err)
	}
}
func (g *GameService)StopGame(worldName string) {
	launchLock.Lock()
	defer func() {
		launchLock.Unlock()
		if r := recover(); r != nil {
			// Do nothing
		}
	}()
		g.logRecord.RecordLog(worldName,model.STOP)
		g.ShutdownGame(worldName)
		time.Sleep(3 * time.Second)
		if g.GetGameStatus(worldName) {
			var i uint8 =1
			for{
				if g.GetGameStatus(worldName){
					break
				}
				g.ShutdownGame(worldName)
				time.Sleep(1 * time.Second)
				i++
				if i>3{
					break
				}
			}
		}
		g.KillGame(worldName)
}