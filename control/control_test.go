package control

import "testing"

func CreateRepo() interface{} {
	r := RepoCreate("testrepo",
		"https://github/mchirico/date.git",
		"master",
		"github.com/mchirico/testrepo")
	return r
}
func TestBuildUnitSH(t *testing.T) {

	r := CreateRepo()
	c := CreateCI()
	c.BuildUnitSH(r)
}

func TestBuildSH(t *testing.T) {

	r := CreateRepo()
	c := CreateCI()
	c.BuildSH(r)
}

func TestRunCI(t *testing.T) {

	r := CreateRepo()
	c := CreateCI()
	c.RunCI(r)
}

func TestBuildUnitTaskYML(t *testing.T) {

	c := CreateCI()
	r := CreateRepo()
	c.BuildUnitTaskYML(r)
}

func TestBuildPipelineYML(t *testing.T) {

	c := CreateCI()
	r := CreateRepo()
	c.BuildPipeline(r)
}

func TestBuildTaskYML(t *testing.T) {

	c := CreateCI()
	r := CreateRepo()
	c.BuildTaskYML(r)
}

func TestBuildNotes(t *testing.T) {

	c := CreateCI()
	r := CreateRepo()
	c.BuildNotes(r)
}

func TestInformTaskYML(t *testing.T) {

	c := CreateCI()
	r := CreateRepo()
	c.BuildInformTaskYML(r)
}

func TestInformSH(t *testing.T) {

	c := CreateCI()
	r := CreateRepo()
	c.InformSH(r)
}
