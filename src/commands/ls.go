package commands

import (
	"cadagen/microservices/fs/src/utils"
	"os"
)

type LsCliResponseItemType string

const (
	LsCliResponseItemTypeFile      LsCliResponseItemType = "file"
	LsCliResponseItemTypeDirectory LsCliResponseItemType = "directory"
)

type LsCliResponseItem struct {
	Name string                `json:"name"`
	Type LsCliResponseItemType `json:"type"`
}

func Ls(dirPath string) utils.CliResponse[[]LsCliResponseItem] {
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		utils.Throw(err)
	}

	output := make([]LsCliResponseItem, 0)

	for _, entry := range entries {
		output = append(output, LsCliResponseItem{
			Name: entry.Name(),
			Type: utils.Ifs(
				entry.IsDir(),
				LsCliResponseItemTypeDirectory,
				LsCliResponseItemTypeFile,
			),
		})
	}

	return utils.NewCliResponse(output)
}
