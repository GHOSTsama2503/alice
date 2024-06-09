package modules

import "fmt"

func NewSwitchButton(enabled bool, text string) string {
	if enabled {
		return fmt.Sprintf("%s %s", "☑️", text)
	}

	return fmt.Sprintf("%s %s", "✖️", text)
}
