package initialize

import (
	"fmt"
	"os"
	"ziyue/global"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Gorm() *gorm.DB {
	switch global.Z_CONFIG.System.DbType {
	case "mysql":
		return GormMysql()
	default:
		return GormMysql()
	}
}

func GormMysql() *gorm.DB {
	m := global.Z_CONFIG.Mysql
	dsn := m.Username + ":" + m.Password + "@tcp(" + m.Path + ")/" + m.Dbname + "?" + m.Config
	fmt.Println(dsn)
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         255,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	// if db, err := gorm.Open(mysql.New(mysqlConfig), gormConfig(m.LogMode)); err != nil {
	if db, err := gorm.Open(mysql.New(mysqlConfig)); err != nil {
		fmt.Println("Mysql error")
		os.Exit(0)
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}

}

// func gormConfig(mod bool) *gorm.Config {
// 	var config = &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}
// 	switch global.ServerConfig.Mysql.LogZap {
// 	case "silent", "Silent":
// 		config.Logger = internal.Default.LogMode(logger.Silent)
// 	case "error", "Error":
// 		config.Logger = internal.Default.LogMode(logger.Error)
// 	case "warn", "Warn":
// 		config.Logger = internal.Default.LogMode(logger.Warn)
// 	case "info", "Info":
// 		config.Logger = internal.Default.LogMode(logger.Info)
// 	case "zap", "Zap":
// 		config.Logger = internal.Default.LogMode(logger.Info)
// 	default:
// 		if mod {
// 			config.Logger = internal.Default.LogMode(logger.Info)
// 			break
// 		}
// 		config.Logger = internal.Default.LogMode(logger.Silent)
// 	}
// 	return config
// }
