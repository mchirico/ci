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
	dir         string
	buildUnitSH string
	unitTaskYml string
	pipeline    string
	buildTask   string
	buildSH     string
	runCI       string
	notes       string
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
