package editor

import (
	"github.com/github/hub/github"
	"github.com/joshbohde/lab"
)

// An editor that stores files inside the .git directory
type Editor struct{}

func (e Editor) New(filename, topic, message string) (lab.FileEditor, error) {
	return github.NewEditor(filename, topic, message)
}
