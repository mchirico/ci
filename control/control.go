package control

import (
	"github.com/mchirico/ci/pkg"
	"github.com/mchirico/ci/static"
	"github.com/mchirico/ci/templates"
	"log"
)

type Repo struct {
	Reposhort, RepoHttp, Branch, Path string
}

func RepoCreate(Reposhort, RepoHttp, Branch, Path string) interface{} {
	// Prepare some data to insert into the template.

	var r = Repo{
		Reposhort,
		RepoHttp,
		Branch,
		Path,
	}
	return r
}

type CI struct {
	dir           string
	buildUnitSH   string
	unitTaskYml   string
	pipeline      string
	buildTask     string
	buildSH       string
	runCI         string
	notes         string
	informSH      string
	informTaskYml string
	dockerSH      string
	dockerTaskYml string
}

func CreateCI() CI {
	c := CI{"ci", "unit.sh",
		"unit-task.yml",
		"build-golang-pipeline.yml",
		"build-task.yml",
		"build.sh",
		"run_ci.sh",
		"NOTES.md",
		"inform.sh",
		"inform-task.yml",
		"docker.sh",
		"docker-task.yml"}
	return c
}

func (c CI) BuildPipeline(r interface{}) {

	p := templates.Pipeline()
	b := pkg.TmpCreate(p, r)
	pkg.Mkdir(c.dir)

	err := pkg.WriteString(c.dir+"/"+c.pipeline, b.String(), 0755)
	if err != nil {
		log.Printf("error: %s\n", err)
	}

}

func (c CI) BuildUnitSH(r interface{}) {

	p := templates.UnitSH()
	b := pkg.TmpCreate(p, r)
	pkg.Mkdir(c.dir)

	err := pkg.WriteString(c.dir+"/"+c.buildUnitSH, b.String(), 0755)
	if err != nil {
		log.Printf("error: %s\n", err)
	}

}

func (c CI) InformSH(r interface{}) {

	p := templates.InformSH()
	b := pkg.TmpCreate(p, r)
	pkg.Mkdir(c.dir)

	err := pkg.WriteString(c.dir+"/"+c.informSH, b.String(), 0755)
	if err != nil {
		log.Printf("error: %s\n", err)
	}

}

func (c CI) DockerSH(r interface{}) {

	p := templates.DockerSH()
	b := pkg.TmpCreate(p, r)
	pkg.Mkdir(c.dir)

	err := pkg.WriteString(c.dir+"/"+c.dockerSH, b.String(), 0755)
	if err != nil {
		log.Printf("error: %s\n", err)
	}

}

func (c CI) BuildSH(r interface{}) {

	p := templates.BuildSH()
	b := pkg.TmpCreate(p, r)
	pkg.Mkdir(c.dir)

	err := pkg.WriteString(c.dir+"/"+c.buildSH, b.String(), 0755)
	if err != nil {
		log.Printf("error: %s\n", err)
	}

}

func (c CI) RunCI(r interface{}) {

	p := templates.RunCI()
	b := pkg.TmpCreate(p, r)
	pkg.Mkdir(c.dir)

	err := pkg.WriteString(c.dir+"/"+c.runCI, b.String(), 0755)
	if err != nil {
		log.Printf("error: %s\n", err)
	}

}

func (c CI) BuildUnitTaskYML(r interface{}) {

	p := templates.UnitTaskYML()
	b := pkg.TmpCreate(p, r)
	pkg.Mkdir(c.dir)

	err := pkg.WriteString(c.dir+"/"+c.unitTaskYml, b.String(), 0644)
	if err != nil {
		log.Printf("error: %s\n", err)
	}

}

func (c CI) BuildInformTaskYML(r interface{}) {

	p := templates.InformTask()
	b := pkg.TmpCreate(p, r)
	pkg.Mkdir(c.dir)

	err := pkg.WriteString(c.dir+"/"+c.informTaskYml, b.String(), 0644)
	if err != nil {
		log.Printf("error: %s\n", err)
	}

}

func (c CI) BuildDockerTaskYML(r interface{}) {

	p := templates.DockerTask()
	b := pkg.TmpCreate(p, r)
	pkg.Mkdir(c.dir)

	err := pkg.WriteString(c.dir+"/"+c.dockerTaskYml, b.String(), 0644)
	if err != nil {
		log.Printf("error: %s\n", err)
	}

}

func (c CI) BuildTaskYML(r interface{}) {

	p := templates.BuildTask()
	b := pkg.TmpCreate(p, r)
	pkg.Mkdir(c.dir)

	err := pkg.WriteString(c.dir+"/"+c.buildTask, b.String(), 0644)
	if err != nil {
		log.Printf("error: %s\n", err)
	}

}

func (c CI) BuildNotes(r interface{}) {

	p := static.Notes()
	b := pkg.TmpCreate(p, r)
	pkg.Mkdir(c.dir)

	err := pkg.WriteString(c.dir+"/"+c.notes, b.String(), 0644)
	if err != nil {
		log.Printf("error: %s\n", err)
	}

}
