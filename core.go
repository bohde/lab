package lab

type Git interface {
	LocalBranch() (string, error)
	RemoteProject() (RemoteProject, error)
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
