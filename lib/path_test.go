package lib

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func Test_searchBins(t *testing.T) {
	emptyPath, _ := filepath.Abs("../testdata/empty")
	bin1Path, _ := filepath.Abs("../testdata/path/bin1")
	bin2Path, _ := filepath.Abs("../testdata/path/bin2")
	extraPath, _ := filepath.Abs("../testdata/path/extra")

	type args struct {
		binPrefix        string
		searchPathEnv string
		extraSearchPaths []string
	}
	tests := []struct {
		name    string
		args    args
		pathEnv string
		want    []string
		wantErr bool
	}{
		{
			name: "empty",
			args: args{
				binPrefix: "sample",
				searchPathEnv: "X_TEST_PATH",
			},
			pathEnv: "",
		},
		{
			name: "only one env path (no match file)",
			args: args{
				binPrefix: "sample",
				searchPathEnv: "X_TEST_PATH",
			},
			pathEnv: emptyPath,
		},
		{
			name: "only one env path(with match file)",
			args: args{
				binPrefix: "sample",
				searchPathEnv: "X_TEST_PATH",
			},
			pathEnv: bin1Path,
			want: []string{filepath.Join(bin1Path, "sample-ok")},
		},
		{
			name: "multiple env path",
			args: args{
				binPrefix: "sample",
				searchPathEnv: "X_TEST_PATH",
			},
			pathEnv: bin1Path + ":" + bin2Path,
			want: []string{filepath.Join(bin1Path, "sample-ok"), filepath.Join(bin2Path, "sample-ok")},
		},
		{
			name: "only extraPath",
			args: args{
				binPrefix: "sample",
				searchPathEnv: "X_TEST_PATH",
				extraSearchPaths: []string{extraPath},
			},
			pathEnv: "",
			want: []string{filepath.Join(extraPath, "sample-ok")},
		},
		{
			name: "multiple path with extraPath",
			args: args{
				binPrefix: "sample",
				searchPathEnv: "X_TEST_PATH",
				extraSearchPaths: []string{extraPath},
			},
			pathEnv: bin1Path + ":" + bin2Path,
			want: []string{
				filepath.Join(bin1Path, "sample-ok"),
				filepath.Join(bin2Path, "sample-ok"),
				filepath.Join(extraPath, "sample-ok"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer os.Unsetenv(tt.args.searchPathEnv)
			os.Setenv(tt.args.searchPathEnv, tt.pathEnv)

			got, err := SearchBins(tt.args.binPrefix, tt.args.searchPathEnv, tt.args.extraSearchPaths)
			if (err != nil) != tt.wantErr {
				t.Errorf("searchBins() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchBins() got = %v, want %v", got, tt.want)
			}
		})
	}
}