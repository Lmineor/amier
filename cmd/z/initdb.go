package z

import (
	"github.com/gookit/color"
	"github.com/spf13/cobra"

	"ziyue/core"
	_ "ziyue/core"
	"ziyue/global"
)

// initdbCmd represents the initdb command
var initdbCmd = &cobra.Command{
	Use:   "initdb",
	Short: "初始化数据",
	Long: `初始化数据适配数据库情况: 
1. mysql完美适配,
2. postgresql不能保证完美适配,
3. sqlite未适配,
4. sqlserver未适配`,
	Run: func(cmd *cobra.Command, args []string) {
		frame, _ := cmd.Flags().GetString("frame")
		path, _ := cmd.Flags().GetString("path")
		global.Z_VP = core.Viper(path)
		// global.GVA_LOG = core.Zap() // 初始化zap日志库
		Mysql.CheckDatabase()
		Mysql.CheckUtf8mb4()
		Mysql.Info()
		Mysql.Init()
		switch frame {
		case "gin":
			if global.Z_CONFIG.System.DbType == "mysql" {
				Mysql.AutoMigrateTables()
				// Mysql.InitData()
			}
		case "gf":
			color.Info.Println("gf功能开发中")
		}
	},
}

func init() {
	rootCmd.AddCommand(initdbCmd)
	initdbCmd.Flags().StringP("path", "p", "./config.yaml", "自定配置文件路径(绝对路径)")
	initdbCmd.Flags().StringP("frame", "f", "gin", "可选参数为gin,gf")
	initdbCmd.Flags().StringP("type", "t", "mysql", "可选参数为mysql,postgresql,sqlite,sqlserver")
}
