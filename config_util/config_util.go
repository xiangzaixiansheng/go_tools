package config_util

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/creasty/defaults"
	"gopkg.in/yaml.v2"
)

type Configuration struct {
	App   AppConfig   `yaml:"app-config"`
	Redis RedisConfig `yaml:"redis-config"`
}

type RedisConfig struct {
	Url      string `yaml:"url"`
	Db       int    `yaml:"db"`
	Password string `yaml:"password"`
	RDBPath  string `yaml:"rdb_path"`
}

type AppConfig struct {
	Name          string `yaml:"name"`
	DbPassword    string `yaml:"db_password"`
	AdminUser     string `yaml:"admin_user"`
	AdminPassword string `yaml:"admin_passowrd"`
	Plugin        string `yaml:"plugin" default:"default"`
}

var Config Configuration

func InitConfig() {
	absPath, err := filepath.Abs("./configuration.yaml")
	fmt.Println("absPath", absPath)
	if err != nil {
	}
	yamlFile, err := ioutil.ReadFile(absPath)
	if err != nil {
		return
	}
	yaml.Unmarshal(yamlFile, &Config)
	setDefaults(&Config)
}

func setDefaults(v interface{}) {
	if err := defaults.Set(v); err != nil {
		panic(err)
	}
}
