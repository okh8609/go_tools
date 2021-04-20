package cmd

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/okh8609/go_tools/internal/timer"
	"github.com/spf13/cobra"
)

var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "時間格式處理",
	Long:  "時間格式處理",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var nowTimeCmd = &cobra.Command{
	Use:   "now",
	Short: "取得目前時間",
	Long:  "取得目前時間",
	Run: func(cmd *cobra.Command, args []string) {
		nowTime := timer.GetNowTime()
		log.Printf("输出结果: %s, %d", nowTime.Format("2006-01-02 15:04:05"), nowTime.Unix())
		log.Printf("输出结果: %s, %d", nowTime.Format(time.RFC3339), nowTime.Unix())

	},
}

var calcTime string
var duration string

var calcTimeCmd = &cobra.Command{
	Use:   "calc",
	Short: "計算所需時間",
	Long:  "計算所需時間",
	Run: func(cmd *cobra.Command, args []string) {
		var currTime time.Time
		var layout = ""
		if calcTime == "" {
			currTime = timer.GetNowTime()
			layout = "2006-01-02 15:04:05"
		} else {
			space := strings.Count(calcTime, " ")
			if space == 0 {
				layout = "2006-01-02"
			}
			if space == 1 {
				layout = "2006-01-02 15:04"
			}

			var err error
			currTime, err = time.Parse(layout, calcTime)

			if err != nil { // timestamp
				t, _ := strconv.Atoi(calcTime)
				currTime = time.Unix(int64(t), 0)
				layout = "2006-01-02 15:04:05"
			}
		}

		t, err := timer.GetCalcTime(currTime, duration)
		if err != nil {
			log.Fatalf("timer.GetCalcTime err: %v", err)
		}

		log.Printf("输出结果: %s, %d", t.Format(layout), t.Unix())
	},
}

func init() {
	timeCmd.AddCommand(nowTimeCmd)
	timeCmd.AddCommand(calcTimeCmd)
	calcTimeCmd.Flags().StringVarP(&calcTime, "calculate", "c", "", `需要計算的時間，有效單位為時間戳或已格式化後的時間`)
	calcTimeCmd.Flags().StringVarP(&duration, "duration", "d", "", `持續時間，有效時間單位為"ns", "us" (or "µs"), "ms", "s", "m", "h"`)
}
