package services

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/psycomentis/psycofolio++/src"
)

type Config struct {
	Version string `json:"version"`

	ServerPort string `json:"server_port"`

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
	defer file.Close()
	if openErr != nil {
		cnf := CreateDefaultConfig()
		cnf.ExportToJson(path)
		return cnf, nil
	}
	defer file.Close()
	var mapData map[string]interface{}
	decodeErr := json.NewDecoder(file).Decode(&mapData)
	if decodeErr != nil {
		return Config{}, decodeErr
	}
	configVer := mapData["version"]
	if configVer != "" {
		switch configVer {
		case "1":
			return parseConfigV1(&mapData), nil
		default:
			return Config{}, errors.New("Unkown Configuration Version")
		}
	}
	return Config{}, errors.New("Configuration Version is missing")
}

func parseConfigV1(m *map[string]interface{}) Config {
	cnf := Config{
		Version:    "1",
		ServerPort: src.GetOrDefault(m, "server_port", "8080"),
	}

	// Build AdminConfig
	adminMap, ok := (*m)["admin"].(map[string]interface{})
	if !ok {
		adminMap = make(map[string]interface{})
	}
	adminConfig := AdminConfig{
		Username: src.GetOrDefault(&adminMap, "username", ""),
		Password: src.GetOrDefault(&adminMap, "password", ""),
	}
	cnf.Admin = adminConfig

	// Build DatabaseConfig
	dbMap, ok := (*m)["database"].(map[string]interface{})
	if !ok {
		dbMap = make(map[string]interface{})
	}
	dbConfig := DatabaseConfig{
		ConnectionString: src.GetOrDefault(&dbMap, "connection_string", ""),
		Engine:           src.GetOrDefault(&dbMap, "engine", "sqlite"),
	}
	cnf.Database = dbConfig

	// Build LocaleConfig
	localeMap, ok := (*m)["locale"].(map[string]interface{})
	if !ok {
		localeMap = make(map[string]interface{})
	}
	localeConfig := LocaleConfig{
		Default:  src.GetOrDefault(&localeMap, "default", "en_US"),
		Selected: src.GetOrDefault(&localeMap, "selected", "en_US"),
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
	err = os.WriteFile(path, d, 0600)
	if err != nil {
		return err
	}
	return nil
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
