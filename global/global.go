package global

import (
	"go.uber.org/zap"
	"ziyue/config"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	Z_CONFIG config.Server
	Z_VP     *viper.Viper
	Z_LOG    *zap.Logger
	Z_DB     *gorm.DB
)
