package main

import (
	"github.com/l58193/terrabutler/internal/cli"
	"github.com/l58193/terrabutler/internal/logger"

	"github.com/spf13/afero"
)

var version string

func main() {

	// Using Real FileSystem
	fs := afero.NewOsFs()

	err := cli.Run(version, fs)

	if err != nil {
		logger.Zap.Error(err.Error())
	}
}
