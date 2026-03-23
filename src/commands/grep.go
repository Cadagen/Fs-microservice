package commands

import (
	"cadagen/microservices/fs/src/fsutils"
	"cadagen/microservices/fs/src/utils"
)

type GrepCliResponseItem struct {
	Path    string `json:"path"`
	Offsets []int  `json:"offsets"`
}

type GrepCliResponse []GrepCliResponseItem

func Grep(dirPath string, text string, filePattern string, excludedDirs []string) utils.CliResponse[GrepCliResponse] {
	files := fsutils.Glob(dirPath, filePattern, excludedDirs)

	output := make(GrepCliResponse, 0)

	for _, file := range files {
		occurences, err := fsutils.FindOccurences(file, text)
		if err != nil {
			continue
		}

		if len(occurences) == 0 {
			continue
		}

		output = append(output, GrepCliResponseItem{
			Path:    file,
			Offsets: occurences,
		})
	}

	return utils.NewCliResponse(output)
}
