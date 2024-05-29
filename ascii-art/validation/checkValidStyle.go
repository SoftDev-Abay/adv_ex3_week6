package validation

import "fmt"

//  valid styles are "standard", "thinkertoy", "shadow"

var validStyles = map[string]bool{
	"standard":   true,
	"thinkertoy": true,
	"shadow":     true,
}

func CheckValidStyle(styleName string) error {
	if !validStyles[styleName] {
		return fmt.Errorf("Invalid style name, only 'standard', 'thinkertoy', 'shadow' are allowed.")
	}
	return nil
}
