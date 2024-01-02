package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var Config = viper.New()

func Init() *viper.Viper {
    Config.SetConfigName("config")
    Config.SetConfigType("toml")

    Config.AddConfigPath("$HOME/.config/moco")

    err := Config.ReadInConfig()
    if err != nil {
        fmt.Println("Error reading config:", err)
    }

    return Config
}
