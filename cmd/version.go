/*
Copyright © 2026 eamat. <eamat.dot@gmail.com>
*/
package cmd

import (
	"fmt"
	"runtime/debug"
	"time"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of the CLI",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(APP_NAME + " version " + VersionString())
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

func VersionString() string {
	version := "dev"
	commit := "unknown"
	buildDate := "unknown"

	if info, ok := debug.ReadBuildInfo(); ok {
		version = info.Main.Version

		for _, s := range info.Settings {
			switch s.Key {
			case "vcs.revision":
				if len(s.Value) >= 7 {
					commit = s.Value[:7]
				} else {
					commit = s.Value
				}

			case "vcs.time":
				if t, err := time.Parse(time.RFC3339, s.Value); err == nil {
					buildDate = t.Format("2006-01-02")
				}
			}
		}
	}

	return fmt.Sprintf(
		"%s (%s) build %s",
		version,
		buildDate,
		commit,
	)
}
