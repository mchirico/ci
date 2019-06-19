package userMsg

import (
	"strings"
	"testing"
)

func TestNotFound(t *testing.T) {
	s := NotFound()
	strings.Contains(s, "You must be in a git repo project directory")
}
