package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type ConfigType struct {
	Env    string
	Server struct {
		Port int
	}
	Database struct {
		Postgres struct {
			Host     string
			User     string
			Password string
			DBname   string
			Port     string
			SSLmode  string
		}
	}
	Bcrypt struct {
		Key  string
		Cost int
	}
}

var Config ConfigType

func init() {
	viper.SetEnvPrefix("go.bank")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
	Config.Env = viper.GetString("env")

	if Config.Env != "" && Config.Env != "prod" {
		viper.SetConfigName("config_" + Config.Env)
	} else {
		viper.SetConfigName("config")
	}
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./assets/config/")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("[Error] Loading config failed: ", err)
		panic(err)
	}

	if err := viper.Unmarshal(&Config); err != nil {
		fmt.Println("[Error] Loading config failed: ", err)
		panic(err)
	}
}
