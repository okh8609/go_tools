package cmd

import (
	"log"
	"strings"

	"github.com/okh8609/go_tools/internal/word"
	"github.com/spf13/cobra"
)

const (
	ModeUpper                      = iota + 1 // 全部轉大寫
	ModeLower                                 // 全部轉小寫
	ModeUnderscoreToUpperCamelCase            // 下劃線轉大寫駝峰
	ModeUnderscoreToLowerCamelCase            // 下線線轉小寫駝峰
	ModeCamelCaseToUnderscore                 // 駝峰轉下劃線
)

var info = strings.Join([]string{
	"該子命令支持各種單詞格式轉換，模式如下：",
	"1：全部轉大寫",
	"2：全部轉小寫",
	"3：下劃線轉大寫駝峰",
	"4：下劃線轉小寫駝峰",
	"5：駝峰轉下劃線",
}, "\n")

var mode int8
var ssss string

var wordCmd = &cobra.Command{
	Use:   "word",
	Short: "單字格式轉換",
	Long:  info,
	Run: func(cmd *cobra.Command, args []string) {
		var result string
		switch mode {
		case ModeUpper:
			result = word.ToUpper(ssss)
		case ModeLower:
			result = word.ToLower(ssss)
		case ModeUnderscoreToUpperCamelCase:
			result = word.Underscore_To_UpperCamelCase(ssss)
		case ModeUnderscoreToLowerCamelCase:
			result = word.Underscore_To_LowerCamelCase(ssss)
		case ModeCamelCaseToUnderscore:
			result = word.CamelCase_To_Underscore(ssss)
		default:
			log.Fatalf("暫不支持該轉換模式，請執行 help word 查看幫助文檔")
		}

		log.Printf("輸出結果: %s", result)
	},
}

func init() {
	wordCmd.Flags().StringVarP(&ssss, "str", "s", "Test...", "請輸入單字內容")
	wordCmd.Flags().Int8VarP(&mode, "mode", "m", 0, "請輸入單字轉換模式")
}
