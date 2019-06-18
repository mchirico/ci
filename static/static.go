package static

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func Notes() string {
	file := filepath.Join("../staticFile", "NOTES.md")
	dat, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Printf("Cannot read file ./NOTES.md")
	}
	return string(dat)

}
