package filepath

import (
	"os"
	"strings"
)

func RenameDirToCbz(rootDir string) {
	cbzDir := rootDir + ".cbz"
	zipDir := rootDir + ".zip"
	os.Rename(zipDir, cbzDir)
}

func ConvDir(rootDir string) string {
	return strings.ReplaceAll(rootDir, "\\", "/")
}
