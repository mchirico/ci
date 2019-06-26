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

func find(s string, list []string) bool {
	for _, v := range list {
		if strings.Contains(v, s) {
			return true
		}
	}
	return false
}

func TestListFiles(t *testing.T) {
	defer ConstructDir()()
	pwd := PWD()

	WriteString("bozo", "test", 0600)
	list := ListFiles(pwd)

	if !find("bozo", list) {
		t.Fatalf("not found")
	}

}
