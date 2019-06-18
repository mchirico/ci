package pkg

import (
	"fmt"
	"github.com/mchirico/ci/templates"
	"testing"
)

func TestPkg(t *testing.T) {

	Sample()
}

func TestSample2(t *testing.T) {
	Sample2()
}

func TestMkdir(t *testing.T) {

	top := "/tmp/spock/"
	full := top + "junk/stuff/dir"
	file := full + "/file"

	Rmdir(top)
	err := WriteString(top+"sample", "file", 0600)
	if err == nil {
		t.FailNow()
	}

	Mkdir(full)
	err = WriteString(file, "This is sample data", 0600)
	if err != nil {
		t.FailNow()
	}

	Rmdir(top)
}

func TestPipeline(t *testing.T) {
	templates.Pipeline()
}

func RepoCreate() interface{} {
	// Prepare some data to insert into the template.
	type Repo struct {
		Reposhort, RepoHttp, Branch, Path string
	}
	var r = Repo{
		"testrepo",
		"https://github/mchirico/date.git",
		"master",
		"github.com/mchirico/testrepo",
	}
	return r
}

func TestTmpCreate(t *testing.T) {

	r := RepoCreate()
	p := templates.Pipeline()
	TmpCreate(p, r)
}

func TestBuildTask(t *testing.T) {

	r := RepoCreate()
	p := templates.BuildTask()
	TmpCreate(p, r)
}

func TestUnitTask(t *testing.T) {

	r := RepoCreate()
	p := templates.UnitTaskYML()
	TmpCreate(p, r)
}

func TestBuildSH(t *testing.T) {

	r := RepoCreate()
	p := templates.BuildSH()
	TmpCreate(p, r)
}

func TestBuildUnit(t *testing.T) {

	r := RepoCreate()
	p := templates.BuildUnit()
	TmpCreate(p, r)
}

func TestRunCI(t *testing.T) {

	r := RepoCreate()
	p := templates.RunCI()
	TmpCreate(p, r)
}

func TestUnitSH(t *testing.T) {

	r := RepoCreate()
	p := templates.UnitSH()
	TmpCreate(p, r)
}

func TestGetPWD(t *testing.T) {
	d, err := GetPWD()
	if err != nil {
		t.FailNow()
	}
	fmt.Printf("d: %v\n", d)
}

func TestGetRepo(t *testing.T) {
	s := "/Users/mchirico/ci/src/github.com/mchirico/ci/pkg"
	if r, found := GithubRepo(s); found {
		if r != "ci" {
			t.FailNow()
		}
	}

	s2 := "/Users/mchirico/ci/src/mchirico/ci/pkg"
	if r, found := GithubRepo(s2); !found {
		if r != "" {
			t.FailNow()
		}
	}
	s3 := "/Users/mchirico/ci/src/github.com"
	if r, found := GithubRepo(s3); !found {
		if r != "" {
			t.FailNow()
		}
	}

}

func TestGetUser(t *testing.T) {
	s := "/Users/mchirico/ci/src/github.com/spock/ci/pkg"
	if r, found := GithubUser(s); found {
		if r != "spock" {
			t.FailNow()
		}
	}

	s3 := "/Users/mchirico/ci/src/github.com"
	if r, found := GithubUser(s3); !found {
		if r != "" {
			t.FailNow()
		}
	}

}
