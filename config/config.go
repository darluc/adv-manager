package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Config struct {
	DbFile       string `json:"database_file"`
	StaticDir    string `json:"static_dir"`
	CookieSecret string `json:"cookie_secret"`
}

type CfgFunc func(config *Config) error

func LoadConfig(loaders ...CfgFunc) *Config {
	config := new(Config)
	for _, loader := range loaders {
		if err := loader(config); err != nil {
			panic(err)
		}
	}
	return config
}

func FromJsonFile(filepath string) CfgFunc {
	return func(config *Config) error {
		if content, err := ioutil.ReadFile(filepath); err == nil {
			err = json.Unmarshal(content, config)
			return err
		} else {
			return err
		}
	}
}

func FromEnv(config *Config) error {
	if db := os.Getenv("ADV_DATABASE"); db != "" {
		config.DbFile = db
	}
	if secret := os.Getenv("ADV_COOKIE_SECRET"); secret != "" {
		config.CookieSecret = secret
	}
	if staticDir := os.Getenv("STATIC_DIR"); staticDir != "" {
		config.StaticDir = staticDir
	}
	return nil
}
