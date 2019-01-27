package lab

import (
	"errors"
	"fmt"
	"os/exec"

	"github.com/github/hub/git"
	"github.com/github/hub/github"
	"github.com/joshbohde/lab"
)

// Git is a client for the local Git repo
type Git struct{}

func New() *Git {
	return &Git{}
}

// LocalBranch returns the name of the local branch
func (g *Git) LocalBranch() (string, error) {
	return git.Head()
}

// RemoteProject returns the remote Gitlab project of the local branch
func (g *Git) RemoteProject() (lab.RemoteProject, error) {
	remotes, err := github.Remotes()
	if err != nil {
		return lab.RemoteProject{}, err
	}

	if len(remotes) < 1 {
		return lab.RemoteProject{}, errors.New("no remotes found")
	}

	r := lab.ParseRemoteProject(*remotes[0].URL)

	authToken, err := g.Get(fmt.Sprintf("lab.%s.token", r.Host))

	if err != nil || authToken == "" {
		err = fmt.Errorf("create an access token and configure it by running `git config --global lab.%s.token <token>`", r.Host)
		return r, err
	}

	r.Token = authToken
	return r, nil

}

func (g *Git) Get(key string) (string, error) {
	out, err := exec.Command("git", "config", key).Output()
	if err != nil {
		return "", err
	}
	return string(out), err
}

func (g *Git) Set(key, val string, global bool) error {
	args := []string{"config"}
	if global {
		args = append(args, "--global")
	}
	args = append(args, key, val)

	_, err := exec.Command("git", args...).Output()
	return err
}
