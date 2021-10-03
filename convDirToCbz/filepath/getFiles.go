package filepath

import (
	"fmt"
	"os"
)

var err error = nil

func GetFileOs(files []string) map[string]os.FileInfo {
	filesOS := make(map[string]os.FileInfo)

	for _, file := range files {
		filesOS[file], err = os.Stat(file)
		if err != nil {
			fmt.Print(err)
		}
	}
	return (filesOS)
}

func GetSubDirs(rootDir string) []string {
	subDirs := WalkDir_FindSubDirs(rootDir)
	_, subDirs = subDirs[0], subDirs[1:]
	return subDirs
}
