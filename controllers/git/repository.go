package git_controller

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/Minnek-Digital-Studio/cominnek/pkg/shell"
	"github.com/briandowns/spinner"
)

type IOriginInfo struct {
	Original string
	Owner    string
	Repo     string
}

func GetOrigin() *IOriginInfo {
	loader := spinner.New(spinner.CharSets[23], 100*time.Millisecond)
	loader.Start()
	loader.Prefix = "Getting repo info "

	origin := getOrigin()
	loader.Stop()

	return origin
}

func getOrigin() *IOriginInfo {
	out, errout, err := shell.Out("git ls-remote --get-url")
	sshRepo := "git@github.com:"
	httpsRepo := "https://github.com/"

	var originURL, owner, repo string

	if err != nil {
		fmt.Println(out)
		fmt.Println(errout)
		log.Fatal(errout)
	}

	if strings.HasPrefix(out, sshRepo) {
		originURL = strings.TrimPrefix(out, sshRepo)
	}

	if strings.HasPrefix(out, httpsRepo) {
		originURL = strings.TrimPrefix(out, httpsRepo)
	}

	if originURL != "" {
		owner = strings.Split(originURL, "/")[0]
		repoWithDot := strings.Split(originURL, "/")[1]
		repo = strings.Split(repoWithDot, ".")[0]
	}

	return &IOriginInfo{
		Original: out,
		Owner:    owner,
		Repo:     repo,
	}
}
	
func CheckGitRepo() bool {
    _, _, err := shell.Out("git rev-parse --is-inside-work-tree")
    return err == nil
}