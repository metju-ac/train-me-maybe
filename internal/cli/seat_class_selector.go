package cli

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/metju-ac/train-me-maybe/openapi"
)

func formatSeatClass(seatClass openapi.SeatClass) string {
	return fmt.Sprintf("%s - %s [%s]", seatClass.Title, seatClass.ShortDescription, seatClass.Key)
}

func SelectSeatClasses(seatClasses []openapi.SeatClass) ([]string, error) {
	var selectedSeatClasses []string

	formattedSeatClasses := make([]string, 0, len(seatClasses))
	seatClassMap := make(map[string]string)
	for _, seatClass := range seatClasses {
		formatted := formatSeatClass(seatClass)
		formattedSeatClasses = append(formattedSeatClasses, formatted)
		seatClassMap[formatted] = seatClass.Key
	}

	selectPrompt := &survey.MultiSelect{
		Message: "Select seat classes:",
		Options: formattedSeatClasses,
	}
	err := survey.AskOne(selectPrompt, &selectedSeatClasses)
	if err != nil {
		return nil, err
	}

	var selectedKeys []string
	for _, selected := range selectedSeatClasses {
		selectedKeys = append(selectedKeys, seatClassMap[selected])
	}

	return selectedKeys, nil
}
