package adventofcode

import (
	"fmt"
	"io/ioutil"
)

func ReadInputData(path string) string {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return fmt.Sprintf("Cannot reach input file using relative path: %s\n", path)
	}

	return string(bytes)
}
