package services

import (
	"os"

	"github.com/rs/zerolog/log"
)

type Config struct {
	Version string `json:"version" default:"1"`

	DataDir   string `json:"data_dir" default:"$HOME/.locale/share/psycofolio2"`
	ConfigDir string `json:"config_dir" default:"$HOME/.config/psycofolio2"`

	ServerPort uint16 `json:"server_port" default:"80"`

	AdminUsername  string `json:"admin_username" default:"admin"`
	AdminPassword  string `json:"admin_password" default:"admin"`
	PasswordMethod string `json:"password_method" default:"plain"`

	DatabaseConnectionString string `json:"db_connection"`
	DatabaseEngine           string `json:"database_engine" default:"sqlite"`

	DefaultLanguage string `json:"default_lang" default:"en_US"`
}

// Loads file from default directory
func LoadConfig(path string) (Config, error) {
	_, openErr := os.Open(ConfigFile)
	if openErr != nil {
		log.Print("Config File not found. Using Defaults")
	}
	return Config{}, nil
}

func parseConfigV1() {}

func CreateDefaultConfig() Config {
	return Config{
		Version:                  "1",
		ServerPort:               80,
		AdminUsername:            "admin",
		AdminPassword:            "admin",
		PasswordMethod:           "plain",
		DefaultLanguage:          "en_US",
		DatabaseEngine:           "sqlite",
		DatabaseConnectionString: ApplicationDir + "psyco.db",
	}
}

func (cnf *Config) ExportToJson(path string) {
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
