package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

var Config = viper.New()

func Init() *viper.Viper {
    Config.SetConfigName("config")
    Config.SetConfigType("toml")

    home, err := os.UserHomeDir()
    if err != nil {
        fmt.Println("Error getting home directory:", err)
        os.Exit(1)
    }

    path := home + "/.config/moco"
    filename := "config.toml"

    Config.AddConfigPath(path)

    if _, err := os.Stat(path + "/" + filename); os.IsNotExist(err) {
        os.MkdirAll(path, 0755)
        f, _ := os.Create(path + "/" + filename)
        f.Close()
    }

    err = Config.ReadInConfig()
    if err != nil {
        fmt.Println("Error reading config:", err)
    }

    return Config
}
