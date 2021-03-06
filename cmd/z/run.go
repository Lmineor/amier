package z

import (
	"os"
	"ziyue/utils"

	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "run ziyue server.",
	Long: `
	The "run" command is used for running ziyue server with hot-compiled-like feature,     
	which compiles and runs the go codes asynchronously when codes change.
`,
	Run: func(cmd *cobra.Command, args []string) {
		w := utils.NewWatch()
		t := utils.NewT()
		cmdPath, _ := os.Getwd()
		go w.Watch(cmdPath, t)
		t.RunTask()
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
