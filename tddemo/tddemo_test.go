package tddemo

import (
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)
func processFile(r io.Reader) bool {
	buf := make([]byte, 10)
	n, _ := r.Read(buf)
	return n != 0
}

func TestDataExample(t *testing.T) {
	f, _ := os.Open("testdata/fixture.json")
	defer f.Close()

	assert.True(t, processFile(f))
}
