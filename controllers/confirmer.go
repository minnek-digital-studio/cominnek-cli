package controllers

import (
	"fmt"
)

func Confirm(msg string, defaultAnswer ...string) bool {
	var answer string
	yesOrNo := yesOrNo(defaultAnswer...)

	fmt.Println(msg + yesOrNo)
	fmt.Scanln(&answer)

	if len(defaultAnswer) > 0 && answer == "" {
		answer = defaultAnswer[0]
	}

	for answer != "y" && answer != "n" && answer != "yes" && answer != "no" {
		fmt.Println("Please enter yes or no")
		fmt.Scanln(&answer)
	}

	return answer == "y"
}

func yesOrNo(defaultAnswer ...string) string {
	yesText := "y"
	noText := "n"

	if len(defaultAnswer) > 0 {
		if defaultAnswer[0] == "yes" || defaultAnswer[0] == "y" {
			yesText = "Y"
		}

		if defaultAnswer[0] == "no" || defaultAnswer[0] == "n" {
			noText = "N"
		}
	}

	return fmt.Sprintf(" [%s/%s]", yesText, noText)
}
