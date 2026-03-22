package commands

import (
	"cadagen/microservices/fs/src/utils"
	"io/fs"
	"path/filepath"
)

type GlobCliResponse []string

func Glob(dirPath string, pattern string) utils.CliResponse[GlobCliResponse] {
	var matches []string

	filepath.WalkDir(dirPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}

		if d.IsDir() {
			// TODO: Skip ignored folders by name
			// if d.Name() == ".git" || d.Name() == "node_modules" {
			//	 return filepath.SkipDir
			// }

			return nil
		}

		matched, _ := filepath.Match(pattern, d.Name())
		if matched {
			matches = append(matches, path)
		}

		return nil
	})

	return utils.NewCliResponse(GlobCliResponse(matches))
}
