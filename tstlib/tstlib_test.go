package tstlib

import (
	"strings"
	"testing"
)

func TestConstructDir(t *testing.T) {
	defer ConstructDir()()
	pwd := PWD()
	if !strings.Contains(pwd, repo) {
		msg := `
             pwd: %s
             expected: %s
`
		t.Fatalf(msg, pwd, repo)
	}
}

