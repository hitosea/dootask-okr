package command

import (
	"dootask-okr/config"
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCommand.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "查看版本号",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Version: %s\nCommitSHA: %s\n", config.Version, config.CommitSHA)
	},
}
