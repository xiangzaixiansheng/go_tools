package config_util

import (
	"fmt"
	"testing"
)

func TestConfig(t *testing.T) {
	InitConfig()
	fmt.Println("Redis.Url", Config.Redis.Url)
	fmt.Println("AppConfig.Name", Config.App.Name)

}
