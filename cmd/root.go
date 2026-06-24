/*
Copyright © 2026 eamat. <eamat.dot@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ~/.config/"+APP_NAME+"/config.toml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" { // --config 指定時
		viper.SetConfigFile(cfgFile)
	} else {
		configDir, err := getConfigDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(filepath.Join(configDir, APP_NAME))
		viper.SetConfigType("toml")
		viper.SetConfigName("config")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

// getConfigDir は、XDG Base Directory Specification に従って設定ディレクトリを取得する
func getConfigDir() (string, error) {
	if dir := os.Getenv("XDG_CONFIG_HOME"); dir != "" {
		return dir, nil
	}

	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	return filepath.Join(home, ".config"), nil
}
