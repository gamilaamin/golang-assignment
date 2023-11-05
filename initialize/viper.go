package initialize

import (
	"fmt"
	"golang-assignment/global"
	"log"

	"github.com/spf13/viper"
)

func InitViper(path string) *viper.Viper {

	v := viper.New()
	v.SetConfigFile(path)

	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	} else {
		log.Printf("config file loaded successfully\n")
	}
	v.WatchConfig()
	if err := v.Unmarshal(&global.CONFIG); err != nil {
		fmt.Println(err)
	}
	return v
}
