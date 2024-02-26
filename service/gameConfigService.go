package service

import (
	"dst-admin-go/constant"
	"dst-admin-go/utils/dstConfigUtils"
	"dst-admin-go/utils/dstUtils"
	"dst-admin-go/utils/fileUtils"
	"dst-admin-go/utils/levelConfigUtils"
	"dst-admin-go/vo"
	"log"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

var server_init_template = "./static/template/serverconfig.xml"
var serveradmin_init_template = "./static/template/serveradmin.xml"
type GameConfigService struct {
	w HomeService
}

func (c *GameConfigService) GetConfig(worldName string) vo.GameConfigVO {

	gameConfig := vo.NewGameConfigVO()
	c.GetServerIni(worldName, gameConfig)
	gameConfig.ModData = c.getModoverrides(worldName)

	return *gameConfig
}



func (c *GameConfigService) GetServerIni(worldName string, gameconfig *vo.GameConfigVO) {

	clusterIniPath := dstUtils.GetClusterIniPath(worldName)
	clusterIni, err := fileUtils.ReadLnFile(clusterIniPath)
	if err != nil {
		panic("read cluster.ini file error: " + err.Error())
	}
	for _, value := range clusterIni {
		if value == "" {
			continue
		}
		if strings.Contains(value, "game_mod") {
			split := strings.Split(value, "=")
			if len(split) > 1 {
				s := strings.TrimSpace(split[1])
				gameconfig.GameMode = s
			}
		}
		if strings.Contains(value, "max_players") {
			split := strings.Split(value, "=")
			if len(split) > 1 {
				s := strings.TrimSpace(split[1])
				n, err := strconv.ParseUint(s, 10, 8)
				if err == nil {
					gameconfig.MaxPlayers = uint8(n)
				}
			}
		}
		if strings.Contains(value, "pvp") {
			split := strings.Split(value, "=")
			if len(split) > 1 {
				s := strings.TrimSpace(split[1])
				b, err := strconv.ParseBool(s)
				if err == nil {
					gameconfig.Pvp = b
				}
			}
		}
		if strings.Contains(value, "pause_when_empty") {
			split := strings.Split(value, "=")
			if len(split) > 1 {
				s := strings.TrimSpace(split[1])
				b, err := strconv.ParseBool(s)
				if err == nil {
					gameconfig.PauseWhenNobody = b
				}
			}
		}
		if strings.Contains(value, "cluster_intention") {
			split := strings.Split(value, "=")
			if len(split) > 1 {
				s := strings.TrimSpace(split[1])
				gameconfig.ClusterIntention = s
			}
		}
		if strings.Contains(value, "cluster_password") {
			split := strings.Split(value, "=")
			if len(split) > 1 {
				s := strings.TrimSpace(split[1])
				gameconfig.ClusterPassword = s
			}
		}
		if strings.Contains(value, "cluster_description") {
			split := strings.Split(value, "=")
			if len(split) > 1 {
				s := strings.TrimSpace(split[1])
				gameconfig.ClusterDescription = s
			}
		}
		if strings.Contains(value, "cluster_name") {
			split := strings.Split(value, "=")
			if len(split) > 1 {
				s := strings.TrimSpace(split[1])
				gameconfig.ClusterName = s
			}
		}

	}
}



func (c *GameConfigService) SaveConfig(worldName string, gameConfigVo vo.GameConfigVO) {
	//创建mod设置
	c.createModoverrides(worldName, gameConfigVo.ModData)

}

func (c *GameConfigService) createMyDediServerDir() {
	dstConfig := dstConfigUtils.GetDstConfig()
	basePath := constant.GET_DST_USER_GAME_CONFG_PATH()
	myDediServerPath := path.Join(basePath, dstConfig.Cluster)
	log.Println("生成 myDediServer 目录：" + myDediServerPath)
	fileUtils.CreateDir(myDediServerPath)
}

func (c *GameConfigService) createClusterIni(worldName string, gameConfigVo vo.GameConfigVO) {
	clusterIniPath := dstUtils.GetClusterIniPath(worldName)
	log.Println("生成游戏配置文件 cluster.ini文件: ", clusterIniPath)
	oldCluster := c.w.GetClusterIni(worldName)

	oldCluster.ClusterName = gameConfigVo.ClusterName
	oldCluster.ClusterDescription = gameConfigVo.ClusterDescription
	oldCluster.GameMode = gameConfigVo.GameMode
	oldCluster.MaxPlayers = uint(gameConfigVo.MaxPlayers)
	oldCluster.Pvp = gameConfigVo.Pvp
	oldCluster.VoteEnabled = gameConfigVo.VoteEnabled
	oldCluster.PauseWhenNobody = gameConfigVo.PauseWhenNobody
	oldCluster.ClusterPassword = gameConfigVo.ClusterPassword

	clusterIni := dstUtils.ParseTemplate(cluster_init_template, oldCluster)
	fileUtils.WriterTXT(clusterIniPath, clusterIni)
}

func (c *GameConfigService) createClusterToken(worldName, token string) {
	fileUtils.WriterTXT(dstUtils.GetClusterTokenPath(worldName), token)
}

func (c *GameConfigService) createMasteLeveldataoverride(worldName string, mapConfig string) {
	leveldataoverridePath := dstUtils.GetMasterLeveldataoverridePath(worldName)
	log.Println("生成master_leveldataoverride.txt 文件 ", leveldataoverridePath)
	if mapConfig != "" {
		fileUtils.WriterTXT(leveldataoverridePath, mapConfig)
	} else {
		//置空
		fileUtils.WriterTXT(leveldataoverridePath, "")
	}
}
func (c *GameConfigService) createCavesLeveldataoverride(worldName string, mapConfig string) {
	leveldataoverridePath := dstUtils.GetCavesLeveldataoverridePath(worldName)
	log.Println("生成caves_leveldataoverride.lua 文件 ", leveldataoverridePath)
	if mapConfig != "" {
		fileUtils.WriterTXT(leveldataoverridePath, mapConfig)
	} else {
		//置空
		fileUtils.WriterTXT(leveldataoverridePath, "")
	}
}
func (c *GameConfigService) createModoverrides(worldName string, modConfig string) {

	if modConfig != "" {

		config, _ := levelConfigUtils.GetLevelConfig(worldName)
		for i := range config.LevelList {
			fileUtils.WriterTXT(filepath.Join(dstUtils.GetClusterBasePath(worldName), config.LevelList[i].File, "modoverrides.lua"), modConfig)
		}
		var serverModSetup = ""
		//TODO 添加m
		workshopIds := dstUtils.WorkshopIds(modConfig)
		for _, workshopId := range workshopIds {
			serverModSetup += "ServerModSetup(\"" + workshopId + "\")\n"
		}
		fileUtils.WriterTXT(dstUtils.GetModSetup(worldName), serverModSetup)
	}

}

func (c *GameConfigService) UpdateDedicatedServerModsSetup(worldName, modConfig string) {
	if modConfig != "" {
		var serverModSetup = ""
		workshopIds := dstUtils.WorkshopIds(modConfig)
		for _, workshopId := range workshopIds {
			serverModSetup += "ServerModSetup(\"" + workshopId + "\")\n"
		}
		fileUtils.WriterTXT(dstUtils.GetModSetup(worldName), serverModSetup)
	}

}
