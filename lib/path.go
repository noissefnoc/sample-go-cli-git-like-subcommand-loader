package lib

import (
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type Path struct {

}

const (
	pathEnvSeparator    = ":"
	filePrefixSeparator = "-"
)

func SearchBins(binPrefix string, searchPathEnv string, extraSearchPaths []string) (map[string]string, error) {

	binPrefix = binPrefix + filePrefixSeparator
	pathString := os.Getenv(searchPathEnv)
	searchPaths := strings.Split(pathString, pathEnvSeparator)

	if len(searchPaths) > 0 {
		sort.Sort(sort.Reverse(sort.StringSlice(searchPaths)))
	}

	if len(extraSearchPaths) != 0 {
		searchPaths = append(searchPaths, extraSearchPaths...)
	}

	var subcommandPath = make(map[string]string)

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

				subcommandName := strings.TrimLeft(fileName, binPrefix)
				subcommandPath[subcommandName] = path

				return nil
			},
		)
	}

	return subcommandPath, nil
}
