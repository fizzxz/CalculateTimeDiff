package filepath

import (
	"fmt"
	"os"
)

func getFileOs(files []string) map[string]os.FileInfo {
	filesOS := make(map[string]os.FileInfo)

	for _, file := range files {
		filesOS[file], err = os.Stat(file)
		if err != nil {
			fmt.Print(err)
		}
	}
	return (filesOS)
}

func getSubDirs(rootDir string) []string {
	subDirs := walkDir_FindSubDirs(rootDir)
	_, subDirs = subDirs[0], subDirs[1:]
}
