package testifydemo

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDemoAssert(t *testing.T) {

	assert.Equal(t, 4, 2+2)

	assert.Contains(t, "Hello World", "World")
	assert.Contains(t, []string{"Hello", "World"}, "World")

	assert.DirExists(t, "testdata")

	assert.JSONEq(t,
		`{"hello": "world", "foo": "bar"}`,
		`{"foo": "bar", "hello": "world"}`,
	)

	assert.HTTPStatusCode(t, myHandler, "GET", "/home", nil, 200)
}

func myHandler(http.ResponseWriter, *http.Request) {
}
