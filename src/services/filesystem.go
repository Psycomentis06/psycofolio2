package services

import (
	"github.com/psycomentis/psycofolio++/src"
	"github.com/rs/zerolog/log"
)

// Application Defaults
var ApplicationDir = "$HOME/.locale/share/psycofolio2"
var LocalesDir = ApplicationDir + "/locales"
var ConfigDir = "$HOME/.config/psycofolio2"
var ConfigFile = ConfigDir + "/config.json"

func InitApplicationFolders() {
	appDirPath, appDirErr := src.GetFullPath(ApplicationDir)
	if appDirErr != nil {
		log.Err(appDirErr)
		return
	}
	ApplicationDir = appDirPath

	cnfDirPath, cnfDirErr := src.GetFullPath(ConfigDir)
	if cnfDirErr != nil {
		log.Err(cnfDirErr)
		return
	}
	ConfigDir = cnfDirPath
}
