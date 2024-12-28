package main

import (
	"fmt"
	"os"
	"time"

	"github.com/metju-ac/train-me-maybe/internal/config"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println("Poll interval in seconds:", config.General.PollInterval)

	// start a timer
	start := time.Now()

	time.Sleep(time.Duration(config.General.PollInterval) * time.Second)

	// stop the timer
	elapsed := time.Since(start)

	fmt.Println("Time elapsed:", elapsed)
}
