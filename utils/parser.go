package utils

import (
	"fmt"
	"strings"
)

func ParseInfoCommandResponse(input string) map[string]map[string]string {
	var sectionName string
	out := make(map[string]map[string]string)
	lines := strings.Split(input, "\r\n")
	for _, line := range lines {
		fmt.Println(line)
		if len(line) > 0 {
			if string(line[0]) == "#" {
				sectionName = line
				out[sectionName] = make(map[string]string)
			} else {
				keyValue := strings.Split(line, ":")
				if len(keyValue) == 2 {
					out[sectionName][keyValue[0]] = keyValue[1]
				} else {
					out[sectionName][keyValue[0]] = ""
				}

			}
		}

	}
	return out

}
