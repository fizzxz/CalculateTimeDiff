package compressionTypes

import (
	"archive/zip"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"convDirToCbz/filepath"
)

func ZipArchiveDir(rootDir string) bool {
	zipDir := rootDir + ".zip"
	if cap(filepath.FindDir(rootDir)) > 0 {

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
