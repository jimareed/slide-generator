package slides

import (
	"strings"
)


func Execute(path string) (string, error) {

	output := strings.ToUpper(path)

	return output, nil
}