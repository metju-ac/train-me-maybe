package main

import (
	"fmt"
	"os"

	"github.com/metju-ac/train-me-maybe/internal/config"
	"github.com/metju-ac/train-me-maybe/internal/notification"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	notification.EmailNotification(&config.Smtp, 1, 2, []int64{1, 2, 3}, []string{"REGULAR"})
}
