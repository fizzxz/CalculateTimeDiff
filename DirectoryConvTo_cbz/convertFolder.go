package main

import (
	"archive/zip"
	"context"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/karrick/godirwalk"
	"github.com/saracen/fastzip"
	"github.com/urfave/cli"
)

const (
	fastZipCompression = "fz"
	glZipCompression   = "gz"
)

var (
	app = cli.NewApp()
	err error
)

func info() {
	app.Name = "Image dir conversion to cbz"
	app.Usage = "Provide a root dir that contains folders for the application to convert to cbz"
	app.Authors = []cli.Author{{Name: "Fizzxz", Email: "faisalk96@outlook.com"}}
	app.Version = "0.0.1"
}
func commands() {
	app.Commands = []cli.Command{
		{
			Name:    "convert",
			Aliases: []string{"c", "conv", "convert"},
			Usage:   "Uses provided values to convert a dir into cbz",
			Flags: []cli.Flag{
				&cli.StringFlag{Name: "dir", Value: "", Usage: "root dir to convert to cbz"},
				&cli.StringFlag{Name: "compression", Usage: "compression type to use," +
					"default fastZip(fz), other option is golang zip(gz)", Value: ""},
			},
			Action: func(c *cli.Context) error {
				findDirToArchive(c.String("dir"), c.String("compression"))
				return nil
			},
		},
	}
}

func main() {
	defer timeTrack(time.Now(), "Converting files")
	info()
	commands()

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}

func findDirToArchive(rootDir, zipTypeCompression string) {

	//Default to the faster zip conversion
	if zipTypeCompression == fastZipCompression ||
		zipTypeCompression == "" {
		//fastzip and godirwalk
		// using a ryzen 3700x
		// can convert 1,150 Files, 58 Folders (1.04GB) in 2.6 seconds
		subDirs := walkDir_FindSubDirs(rootDir)
		_, subDirs = subDirs[0], subDirs[1:]

		for _, subDir := range subDirs {
			foundImagesOSFiles := walkDir_FindImages(subDir)
			filesOS := getOsFile(foundImagesOSFiles)
			if len(filesOS) > 0 {
				if zipArchiveDir_FastZip(subDir, filesOS) {
					cbzDir := subDir + ".cbz"
					zipDir := subDir + ".zip"
					os.Rename(zipDir, cbzDir)
				}
			}
		}
	}

	// Golang archive/zip and walk
	// using a ryzen 3700x
	// can convert 1,150 Files, 58 Folders (1.04GB) in 30 seconds
	if zipTypeCompression == glZipCompression {

		subDirs := walkDir_FindSubDirs(rootDir)
		_, subDirs = subDirs[0], subDirs[1:]

		for _, subDir := range subDirs {

			dirConv := convDir(subDir)
			if zipArchiveDir(dirConv) {
				cbzDir := dirConv + ".cbz"
				os.Rename(dirConv+".zip", cbzDir)
			} else {
				//Remove the created empty zip file
				//when a failed conversion occurs
				os.Remove(dirConv + ".zip")
			}
		}
	}

}

func zipArchiveDir_FastZip(rootDir string, filesOS map[string]os.FileInfo) bool {
	// Create archive file
	zipDir := rootDir + ".zip"

	if _, err := os.Stat(zipDir); os.IsNotExist(err) {

		w, err := os.Create(zipDir)
		if err != nil {
			panic(err)
		}
		defer w.Close()

		// Create new Archiver
		a, err := fastzip.NewArchiver(w, rootDir, fastzip.WithArchiverConcurrency(1))
		if err != nil {
			panic(err)

		}

		defer a.Close()

		// Register a non-default level compressor if required
		a.RegisterCompressor(zip.Deflate, fastzip.FlateCompressor(5))

		// Archive
		if err = a.Archive(context.Background(), filesOS); err != nil {
			panic(err)

		}

		return true
	} else {
		return false
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

func getOsFile(files []string) map[string]os.FileInfo {
	filesOS := make(map[string]os.FileInfo)

	for _, file := range files {
		filesOS[file], err = os.Stat(file)
		if err != nil {
			fmt.Print(err)
		}
	}
	return (filesOS)
}

func zipArchiveDir(rootDir string) bool {
	zipDir := rootDir + ".zip"
	if cap(findDir(rootDir)) > 0 {

		// Get a Buffer to Write To
		outFile, err := os.Create(zipDir)
		if err != nil {
			fmt.Println(err)
			return false
		}
		defer outFile.Close()

		// Create a new zip archive.
		w := zip.NewWriter(outFile)

		// Add some files to the archive.
		addFilesErr := addFiles(w, rootDir, "")

		if err != nil {
			fmt.Println(err)
			return false
		}

		// Make sure to check the error on Close.
		err = w.Close()

		//If an error occurs adding files,
		// return false
		if addFilesErr != nil {
			return false
		}
		if err != nil {
			fmt.Println(err)
			return false
		}
		return true
	}
	return false
}

func addFiles(w *zip.Writer, basePath, baseInZip string) error {

	files, err := ioutil.ReadDir(basePath)
	if err != nil {
		fmt.Println(err)
		return err
	}

	for _, file := range files {

		if !file.IsDir() {

			if strings.Contains(file.Name(), ".jpg") ||
				strings.Contains(file.Name(), ".jpeg") ||
				strings.Contains(file.Name(), ".png") ||
				strings.Contains(file.Name(), ".webp") {

				dat, err := ioutil.ReadFile(basePath + "/" + file.Name())
				if err != nil {
					fmt.Println(err)
					return err
				}

				// Add some files to the archive.
				f, err := w.Create(baseInZip + file.Name())
				if err != nil {
					fmt.Println(err)
					return err
				}
				_, err = f.Write(dat)
				if err != nil {
					fmt.Println(err)
					return err
				}
			}

		} else if file.IsDir() {

			// Recurse
			newBase := basePath + file.Name() + "/"

			err := addFiles(w, newBase, baseInZip+file.Name()+"/")
			if err != nil {
				return err
			}
		}
	}
	return err
}

func findDir(rootDir string) []fs.FileInfo {
	files, err := ioutil.ReadDir(rootDir)
	if err != nil {
		fmt.Println(err)
	}

	return files
}

func convDir(rootDir string) string {
	return strings.ReplaceAll(rootDir, "\\", "/")
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
