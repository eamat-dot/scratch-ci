/*
Copyright © 2026 eamat. <eamat.dot@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// sampleCmd represents the sample command
var sampleCmd = &cobra.Command{
	Use:   "sample",
	Short: "コマンドのサンプル",
	Long: `このファイルはコマンドのサンプルです。コマンドを追加する際の参考にしてください。
削除しても問題ありません。コマンドの追加方法は以下の通りです。

>  cobra-cli add <command-name>
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("sample called")
	},
}

func init() {
	rootCmd.AddCommand(sampleCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sampleCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sampleCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
