package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

type config struct {
	Mysql struct {
		DBMS string
		USER string
		PASS string
		HOST string
		PORT string
		DBNAME string
		OPTION string
	}
}

var C config

func LoadConfig() {
	Conf := &C
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(filepath.Join("config"))

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Load config error")
		fmt.Println(err)
		os.Exit(1)
	}

	if err := viper.Unmarshal(&Conf); err != nil {
		fmt.Println("config Unmarshal error")
		fmt.Println(err)
		os.Exit(1)
	}
}