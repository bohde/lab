package lab

import "strings"

type Body struct {
	Title       string
	Description string
}

func (body *Body) ParseContent(content string) {
	splitContent := strings.SplitAfterN(content, "\n", 2)

	if len(splitContent) >= 1 {
		body.Title = strings.Trim(splitContent[0], "\n")
	}

	if len(splitContent) > 1 {
		body.Description = strings.Trim(splitContent[1], "\n")
	}
}
