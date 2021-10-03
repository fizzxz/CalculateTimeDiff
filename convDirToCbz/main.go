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
		subDirs := filepath.GetSubDirs(rootDir)

		for _, subDir := range subDirs {
			foundImagesOSFiles := filepath.WalkDir_FindImages(subDir)
			filesOS := filepath.GetFileOs(foundImagesOSFiles)
			if len(filesOS) > 0 {
				if compressionTypes.ZipArchiveDir_FastZip(subDir, filesOS) {
					filepath.RenameDirToCbz(subDir)
				}
			}
		}
	}

	// Golang archive/zip and walk
	// using a ryzen 3700x
	// can convert 1,150 Files, 58 Folders (1.04GB) in 30 seconds
	if zipTypeCompression == glZipCompression {

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
