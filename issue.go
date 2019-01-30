package lab

import (
	"errors"
	"fmt"
	"io"
)

type Issue struct {
	Body
	URL string
}

type CreateIssueOptions struct {
	Message string
	File    string
	Edit    bool
}

func (opts *CreateIssueOptions) Issue() Issue {
	ret := Issue{}
	ret.ParseContent(opts.Message)
	return ret
}

type IssueService struct {
	Git     Git
	Gitlab  Gitlab
	Message Message
	Writer  io.Writer
}

func (service *IssueService) Create(opts *CreateIssueOptions) error {
	remote, err := service.Git.RemoteProject()
	if err != nil {
		return err
	}

	delete, err := service.Message.GetMessage(&opts.Message, MessageOpts{
		Edit:      opts.Edit,
		InputFile: opts.File,
		EditFile:  "ISSUEMSG",
		Topic:     "issue",
		Comment:   fmt.Sprintf("Creating an issue in %s.\n\nWrite a message for this issue. The first line is the title and the rest is the description.", remote.Path),
	})

	if err != nil {
		return err
	}

	issue := opts.Issue()

	if issue.Title == "" {
		return errors.New("issue title is blank")
	}

	err = service.Gitlab.CreateIssue(remote, &issue)

	if err != nil {
		return err
	}

	fmt.Fprintf(service.Writer, "%s\n", issue.URL)

	err = delete()

	return err
}
