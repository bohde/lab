package lab

type Git interface {
	LocalBranch() (string, error)
	RemoteProject() (RemoteProject, error)
}

type Gitlab interface {
	Project(RemoteProject) (Project, error)
	CreateMergeRequest(RemoteProject, *MergeRequest) error
}

type Editor interface {
	New(filename, topic, message string) (FileEditor, error)
}

type FileEditor interface {
	AddCommentedSection(text string)
	DeleteFile() error
	EditContent() (content string, err error)
}

type FileReader interface {
	Read(string) (content string, err error)
}
