package services

import (
	"encoding/json"
	"os"

	"github.com/psycomentis/psycofolio++/src"
	"github.com/rs/zerolog/log"
)

type Config struct {
	Version string `json:"version" default:"1"`

	ServerPort string `json:"server_port" default:"8080"`

	AdminUsername  string `json:"admin_username" default:"admin"`
	AdminPassword  string `json:"admin_password" default:"admin"`
	PasswordMethod string `json:"password_method" default:"plain"`

	DatabaseConnectionString string `json:"db_connection" default:"./psyco.db"`
	DatabaseEngine           string `json:"database_engine" default:"sqlite"`

	DefaultLocale  string `json:"default_locale" default:"en_US"`
	SelectedLocale string `json:"selected_locale" default:"en_US"`
}

// Loads file from default directory
func LoadConfig(path string) (Config, error) {
	file, openErr := os.Open(ConfigFile)
	if openErr != nil {
		cnf := CreateDefaultConfig()
		cnf.ExportToJson(path)
		log.Log().Msg("Config file not found. Generating default config to: " + path)
		return cnf, nil
	}
	var mapData map[string]interface{}
	decodeErr := json.NewDecoder(file).Decode(&mapData)
	if decodeErr != nil {
		log.Error().
			Err(decodeErr).
			Msg("Failed to decode configuration file.")
	}
	configVer := mapData["version"]
	if configVer != "" {
		switch configVer {
		case "1":
			return parseConfigV1(&mapData), nil
		}
	}
	return Config{}, nil
}

func parseConfigV1(m *map[string]interface{}) Config {
	cnf := Config{
		Version:                  "1",
		ServerPort:               src.GetOrDefault(m, "server_port", "8080"),
		AdminUsername:            src.GetOrDefault(m, "admin_username", "admin"),
		AdminPassword:            src.GetOrDefault(m, "admin_password", "admin"),
		PasswordMethod:           src.GetOrDefault(m, "password_method", "plain"),
		DatabaseConnectionString: src.GetOrDefault(m, "db_connection", "./psyco.db"),
		DatabaseEngine:           src.GetOrDefault(m, "database_engine", "sqlite"),
		DefaultLocale:            src.GetOrDefault(m, "default_locale", "en_US"),
		SelectedLocale:           src.GetOrDefault(m, "selected_locale", "en_US"),
	}
	return cnf
}

func CreateDefaultConfig() Config {
	return Config{
		Version:                  "1",
		ServerPort:               "8080",
		AdminUsername:            "admin",
		AdminPassword:            "admin",
		PasswordMethod:           "plain",
		DefaultLocale:            "en_US",
		SelectedLocale:           "en_US",
		DatabaseEngine:           "sqlite",
		DatabaseConnectionString: ApplicationDir + "/psyco.db",
	}
}

func (cnf *Config) ExportToJson(path string) error {
	d, err := json.Marshal(cnf)
	if err != nil {
		return err
	}
	return os.WriteFile(path, d, os.ModePerm)
}

/* func CreateDefaultConfig() (Config, error) {
	cnf := Config{}
	ref := reflect.TypeOf(cnf)
	for i := 0; i < ref.NumField(); i++ {
		f := ref.Field(i)
		defaultTag := f.Tag.Get("default")
		if defaultTag != "" {
			value := reflect.ValueOf(&cnf).Elem()
			fVal := value.FieldByName(f.Name)
			fVal.SetString(defaultTag)
		}
	}
	fmt.Print(cnf)
	return cnf, nil
} */
