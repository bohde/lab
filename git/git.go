package git

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"

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
	out, err := exec.Command("git", "symbolic-ref", "--short", "-q", "HEAD").Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSuffix(string(out), "\n"), err

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

	r := lab.ParseRemoteProject(remotes[0].Name, *remotes[0].URL)

	authToken, err := g.Get(fmt.Sprintf("lab.%s.token", r.Host))

	if err != nil || authToken == "" {
		err = lab.MissingToken{
			Host: r.Host,
		}
		return r, err
	}

	r.Token = authToken
	return r, nil

}

// RevList returns a list of revisions between the two commits
func (g *Git) RevList(base, merge string) ([]string, error) {
	commitRange := fmt.Sprintf("%s...%s", base, merge)
	out, err := exec.Command("git", "rev-list", commitRange).Output()
	if err != nil {
		return nil, err
	}

	revList := strings.Split(string(out), "\n")
	if len(revList) >= 1 {
		// trim newline, and original revision
		revList = revList[0 : len(revList)-1]
	}

	return revList, nil
}

// CommitMessage returns the commit message for a ref
func (g *Git) CommitMessage(commit string) (string, error) {
	out, err := exec.Command("git", "log", "--format=%B", "-n", "1", commit).Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}

// CommitMessages returns the first commit message between two refs
func (g *Git) CommitMessages(base, merge string) (string, error) {
	commitRange := fmt.Sprintf("%s...%s", base, merge)
	out, err := exec.Command("git", "log", "--pretty=format:%s", commitRange).Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func (g *Git) SetAccessToken(r lab.RemoteProject, token string) error {
	key := fmt.Sprintf("lab.%s.token", r.Host)
	return g.Set(key, token, true)
}

func (g *Git) Get(key string) (string, error) {
	out, err := exec.Command("git", "config", key).Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSuffix(string(out), "\n"), err
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
