package config

import (
	"7D-admin/utils/fileUtils"
	"log"
	"strings"
)

type DirConfig struct {
	InstallDir     string `yaml:"installDir"`
	SteamcmdDir    string `yaml:"steamcmdDir"`
	ModDir         string `yaml:"modDir"`
	SaveDir        string `yaml:"saveDir"`
	GameLogDir     string `yaml:"gameLogDir"`
	PanelLogDir    string `yaml:"panelLogDir"`
	SeverconfigDir string `yaml:"severconfigDir"`
	ServeradminDir string `yaml:"serveradminDir"`
}

const dir_config_path = "./dir_config"

func NewDirConfig() *DirConfig {
	return &DirConfig{}
}
func GetDirConfig() DirConfig {
	dirConfig := NewDirConfig()
	if !fileUtils.Exists(dir_config_path) {
		if err := fileUtils.CreateFile(dir_config_path); err != nil {
			log.Panicln("create dir_config error", err)
		}
	}
	data, err := fileUtils.ReadLnFile(dir_config_path)
	if err != nil {
		log.Panicln("read dir_config error", err)
	}
	for _, value := range data {
		if value == "" {
			continue
		}
		if strings.Contains(value, "installDir=") {
			split := strings.Split(value, "=")
			if len(split) > 1 {
				s := strings.TrimSpace(split[1])
				dirConfig.InstallDir = strings.Replace(s, "\\n", "", -1)
			}
		}
		if strings.Contains(value, "steamcmdDir=") {
			split := strings.Split(value, "=")
			if len(split) > 1 {
				s := strings.TrimSpace(split[1])
				dirConfig.SteamcmdDir = strings.Replace(s, "\\n", "", -1)
			}
		}
		if strings.Contains(value, "modDir=") {
			split := strings.Split(value, "=")
			if len(split) > 1 {
				s := strings.TrimSpace(split[1])
				dirConfig.ModDir = strings.Replace(s, "\\n", "", -1)
			}
		}
		if strings.Contains(value, "saveDir=") {
			split := strings.Split(value, "=")
			if len(split) > 1 {
				s := strings.TrimSpace(split[1])
				dirConfig.SaveDir = strings.Replace(s, "\\n", "", -1)
			}
		}
		if strings.Contains(value, "gameLogDir=") {
			split := strings.Split(value, "=")
			if len(split) > 1 {
				s := strings.TrimSpace(split[1])
				dirConfig.GameLogDir = strings.Replace(s, "\\n", "", -1)
			}
		}
		if strings.Contains(value, "panelLogDir=") {
			split := strings.Split(value, "=")
			if len(split) > 1 {
				s := strings.TrimSpace(split[1])
				dirConfig.PanelLogDir = strings.Replace(s, "\\n", "", -1)
			}
		}
		if strings.Contains(value, "severconfigDir=") {
			split := strings.Split(value, "=")
			if len(split) > 1 {
				s := strings.TrimSpace(split[1])
				dirConfig.SeverconfigDir = strings.Replace(s, "\\n", "", -1)
			}
		}
		if strings.Contains(value, "serveradminDir=") {
			split := strings.Split(value, "=")
			if len(split) > 1 {
				s := strings.TrimSpace(split[1])
				dirConfig.ServeradminDir = strings.Replace(s, "\\n", "", -1)
			}
		}
		if dirConfig.InstallDir == "" {
			dirConfig.InstallDir = "/root/7DaysToDieServer"
		}
		if dirConfig.SteamcmdDir == "" {
			dirConfig.SteamcmdDir = "/root"
		}
		if dirConfig.ModDir == "" {
			dirConfig.ModDir = "/root/7DaysToDieServer/Mods"
		}
		if dirConfig.SaveDir == "" {
			dirConfig.SaveDir = "/root/.local/share/7DaysToDie/Saves"
		}
		if dirConfig.GameLogDir == "" {
			dirConfig.GameLogDir = "/root/7DaysToDieServer/7DaysToDieServer_Data"
		}
		if dirConfig.PanelLogDir == "" {
			dirConfig.PanelLogDir = "./PanelLogs"
		}
		if dirConfig.SeverconfigDir == "" {
			dirConfig.SeverconfigDir = "/root/7DaysToDieServer/serverconfig.xml"
		}
		if dirConfig.ServeradminDir == "" {
			dirConfig.ServeradminDir = "/root/.local/share/7DaysToDie/Saves/serveradmin.xml"
		}
	}
	return *dirConfig
}
func SaveDirConfig(dirConfig DirConfig) {
	log.Println("dirConfig:", dirConfig)
	err := fileUtils.WriterLnFile(dir_config_path, []string{
		"installDir=" + dirConfig.InstallDir,
		"steamcmdDir=" + dirConfig.SteamcmdDir,
		"modDir=" + dirConfig.ModDir,
		"saveDir=" + dirConfig.SaveDir,
		"gameLogDir=" + dirConfig.GameLogDir,
		"panelLogDir=" + dirConfig.PanelLogDir,
		"severconfigDir=" + dirConfig.SeverconfigDir,
		"serveradminDir=" + dirConfig.ServeradminDir,
	})
	if err != nil {
		log.Panicln("save dir_config error", err)
	}
}
type Config struct {
	Port           string `yaml:"port"`
	Path           string `yaml:"path"`
	Db             string `yaml:"database"`
	Steamcmd       string `yaml:"steamcmd"`
	SteamAPIKey    string `yaml:"steamAPIKey"`
	OPENAI_API_KEY string `yaml:"OPENAI_API_KEY"`
	Prompt         string `yaml:"prompt"`
	Flag           string `yaml:"flag"`

	Token string `yaml:"token"`

	AutoUpdateModinfo struct {
		Enable              bool `yaml:"enable"`
		CheckInterval       int  `yaml:"checkInterval"`
		UpdateCheckInterval int  `yaml:"updateCheckInterval"`
	} `yaml:"autoUpdateModinfo"`
}
