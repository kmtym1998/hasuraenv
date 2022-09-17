package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// RootCmd is root command
var (
	RootCmd = &cobra.Command{
		Use:          "hasuraenv",
		Short:        "Manage multiple hasura-cli versions",
		Long:         "Manage multiple hasura-cli versions. Run 'hasuraenv --help' for usage",
		SilenceUsage: true,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hi")
		},
	}
)

// コマンド実行時に最初に呼ばれる初期化処理
func init() {
	// フラグの定義
	// 第1引数: フラグ名、第2引数: 省略したフラグ名
	// 第3引数: デフォルト値、第4引数: フラグの説明
}
