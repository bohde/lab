package lab

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type AuthService struct {
	Git     Git
	Browser Browser
	Reader  io.Reader
	Writer  io.Writer
}

// AuthService will ensure the user logs in if the remote project requires it
func (service *AuthService) RemoteProject() (RemoteProject, error) {
	remote, err := service.Git.RemoteProject()

	reader := bufio.NewReader(service.Reader)

	if missingToken, ok := err.(MissingToken); ok {

		fmt.Fprintf(service.Writer, "No access token for https://%s/ found. Please create a new access token with api scopes.\n", missingToken.Host)
		fmt.Fprintln(service.Writer, "Press enter to open your browser to your access tokens page.")

		reader.ReadString('\n')

		url := fmt.Sprintf("https://%s/profile/personal_access_tokens", missingToken.Host)
		err = service.Browser.Open(url)
		if err != nil {
			fmt.Fprintf(service.Writer, "Error trying open personal access token page.\n")
			return remote, missingToken
		}

		if err != nil {
			fmt.Fprintf(service.Writer, "Error trying to read string.\n")
			return remote, missingToken
		}

		fmt.Fprint(service.Writer, "Please paste the access token here: ")

		text, err := reader.ReadString('\n')

		err = service.Git.SetAccessToken(remote, strings.Trim(text, "\n"))
		return remote, err
	}

	return remote, err
}
