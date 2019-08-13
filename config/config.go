package config

import (
	"encoding/json"
	"io/ioutil"
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
