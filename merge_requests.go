package lab

import (
	"errors"
	"fmt"
	"strings"
)

type MergeRequest struct {
	URL          string
	Title        string
	Description  string
	SourceBranch string
	TargetBranch string
	KeepSource   bool
}

func (mr *MergeRequest) ParseContent(content string) {
	splitContent := strings.SplitAfterN(content, "\n", 2)

	if len(splitContent) >= 1 {
		mr.Title = strings.Trim(splitContent[0], "\n")
	}

	if len(splitContent) > 1 {
		mr.Description = strings.Trim(splitContent[1], "\n")
	}
}

type CreateMergeRequestOptions struct {
	Message      string
	File         string
	Edit         bool
	SourceBranch string
	TargetBranch string
	KeepSource   bool
}

func (opts *CreateMergeRequestOptions) MergeRequest() MergeRequest {
	ret := MergeRequest{
		SourceBranch: opts.SourceBranch,
		TargetBranch: opts.TargetBranch,
		KeepSource:   opts.KeepSource,
	}
	ret.ParseContent(opts.Message)
	return ret

}

type MergeRequestService struct {
	Git    Git
	Gitlab Gitlab
	Editor Editor
	Reader FileReader
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

	if opts.File != "" {
		content, err := service.Reader.Read(opts.File)
		if err != nil {
			return err
		}
		opts.Message = content
	}

	var fileEditor FileEditor

	if opts.Message == "" || opts.Edit {
		fileEditor, err = service.edit(opts)
		if err != nil {
			return err
		}
	}

	mr := opts.MergeRequest()

	if mr.Title == "" {
		return errors.New("merge request title is blank")
	}

	err = service.Gitlab.CreateMergeRequest(remote, &mr)

	if err != nil {
		return err
	}

	fmt.Printf("%s\n", mr.URL)

	if fileEditor != nil {
		err = fileEditor.DeleteFile()
	}

	return err
}

func (service *MergeRequestService) edit(opts *CreateMergeRequestOptions) (FileEditor, error) {
	fileEditor, err := service.Editor.New("MERGE_REQUESTMSG", "merge request", opts.Message)

	if err != nil {
		return fileEditor, err
	}

	commentedSection := fmt.Sprintf("Requesting a merge from %s to %s.\n\nWrite a message for this merge request. The first line is the title and the rest is the description.", opts.SourceBranch, opts.TargetBranch)

	fileEditor.AddCommentedSection(commentedSection)
	content, err := fileEditor.EditContent()

	if err != nil {
		return fileEditor, err
	}

	opts.Message = content

	return fileEditor, nil

}
