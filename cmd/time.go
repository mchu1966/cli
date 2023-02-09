package cmd

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/mchu1966/cli/internal/timer"
	"github.com/spf13/cobra"
)

var calculateTime string
var duration string

func init() {
	timeCmd.AddCommand(nowTimeCmd)
	timeCmd.AddCommand(calculateTimeCmd)

	calculateTimeCmd.Flags().StringVarP(&calculateTime, "calculate", "c", "", ` target time, accept timestamp and format time string `)
	calculateTimeCmd.Flags().StringVarP(&duration, "duration", "d", "", ` duration, accept time unit of 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'`)
}

var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "time format handling",
	Long:  "time format handling",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var nowTimeCmd = &cobra.Command{
	Use:   "now",
	Short: "get current time",
	Long:  "get current time",
	Run: func(cmd *cobra.Command, args []string) {
		nowTime := timer.GetNowTime()
		log.Printf("Current time: %s, %d", nowTime.Format(time.RFC3339), nowTime.Unix())
	},
}

var calculateTimeCmd = &cobra.Command{
	Use:   "calc",
	Short: "Calculate time plus duration",
	Long:  "Calculate time plus duration",
	Run: func(cmd *cobra.Command, args []string) {
		var currentTimer time.Time
		var layout = "2006-01-02 15:04:05"
		if calculateTime == "" {
			currentTimer = timer.GetNowTime()
		} else {
			var err error
			switch strings.Count(calculateTime, ":") {
			case 0:
				layout = "2006-01-02"
			case 1:
				layout = "2006-01-02 15:04"
			}
			currentTimer, err = time.Parse(layout, calculateTime)
			if err != nil {
				t, _ := strconv.Atoi(calculateTime)
				currentTimer = time.Unix(int64(t), 0)
			}
		}
		t, err := timer.GetCalculateTime(currentTimer, duration)
		if err != nil {
			log.Fatalf("timer.GetCalculateTime err: %v", err)
		}

		log.Printf("Output: %s, %d", t.Format(layout), t.Unix())
	},
}
