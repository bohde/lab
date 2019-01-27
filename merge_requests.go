package lab

import (
	"errors"
	"fmt"
	"strings"
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

func (opts *CreateMergeRequestOptions) ParseContent(content string) {
	splitContent := strings.SplitAfterN(content, "\n", 2)

	if len(splitContent) >= 1 {
		opts.Title = strings.Trim(splitContent[0], "\n")
	}

	if len(splitContent) > 1 {
		opts.Description = strings.Trim(splitContent[1], "\n")
	}
}

type MergeRequestService struct {
	Git    Git
	Gitlab Gitlab
	Editor Editor
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

	var fileEditor FileEditor
	if opts.Title == "" {
		fileEditor, err = service.edit(opts)
		if err != nil {
			return err
		}
	}

	if opts.Title == "" {
		return errors.New("merge request title is blank")
	}

	mr, err := service.Gitlab.CreateMergeRequest(remote, opts)

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
	msg := fmt.Sprintf("%s\n%s", opts.Title, opts.Description)
	fileEditor, err := service.Editor.New("MERGE_REQUESTMSG", "merge request", msg)

	if err != nil {
		return fileEditor, err
	}

	commentedSection := fmt.Sprintf("Requesting a merge from %s to %s.\n\nWrite a message for this merge request. The first line is the title and the rest is the description.", opts.SourceBranch, opts.TargetBranch)

	fileEditor.AddCommentedSection(commentedSection)
	content, err := fileEditor.EditContent()

	if err != nil {
		return fileEditor, err
	}

	opts.ParseContent(content)

	return fileEditor, nil

}
