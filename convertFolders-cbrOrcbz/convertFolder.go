package main

import (
	// "archive/zip"
	"context"
	"fmt"

	"os"

	// "path/filepath"
	// "fmt"
	"strings"

	"github.com/karrick/godirwalk"

	"github.com/saracen/fastzip"
)

// var filesOS = make(map[string]os.FileInfo)

// Zips "./input" into "./output.zip"
func main() {

	// archiveDir := "D:\\TestFolder\\archive.zip"
	rootDir := "D:\\TestFolder\\"
	findDirToArchive(rootDir, "cbz")
}

func findDirToArchive(rootDir string, archiveType string) {
	subDirs := walkDir_FindSubDirs(rootDir)
	for _, subDir := range subDirs {
		zipArchiveDir(subDir)
	}

}

func zipArchiveDir(rootDir string) {
	// Create archive file
	w, err := os.Create(rootDir + ".zip")
	if err != nil {
		panic(err)
	}
	defer w.Close()

	// Create new Archiver
	a, err := fastzip.NewArchiver(w, rootDir)
	if err != nil {
		panic(err)
	}
	defer a.Close()

	foundImagesOSFiles := walkDir_FindImages(rootDir)
	filesOS := getOsFile(foundImagesOSFiles)
	// Register a non-default level compressor if required
	// a.RegisterCompressor(zip.Deflate, fastzip.FlateCompressor(1))

	// Archive
	if err = a.Archive(context.Background(), filesOS); err != nil {
		panic(err)
	}
}

func walkDir_FindSubDirs(rootDir string) []string {
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

			// fmt.Printf(" %s\n", osPathname)
			return nil
		},
		Unsorted: true, // (optional) set true for faster yet non-deterministic enumeration (see godoc)
	})
	if err != nil {
		fmt.Print(err)
	}
	return dir
}

func walkDir_FindImages(rootDir string) []string {
	// Walk directory, adding the files we want to add

	var files []string
	err := godirwalk.Walk(rootDir, &godirwalk.Options{
		Callback: func(osPathname string, de *godirwalk.Dirent) error {
			// Following string operation is not most performant way
			// of doing this, but common enough to warrant a simple
			// example here:
			if strings.Contains(osPathname, ".jpg") ||
				strings.Contains(osPathname, ".jpeg") ||
				strings.Contains(osPathname, ".png") ||
				strings.Contains(osPathname, ".webp") {
				files = append(files, osPathname)
			}

			// return godirwalk.SkipThis
			// fmt.Printf(" %s\n", osPathname)
			return nil
		},
		Unsorted: true, // (optional) set true for faster yet non-deterministic enumeration (see godoc)
	})
	if err != nil {
		fmt.Print(err)
	}
	return files
	// filepath.Walk(sourceDir, func(pathname string, info os.FileInfo, err error) error {
	// 	files[pathname] = info
	// 	return nil
	// })
}

func getOsFile(files []string) map[string]os.FileInfo {
	filesOS := make(map[string]os.FileInfo)
	for _, file := range files {
		filesOS[file], _ = os.Stat(file)
		// if err != nil {
		// 	fmt.Print(err)
		// }
	}
	return filesOS
}
