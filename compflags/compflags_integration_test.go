//go:build integration

package compflags

import (
	"testing"
	"time"
)

func TestSlowExample(t *testing.T) {
	time.Sleep(3 * time.Second)
	if 2+2 != 4 {
		t.Fail()
	}
}
