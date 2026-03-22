package commands

import (
	"cadagen/microservices/fs/src/fsutils"
	"cadagen/microservices/fs/src/utils"
)

type GlobCliResponse []string

func Glob(dirPath string, pattern string) utils.CliResponse[GlobCliResponse] {
	matches := fsutils.Glob(dirPath, pattern)

	return utils.NewCliResponse(GlobCliResponse(matches))
}
