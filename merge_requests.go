package lab

import (
	"fmt"
)

type MergeRequest struct {
	URL         string
	Title       string
	Description string
}

type CreateMergeRequestOptions struct {
	Title        string
	Description  string
	SourceBranch string
	TargetBranch string
	KeepSource   bool
}

type MergeRequestService struct {
	Git    Git
	Gitlab Gitlab
}

func (service *MergeRequestService) Create(opts *CreateMergeRequestOptions) error {
	if opts.SourceBranch == "" {
		localBranch, err := service.Git.LocalBranch()
		if err != nil {
			return err
		}
		opts.SourceBranch = localBranch
	}

	remote, err := service.Git.RemoteProject()
	if err != nil {
		return err
	}

	project, err := service.Gitlab.Project(remote)
	if err != nil {
		return err
	}

	if opts.TargetBranch == "" {
		opts.TargetBranch = project.DefaultBranch
	}

	mr, err := service.Gitlab.CreateMergeRequest(remote, opts)

	if err != nil {
		return err
	}

	fmt.Print(mr.URL)
	return nil
}
