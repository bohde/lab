package lab

import (
	"net/url"
	"strings"
)

type RemoteProject struct {
	Host  string
	Path  string
	Token string
}

type Project struct {
	DefaultBranch string
}

func ParseRemoteProject(url url.URL) RemoteProject {
	path := strings.TrimSuffix(url.EscapedPath(), ".git")

	return RemoteProject{
		Host: url.Hostname(),
		Path: path[1:],
	}

}
