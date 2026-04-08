package analyzer

import (
	"io/fs"
	"path/filepath"
)

func shouldCollect(d fs.DirEntry) bool {
	if d.IsDir() {
		return false
	}
	return d.Name()[0] != '.'
}

func listFiles(dir string, recursive bool) ([]string, error) {
	var files []string

	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() && d.Name() == ".git" {
			return filepath.SkipDir
		}

		if !recursive && d.IsDir() && path != dir {
			return filepath.SkipDir
		}

		if shouldCollect(d) {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}
