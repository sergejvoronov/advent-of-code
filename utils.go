package adventofcode

import (
	"fmt"
	"io/ioutil"
)

func ReadInputData(path string) string {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("Cannot reach input file using relative path: %s\n", path)
		return ""
	}
	return string(bytes)
}
