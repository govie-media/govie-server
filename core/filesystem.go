package core

import (
	"os"
	"path"
	"path/filepath"
	"strings"
)

func DirectoryWalk(root, validDirectory string) ([]string, error) {
	var directories []string

	err := filepath.Walk(root, func(dir string, info os.FileInfo, err error) error {
		if info.IsDir() {
			directory := strings.Split(dir, "/")
			if directory[len(directory)-1] == validDirectory && strings.Count(dir, validDirectory) == 1 {
				directories = append(directories, path.Dir(dir))
			}
		}
		return nil
	})

	return directories, err
}
