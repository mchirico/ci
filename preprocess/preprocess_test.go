package preprocess

import (
	"fmt"
	"github.com/mchirico/ci/pkg"
	"log"
	"os"
	"path/filepath"
	"testing"
)

func ConstructDir() func() {

	old, err := os.Getwd()
	if err != nil {
		log.Fatalf("can't get current dir: %s\n", err)
	}
	repo := "github.com/spock/buildUniverse"
	mockdir := filepath.Join("../test-fixtures", repo)
	err = pkg.Mkdir(mockdir)
	if err != nil {
		log.Fatalf("ConstructDir Failed: %s\n", err)
	}
	os.Chdir(mockdir)

	return func() {
		os.Chdir(old)
		c, _ := os.Getwd()
		fmt.Printf("current: %s\n", c)

		err := os.Chdir("../test-fixtures")
		if err != nil {
			log.Fatalf("can't cd")
		}

		pkg.Rmdir("github.com")
		os.Chdir(old)

	}
}

func Test_CheckForGithubRepro(t *testing.T) {

	cleanup := ConstructDir()
	defer cleanup()

	repo, err := CheckForGithubRepro()
	if err != nil {
		fmt.Printf("err: %s\n", err)
		t.FailNow()
	}
	if repo != "buildUniverse" {
		t.Fatalf("Expected: %s\nGot: %s\n", "buildUniverse", repo)
	}
	fmt.Printf("repo: %s\n", repo)
}

func Test_CheckForGithubUser(t *testing.T) {
	user, err := CheckForGithubUser()
	if err != nil {
		fmt.Printf("err: %s\n", err)
	}
	fmt.Printf("dir: %s\n", user)
}

func TestFixture(t *testing.T) {
	data := filepath.Join("test-fixtures", "stuff.data")
	print(data)
}
