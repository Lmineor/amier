package global

import (
	"ziyue/config"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	ZConfig config.Server
	ZVP     *viper.Viper
	ZDB     *gorm.DB
)
