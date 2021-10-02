package filepath

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"strings"

	"github.com/karrick/godirwalk"
)

func WalkDir_FindSubDirs(rootDir string) []string {
	var dir []string
	err := godirwalk.Walk(rootDir, &godirwalk.Options{
		Callback: func(osPathname string, de *godirwalk.Dirent) error {
			// Following string operation is not most performant way
			// of doing this, but common enough to warrant a simple
			// example here:
			if !de.ModeType().IsDir() {
				return godirwalk.SkipThis
			}
			dir = append(dir, osPathname)

			return nil
		},
		Unsorted: true, // (optional) set true for faster yet non-deterministic enumeration (see godoc)
	})
	if err != nil {
		fmt.Print(err)
	}
	return dir
}

func WalkDir_FindImages(rootDir string) []string {
	// Walk directory, adding the files we want to add

	var files []string
	dirLevel := strings.Count(rootDir, "\\")
	err := godirwalk.Walk(rootDir, &godirwalk.Options{
		Callback: func(subPathname string, de *godirwalk.Dirent) error {
			// Following string operation is not most performant way
			// of doing this, but common enough to warrant a simple
			// example here:
			subDirVal := strings.Count(subPathname, "\\")
			if dirLevel == subDirVal-1 {
				if strings.Contains(subPathname, ".jpg") ||
					strings.Contains(subPathname, ".jpeg") ||
					strings.Contains(subPathname, ".png") ||
					strings.Contains(subPathname, ".webp") {
					files = append(files, subPathname)
				}
			}

			return nil
		},
		Unsorted: true, // (optional) set true for faster yet non-deterministic enumeration (see godoc)
	})
	if err != nil {
		fmt.Print(err)
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
