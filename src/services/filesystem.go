package services

import (
	"os"

	"github.com/psycomentis/psycofolio++/src"
	"github.com/rs/zerolog/log"
)

// Application Defaults
var ApplicationDir = "$HOME/.local/share/psycofolio2"
var LocaleDir = "./locales"
var ConfigDir = "$HOME/.config/psycofolio2"
var ConfigFile = ConfigDir + "/config.json"

func InitApplicationFolders() {
	appDirPath, appDirErr := src.GetFullPath(ApplicationDir)
	if appDirErr != nil {
		log.Err(appDirErr)
		return
	}
	ApplicationDir = appDirPath
	if mkErr := createDirIfNotExist(ApplicationDir); mkErr != nil {
		log.Error().Err(mkErr).Msg("Failed to create application directory")
		return
	}

	cnfDirPath, cnfDirErr := src.GetFullPath(ConfigDir)
	if cnfDirErr != nil {
		log.Err(cnfDirErr)
		return
	}
	ConfigDir = cnfDirPath
	if mkErr := createDirIfNotExist(ConfigDir); mkErr != nil {
		log.Error().Err(mkErr).Msg("Failed to create application directory")
		return
	}
}

func createDirIfNotExist(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return os.Mkdir(path, os.ModePerm)
	}
	return nil
}
