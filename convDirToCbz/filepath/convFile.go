package filepath

import (
	"os"
	"strings"
)

func renameDirToCbz(rootDir string) {
	cbzDir := rootDir + ".cbz"
	zipDir := rootDir + ".zip"
	os.Rename(zipDir, cbzDir)
}

func convDir(rootDir string) string {
	return strings.ReplaceAll(rootDir, "\\", "/")
}
