package preprocess

import (
	"fmt"
	"github.com/mchirico/ci/control"
	"github.com/mchirico/ci/pkg"
	"github.com/mchirico/ci/userMsg"
)

var _repo string
var _user string

func checkForGithubRepo() (string, error) {

	pwd, err := pkg.GetPWD()
	if err != nil {
		return "", fmt.Errorf("error: pkg.GetPWD:%s\n", err)
	}
	repo, found := pkg.GithubRepo(pwd)
	if !found {
		fmt.Printf("%s", userMsg.NotFound())
		return "", fmt.Errorf("Not in project space.")
	}

	return repo, err

}

func CheckForGithubRepo() (string, error) {

	if _repo != "" {
		return _repo, nil
	}

	repo, err := checkForGithubRepo()
	if err != nil {
		return repo, err
	}

	_repo = repo
	return repo, err

}

func checkForGithubUser() (string, error) {

	pwd, err := pkg.GetPWD()
	if err != nil {
		return "", fmt.Errorf("error: pkg.GetPWD:%s\n", err)
	}
	user, found := pkg.GithubUser(pwd)
	if !found {
		fmt.Printf("Not found")
		return "", fmt.Errorf("Not in project space.")
	}

	return user, err

}

func CheckForGithubUser() (string, error) {

	if _user != "" {
		return _user, nil
	}

	user, err := checkForGithubUser()
	if err != nil {
		return user, err
	}

	_user = user
	return user, err

}

func BuildDefaultRepoStruct(branch string) (pkg.Repo, error) {

	repo, err := CheckForGithubRepo()
	if err != nil {
		return pkg.Repo{}, err
	}
	user, err := CheckForGithubUser()
	if err != nil {
		return pkg.Repo{}, err
	}

	http := fmt.Sprintf("https://github.com/%s/%s.git", user, repo)
	path := fmt.Sprintf("gopath/src/github.com/%s/%s", user, repo)
	r := pkg.Repo{repo, http,
		branch,
		path}

	return r, err

}

func BuildDefault(r pkg.Repo) {
	c := control.CreateCI()
	c.BuildUnitSH(r)
	c.BuildSH(r)
	c.RunCI(r)
	c.BuildUnitTaskYML(r)
	c.BuildPipeline(r)
	c.BuildTaskYML(r)
	c.BuildInformTaskYML(r)
	c.InformSH(r)
	c.BuildDockerTaskYML(r)
	c.DockerSH(r)
	c.BuildNotes(r)
	c.BuildDockerDirectory(r)

}
