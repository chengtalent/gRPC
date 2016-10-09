package main

import (
	"github.com/op/go-logging"
	"os"
	"path/filepath"
)

var caLogger = logging.MustGetLogger("ca")

func main() {
	rootPath := "/home/silei/Dianrong/ethereum"
	caDir := "CA1"

	path := filepath.Join(rootPath, caDir)

	caLogger.Info(path)

	if _, err := os.Stat(path); err != nil {
		caLogger.Info("Fresh start; creating databases, key pairs, and certificates.")

		if err := os.MkdirAll(path, 0755); err != nil {
			caLogger.Panic(err)
		}
	}
}
