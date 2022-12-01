package common

import (
	"os"
	"testing"
)

func TestReadFile(t *testing.T) {
	testFilePath := "testfile"
	input := ReadFile(2021, 1, testFilePath)
	os.Remove(testFilePath)
	if len(input) != 9310 {
		t.Errorf("ReadFile expected len 9310, got %d", len(input))
	}
}
