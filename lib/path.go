package lib

import (
	"os"
	"path/filepath"
	"strings"
)

type Path struct {

}

const (
	pathEnvSeparator    = ":"
	filePrefixSeparator = "-"
)

func SearchBins(binPrefix string, searchPathEnv string, extraSearchPaths []string) ([]string, error) {

	binPrefix = binPrefix + filePrefixSeparator
	pathString := os.Getenv(searchPathEnv)
	searchPaths := strings.Split(pathString, pathEnvSeparator)

	if len(extraSearchPaths) != 0 {
		searchPaths = append(searchPaths, extraSearchPaths...)
	}

	var binPaths []string

	for i := range searchPaths {
		filepath.Walk(searchPaths[i],
			func(path string, info os.FileInfo, err error) error {
				_, fileName := filepath.Split(path)

				if !strings.HasPrefix(fileName, binPrefix) {
					return nil
				}

				// TODO: Check work well with Windows OS
				if info.Mode() & 0100 == 0 {
					return nil
				}

				binPaths = append(binPaths, path)

				return nil
			},
		)
	}

	return binPaths, nil
}
