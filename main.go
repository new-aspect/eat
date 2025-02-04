package main

import (
	"fmt"
	"os"

	"eat/cmd"
)

func main() {
	rootCmd := cmd.RootCmd

	// Add global flags
	rootCmd.PersistentFlags().StringP("cpu_usage", "c", "0", "How many cpu would you want eat")
	rootCmd.PersistentFlags().StringP("memory_usage", "m", "0m", "How many memory would you want eat(GB)")
	// such as "300ms", "1.5h", "2h45m". (unit: "ns", "us" (or "µs"), "ms", "s", "m", "h")
	rootCmd.PersistentFlags().StringP("time_deadline", "t", "0", "Deadline to quit eat process")
	// same unit as time_deadline
	rootCmd.PersistentFlags().StringP("memory_refresh_interval", "r", "5m", "How often to trigger a refresh to prevent the ate memory from being swapped out")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
