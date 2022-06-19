package main

import (
	"convDirToCbz/compressionTypes"
	"convDirToCbz/filepath"
	"fmt"
	"log"
	"os"

	"time"

	"github.com/urfave/cli"
)

const (
	fastZipCompression = "fz"
	glZipCompression   = "gz"
)

var (
	app = cli.NewApp()
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
				&cli.StringFlag{Name: "dir", Value: "", Usage: "root dir to conmpress to zip"},
				&cli.StringFlag{Name: "cbz", Value: "", Usage: "root dir to convert to cbz"},
				&cli.StringFlag{Name: "compression", Usage: "compression type to use," +
					"default fastZip(fz), other option is golang zip(gz)", Value: ""},
			},
			Action: func(c *cli.Context) error {
				findDirToArchive(c.String("dir"), c.String("compression"), c.Bool("cbz"))
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

func findDirToArchive(rootDir, zipTypeCompression string, comicBook bool) {

	//Default to the faster zip conversion
	if zipTypeCompression == fastZipCompression ||
		zipTypeCompression == "" {

		subDirs := filepath.GetSubDirs(rootDir)
		for _, subDir := range subDirs {
			foundFiles := findFiles(subDir)
			fmt.Println(foundFiles)
			if len(foundFiles) != 0 {
				if compressionTypes.ZipArchiveDir_FastZip(subDir, foundFiles) {
					if comicBook {
						filepath.RenameDirToCbz(subDir)
					}
				}
			}

		}
	} else if zipTypeCompression == glZipCompression {

		subDirs := filepath.GetSubDirs(rootDir)
		for _, subDir := range subDirs {

			dirConv := filepath.ConvDir(subDir)
			if compressionTypes.ZipArchiveDir(dirConv) {
				filepath.RenameDirToCbz(dirConv)
			} else {
				//Remove the created empty zip file
				//when a failed conversion occurs
				os.Remove(dirConv + ".zip")
			}
		}
	}

}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func findFiles(rootDir string) map[string]os.FileInfo {
	// for _, subDir := range rootDir {

	foundFiles := filepath.WalkDir_FindFiles(rootDir)
	filesOS := filepath.GetFileOs(foundFiles)
	if len(filesOS) == 0 {
		subDirs := filepath.GetSubDirs(rootDir)
		if len(subDirs) != 0 {
			for _, subDir := range subDirs {
				findFiles(subDir)
			}
		}
	}

	return filesOS
}
