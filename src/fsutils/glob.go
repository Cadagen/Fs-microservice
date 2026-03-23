package fsutils

import (
	"io/fs"
	"path/filepath"
)

func Glob(dirPath string, pattern string, excludedDirs []string) []string {
	var matches []string

	filepath.WalkDir(dirPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}

		if d.IsDir() {
			for _, dir := range excludedDirs {
				if d.Name() == dir {
					return filepath.SkipDir
				}
			}

			return nil
		}

		matched, _ := filepath.Match(pattern, d.Name())
		if matched {
			matches = append(matches, path)
		}

		return nil
	})

	return matches
}
