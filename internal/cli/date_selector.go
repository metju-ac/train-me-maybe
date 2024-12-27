package cli

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"time"
)

func SelectDate() (string, error) {
	var selectedDate string
	for {
		prompt := &survey.Input{
			Message: "Enter the departure date (YYYY-MM-DD):",
		}
		err := survey.AskOne(prompt, &selectedDate)
		if err != nil {
			return "", err
		}

		_, err = time.Parse("2006-01-02", selectedDate)
		if err == nil {
			selectedTime, _ := time.Parse("2006-01-02", selectedDate)
			if !selectedTime.Before(time.Now().Truncate(24 * time.Hour)) {
				break
			}
			fmt.Println("The date cannot be in the past. Please try again.")

		} else {
			fmt.Println("Invalid date format. Please try again.")
		}
	}
	return selectedDate, nil
}