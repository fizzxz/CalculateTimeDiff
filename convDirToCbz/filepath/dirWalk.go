package filepath

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func WalkDir_FindSubDirsInCurrDir(rootDir string) []string {
	var dir []string
	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			rootDirCount := strings.Count(rootDir, "\\")
			count := strings.Count(path, "\\")
			rootDirCount++
			if rootDirCount == count {
				dir = append(dir, path)
			} else {
				return nil
			}
		}
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		return nil
	})
	if err != nil {
		fmt.Printf("walk error [%v]\n", err)
	}
	return dir
}

func WalkDir_FindFiles(rootDir string) []string {
	// Walk directory, adding the files we want to add

	var files []string
	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		return nil
	})
	if err != nil {
		fmt.Printf("walk error [%v]\n", err)
	}
	return files
}

func FindDir(rootDir string) []fs.FileInfo {
	files, err := ioutil.ReadDir(rootDir)
	if err != nil {
		fmt.Println(err)
	}
	return files
}
