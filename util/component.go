package util

import (
	"fmt"
	gola "github.com/paulus-otto-harman/golang-module"
	"project/config"
)

func renderTimeout() string {
	if config.SessionTimeout != 0 {
		return fmt.Sprintf("\n%s\n", gola.Tf(gola.Bold, fmt.Sprintf("Timeout %d seconds", config.SessionTimeout), gola.LightYellow))
	}
	return fmt.Sprintf("\n%s\n", gola.Tf(gola.Bold, "Timeout Disabled", gola.LightGreen))
}

func H1(title string) {
	gola.ClearScreen()

	fmt.Println(renderTimeout())

	fmt.Printf("%s\n\n", gola.Tf(gola.Bold, fmt.Sprintf("%s %s - %s %s", "===", config.AppName, title, "==="), gola.LightBlue))
}

func Select(options map[string]interface{}, label string) string {
	for optionValue, option := range options {
		fmt.Printf("[%s] %v", optionValue, option)
	}
	selected, err := gola.ToString(gola.Input(gola.Args(gola.P("label", label))))
	if err != nil {

	}
	return selected
}

func ContinueOrReturn() int {
	if cont, _ := gola.ToString(gola.Input(gola.Args(gola.P("label", "Press Enter to continue or [0] to return :")))); cont == "0" {
		return 0
	}
	return -1
}

func ContinueOrExit() int {
	if cont, _ := gola.ToString(gola.Input(gola.Args(gola.P("label", "Press Enter to continue or [0] to exit :")))); cont == "0" {
		return 0
	}
	return -1
}

func WaitToReturn() int {
	gola.Wait("Press Enter to return")
	return 0
}
