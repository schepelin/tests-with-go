package intro

import (
	"math/rand"
	"testing"
)

func TestExample (t *testing.T) {
	actual := 2 + 2

	if actual != 4 {
		t.Fail()
	}
}

func thisIsNotTest (t *testing.T) {
	t.Fail()
}

func AlsoNotTest (t *testing.T) {
	t.Fail()
}


func TestExampleTwo(t *testing.T) {
	actual := 2 + 2

	if actual != 4 {
		t.Errorf("2 + 2 = %d; Want 4", actual)
	}
}

func TestExampleSkip(t *testing.T) {
	t.Skip()
}

func TestExampleLogs(t *testing.T) {
	t.Skip()
	t.Log("This is a log message")
	got := rand.Intn(4)

	t.Logf("got %d", got)
	if got == 3 {
		t.Log("Going to fail test")
		t.Fail()
	}
}
