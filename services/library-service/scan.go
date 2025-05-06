package library_service

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

var videoExtensions = []string{"mp4", "mkv", "webm", "avi", "mov", "wmv", "mpg", "mpeg", "divx"}

type LibraryScanner struct {
	Libraries []Library
}

type Library struct {
	Path       string
	Extensions []string
	Files      map[string]FileInfo
}

type FileInfo struct {
	Name      string
	Extension string
	Path      string
}

func (scanner *LibraryScanner) Setup() {
	// Load Libraries
	lib1Path, _ := filepath.Abs("./example/movies")
	lib1 := Library{lib1Path, videoExtensions, map[string]FileInfo{}}
	scanner.Libraries = append(scanner.Libraries, lib1)

	lib2Path, _ := filepath.Abs("./example/tv")
	lib2 := Library{lib2Path, videoExtensions, map[string]FileInfo{}}
	scanner.Libraries = append(scanner.Libraries, lib2)
}

func (scanner *LibraryScanner) Scan() {
	for _, library := range scanner.Libraries {
		library.Scan()
	}
}

func (lib *Library) Scan() {
	var matchingFiles []string

	fmt.Printf("\tScanning => %s\n", lib.Path)
	err := filepath.WalkDir(lib.Path, func(path string, de fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if de.IsDir() {
			return nil // Skip directories
		}

		ext := filepath.Ext(de.Name())[1:]
		for _, wantedExt := range lib.Extensions {
			if ext == wantedExt {
				matchingFiles = append(matchingFiles, path)
				break
			}
		}

		return nil
	})

	if err != nil {
		fmt.Errorf("error scanning directory: %w", err)
	}

	fmt.Printf("\tFound => %v\n", matchingFiles)
}
