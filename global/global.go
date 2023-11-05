package global

import (
	"golang-assignment/config"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	VP     *viper.Viper
	DB     *gorm.DB
	CONFIG config.Server
)
