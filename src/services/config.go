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

	Admin AdminConfig `json:"admin"`

	Database DatabaseConfig `json:"database"`

	Locale LocaleConfig `json:"locale"`
}

type AdminConfig struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Theme     string `json:"theme"`
	ThemeMode string `json:"theme_mode"`
	Locale    string `json:"locale"`
}

type DatabaseConfig struct {
	Engine           string `json:"engine"`            // sqlite or postgres
	ConnectionString string `json:"connection_string"` // required for postgres
}

type LocaleConfig struct {
	Default  string `json:"default"` // fallback locale
	Selected string `json:"selected"`
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
	defer file.Close()
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
		Version:    "1",
		ServerPort: src.GetOrDefault(m, "server_port", "8080"),
	}

	// Build AdminConfig
	adminConfig := AdminConfig{
		Username: src.GetOrDefault(m, "admin_username", "admin"),
		Password: src.GetOrDefault(m, "admin_password", "admin"),
	}
	cnf.Admin = adminConfig

	// Build DatabaseConfig
	dbConfig := DatabaseConfig{
		ConnectionString: src.GetOrDefault(m, "db_connection", "./psyco.db"),
		Engine:           src.GetOrDefault(m, "database_engine", "sqlite"),
	}
	cnf.Database = dbConfig

	// Build LocaleConfig
	localeConfig := LocaleConfig{
		Default:  src.GetOrDefault(m, "default_locale", "en_US"),
		Selected: src.GetOrDefault(m, "selected_locale", "en_US"),
	}
	cnf.Locale = localeConfig

	return cnf
}

func CreateDefaultConfig() Config {
	return Config{
		Version:    "1",
		ServerPort: "8080",
		Admin: AdminConfig{
			Username: "admin",
			Password: "admin",
			Locale:   "en_US",
		},
		Locale: LocaleConfig{
			Default:  "en_US",
			Selected: "en_US",
		},
		Database: DatabaseConfig{
			Engine:           "sqlite",
			ConnectionString: ApplicationDir + "/psyco.db",
		},
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
