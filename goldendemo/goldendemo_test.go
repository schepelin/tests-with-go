package goldendemo

import (
	"flag"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var update = flag.Bool("update", false, "update .golden files")

func TestReportData(t *testing.T) {
	goldenfile := "testdata/report.golden"
	actual := generateReport()
	expected, _ := os.ReadFile(goldenfile)
	if *update {
		if err := os.WriteFile(goldenfile, actual, 0666); err != nil {
			t.Error(err)
		}
		return
	}
	assert.Equal(t, string(expected), string(actual))
}

func generateReport() []byte {
	return []byte{'A', ',', 'B', ',', 'C'}
}
