package message

import (
	"io/ioutil"

	"github.com/github/hub/github"
	"github.com/joshbohde/lab"
)

type Service struct{}

func New() *Service {
	return &Service{}
}

func (service *Service) GetMessage(msg *string, opts lab.MessageOpts) (delete func() error, err error) {
	delete = func() error {
		return nil
	}

	if opts.InputFile != "" {
		bytes, err := ioutil.ReadFile(opts.InputFile)
		if err != nil {
			return delete, err
		}
		content := string(bytes)
		*msg = content
	}

	var fileEditor *github.Editor

	if *msg == "" || opts.Edit {
		fileEditor, err = service.edit(msg, opts)
		if err != nil {
			return delete, err
		}
	}

	if fileEditor != nil {
		delete = fileEditor.DeleteFile
	}

	return delete, err
}

func (service *Service) edit(msg *string, opts lab.MessageOpts) (*github.Editor, error) {
	fileEditor, err := github.NewEditor(opts.EditFile, opts.Topic, *msg)

	if err != nil {
		return fileEditor, err
	}

	fileEditor.AddCommentedSection(opts.Comment)
	content, err := fileEditor.EditContent()

	if err != nil {
		return fileEditor, err
	}

	*msg = content

	return fileEditor, nil

}
