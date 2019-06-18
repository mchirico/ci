package preprocess

import (
	"fmt"
	"github.com/mchirico/ci/pkg"
)

func CheckForGithubRepro() (string, error) {
	pwd, err := pkg.GetPWD()
	if err != nil {
		return "", fmt.Errorf("error: pkg.GetPWD:%s\n", err)
	}
	repo, found := pkg.GithubRepo(pwd)
	if !found {
		fmt.Printf("Not found")
		return "", fmt.Errorf("Not in project space.")
	}
	return repo, err

}

func CheckForGithubUser() (string, error) {
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

func BuildDefault(repo string) error {

	return fmt.Errorf("")

}
