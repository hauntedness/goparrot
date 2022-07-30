package main

import (
	"path/filepath"
	"testing"
)

func Test_parseFile(t *testing.T) {
	parseFile(filepath.Join("test.go"), 3, "Time;Open;High;Low;Close&&v[0];v[1];v[2];v[3];v[4]")
}
