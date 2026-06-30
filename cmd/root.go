package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/adrg/xdg"
)

const APP_NAME string = "__REPLACE_MODULE_NAME__"

var cfgFile string

// rootCmd は、サブコマンドなしで呼び出された場合の基本コマンドを表します。
var rootCmd = &cobra.Command{
	Use:   APP_NAME,
	Short: "アプリケーションの簡単な説明",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute は、すべての子コマンドをルート コマンドに追加し、フラグを適切に設定します。
// これは main.main() によって呼び出されます。これは rootCmd に対して 1 回だけ実行する必要があります。
func Execute() {
	rootCmd.Version = VersionString()
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ~/.config/"+APP_NAME+"/config.(toml/yaml/json))")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" { // --config 指定時
		viper.SetConfigFile(cfgFile)
	} else {
		// XDG_CONFIG_HOME の設定があれば、そこを優先
		// 無ければ linux: ~/.config || windows: %AppData%
		viper.AddConfigPath(filepath.Join(xdg.ConfigHome, APP_NAME))

		// Windows の場合は、%USERPROFILE%\.config も探してみる
		if runtime.GOOS == "windows" {
			if home, err := os.UserHomeDir(); err == nil {
				viper.AddConfigPath(filepath.Join(home, ".config", APP_NAME))
			}
		}

		// viper.SetConfigType("toml")
		// 拡張子はviperが自動で補完するため付けない (実際のファイル名は config.toml などにする)
		viper.SetConfigName("config")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
