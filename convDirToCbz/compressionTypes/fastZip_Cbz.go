package compressionTypes

import (
	"archive/zip"
	"context"
	"os"

	"github.com/saracen/fastzip"
)

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
