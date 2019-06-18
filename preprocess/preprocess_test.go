package preprocess

import (
	"fmt"
	"testing"
)

func Test_CheckForGithubRepro(t *testing.T) {
	dir, err := CheckForGithubRepro()
	if err != nil {
		fmt.Printf("err: %s\n", err)
	}
	fmt.Printf("dir: %s\n", dir)
}

func Test_CheckForGithubUser(t *testing.T) {
	user, err := CheckForGithubUser()
	if err != nil {
		fmt.Printf("err: %s\n", err)
	}
	fmt.Printf("dir: %s\n", user)
}
