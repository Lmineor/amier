package core

import (
	"fmt"
	"ziyue/global"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Viper(path ...string) *viper.Viper {
	// var config string
	// if len(path) == 0 {
	// 	flag.StringVar(&config, "c", "", "choose config file.")
	// 	flag.Parse()

	// }
	config := path[0]
	v := viper.New()
	v.SetConfigFile(config)
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed :", e.Name)
		if err := v.Unmarshal(&global.Z_CONFIG); err != nil {
			fmt.Println(err)
		}
	})

	if err := v.Unmarshal(&global.Z_CONFIG); err != nil {
		fmt.Println(err)
	}
	return v
}
