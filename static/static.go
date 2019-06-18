package static

import (
	"fmt"
	"io/ioutil"
)

func Notes() string {
	dat, err := ioutil.ReadFile("./NOTES.md")
	if err != nil {
		fmt.Printf("Cannot read file ./NOTES.md")
	}
	return string(dat)

}
