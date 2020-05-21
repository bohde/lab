package lab

import (
	"net/url"
	"strings"
)

type RemoteProject struct {
	Name  string
	Host  string
	Path  string
	Token string
}

type Project struct {
	DefaultBranch string
}

func ParseRemoteProject(name string, url url.URL) RemoteProject {
	path := strings.TrimSuffix(url.EscapedPath(), ".git")

	return RemoteProject{
		Name: name,
		Host: url.Hostname(),
		Path: path[1:],
	}

}
