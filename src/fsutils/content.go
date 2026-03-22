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

			if j == 0 && !strings.HasSuffix(line, ql) {
				match = false
				break
			} else if j == len(queryLines)-1 && !strings.HasPrefix(line, ql) {
				match = false
				break
			} else if j > 0 && j < len(queryLines)-1 && line != ql {
				match = false
				break
			}
		}

		if match {
			occurences = append(occurences, i+1)
		}
	}

	return occurences, nil
}
