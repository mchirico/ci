package preprocess

import (
	"fmt"
	"github.com/mchirico/ci/control"
	"github.com/mchirico/ci/pkg"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func ListFiles(pwd string) []string {
	var files []string

	err := filepath.Walk(pwd, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	return files
}

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

func TestCheck_LocalVars(t *testing.T) {
	defer ConstructDir()()

	_, err := CheckForGithubRepo()
	if err != nil {
		fmt.Printf("err: %s\n", err)
		t.FailNow()
	}

	if _repo == "" {
		t.FailNow()
	}

	_, err = CheckForGithubUser()
	if err != nil {
		fmt.Printf("err: %s\n", err)
		t.FailNow()
	}

	if _user == "" {
		t.FailNow()
	}

}

func Test_CheckForGithubRepro(t *testing.T) {

	defer ConstructDir()()

	repo, err := CheckForGithubRepo()
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

	defer ConstructDir()()

	user, err := CheckForGithubUser()
	if err != nil {
		t.Fatalf("err: %s\n", err)
	}

	if user != "spock" {
		t.Fatalf("expected user to be spock. God: %s\n", user)
	}

}

func TestBuildDefaultRepoStruct(t *testing.T) {

	defer ConstructDir()()

	r, err := BuildDefaultRepoStruct("develop")

	if err != nil {
		t.Fatalf("")
	}

	if r.Path != "gopath/src/github.com/spock/buildUniverse" {
		t.FailNow()
	}
	if r.RepoHttp != "https://github.com/spock/buildUniverse.git" {
		t.FailNow()
	}

	if r.Branch != "develop" {
		t.FailNow()
	}

	if r.Reposhort != "buildUniverse" {
		t.FailNow()
	}

}

func Test_BuildDefault(t *testing.T) {

	defer ConstructDir()()

	r, err := BuildDefaultRepoStruct("develop")

	if err != nil {
		t.Fatalf("")
	}
	BuildDefault(r)

	pwd, _ := os.Getwd()
	files := ListFiles(pwd)
	if len(files) < 13 {
		t.Fatalf("Not enough files created: %d", len(files))
	}
}

func FileIn(file, pwd string) bool {
	files := ListFiles(pwd)
	for _, value := range files {

		if strings.Contains(value, file) {
			return true
		}
	}
	return false
}

func TestFilesCreated(t *testing.T) {

	t.Parallel()
	defer ConstructDir()()
	pwd, _ := os.Getwd()

	c := control.CreateCI()
	r, err := BuildDefaultRepoStruct("develop")
	if err != nil {
		t.Fatalf("")
	}

	tests := []struct {
		name string
		c    func(r interface{})
		r    pkg.Repo
		file string
	}{
		{name: "test unit.sh", c: c.BuildUnitSH, r: r, file: "unit.sh"},
		{name: "test build.sh", c: c.BuildSH, r: r, file: "build.sh"},
		{name: "test run_ci.sh", c: c.RunCI, r: r, file: "run_ci.sh"},
		{name: "test unit-task.yml", c: c.BuildUnitTaskYML, r: r, file: "unit-task.yml"},
		{name: "test build-golang-pipeline.yml", c: c.BuildPipeline, r: r, file: "build-golang-pipeline.yml"},
		{name: "test build-task.yml", c: c.BuildTaskYML, r: r, file: "build-task.yml"},
		{name: "test inform-task.yml", c: c.BuildInformTaskYML, r: r, file: "inform-task.yml"},
		{name: "test inform.sh", c: c.InformSH, r: r, file: "inform.sh"},
		{name: "test docker-task.yml", c: c.BuildDockerTaskYML, r: r, file: "docker-task.yml"},
		{name: "test docker.sh", c: c.DockerSH, r: r, file: "docker.sh"},
		{name: "test NOTES.md", c: c.BuildNotes, r: r, file: "NOTES.md"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.c(r)
			if !FileIn(tc.file, pwd) {
				t.FailNow()
			}

		})
	}
}
