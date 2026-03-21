package commands

import (
	"bufio"
	"cadagen/microservices/fs/src/utils"
	"io"
	"os"
	"strings"
)

type ReadCliResponse string

func Read(filePath string, start, end int) utils.CliResponse[ReadCliResponse] {
	file, err := os.Open(filePath)
	if err != nil {
		utils.Throw(err)
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	lines := make([]string, 0)
	currentLine := 1

	for {
		if end > 0 && currentLine > end {
			break
		}

		line, err := reader.ReadString('\n')

		if currentLine >= start {
			lines = append(lines, line)
		}

		if err != nil {
			if err == io.EOF {
				break
			}

			utils.Throw(err)
		}

		currentLine++
	}

	return utils.NewCliResponse(ReadCliResponse(strings.Join(lines, "\n")))
}
