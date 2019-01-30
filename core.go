package lab

import "fmt"

type MissingToken struct {
	Host string
}

func (m MissingToken) Error() string {
	return fmt.Sprintf("no access token found for %s. create one by running `lab auth` in this directory.", m.Host)
}

type Git interface {
	LocalBranch() (string, error)
	RemoteProject() (RemoteProject, error)
	SetAccessToken(RemoteProject, string) error
}

type Gitlab interface {
	Project(RemoteProject) (Project, error)
	CreateMergeRequest(RemoteProject, *MergeRequest) error
	CreateIssue(RemoteProject, *Issue) error
}

type MessageOpts struct {
	Edit      bool
	InputFile string
	EditFile  string
	Topic     string
	Comment   string
}

type Message interface {
	GetMessage(*string, MessageOpts) (func() error, error)
}
