package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/lrm25/programWrapper"
	"github.com/stretchr/testify/require"
)

func Test_main(t *testing.T) {

	wd, err := os.Getwd()
	require.NoError(t, err)

	program := programWrapper.NewProgram(filepath.Join(wd, "copypage.exe"))
	output, err := program.Run()
	require.NoError(t, err)
	require.Contains(t, output, "HTML page location must be defined")
}
