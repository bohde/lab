package lab

type Git interface {
	LocalBranch() (string, error)
	RemoteProject() (RemoteProject, error)
}

type Gitlab interface {
	Project(RemoteProject) (Project, error)
	CreateMergeRequest(RemoteProject, *CreateMergeRequestOptions) (MergeRequest, error)
}
