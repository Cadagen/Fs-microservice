package fsutils

import (
	"bytes"
	"os"
	"strings"
)

func FindOccurences(path string, text string) ([]int, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	query := []byte(text)
	if len(query) == 0 {
		return []int{}, nil
	}

	occurences := make([]int, 0)
	lines := bytes.Split(data, []byte("\n"))
	queryLines := strings.Split(text, "\n")

	for i := 0; i <= len(lines)-len(queryLines); i++ {
		match := true

		for j, ql := range queryLines {
			line := string(lines[i+j])

			if len(queryLines) == 1 {
				if !strings.Contains(line, ql) {
					match = false
				}
			} else if j == 0 {
				if !strings.HasSuffix(line, ql) {
					match = false
				}
			} else if j == len(queryLines)-1 {
				if !strings.HasPrefix(line, ql) {
					match = false
				}
			} else {
				if line != ql {
					match = false
				}
			}

			if !match {
				break
			}
		}

		if match {
			occurences = append(occurences, i+1)
		}
	}

	return occurences, nil
}
