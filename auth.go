package lab

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/skratchdot/open-golang/open"
)

type AuthService struct {
	Git Git
}

// AuthService will ensure the user logs in if the remote project requires it
func (service *AuthService) RemoteProject() (RemoteProject, error) {
	remote, err := service.Git.RemoteProject()

	if missingToken, ok := err.(MissingToken); ok {
		reader := bufio.NewReader(os.Stdin)

		fmt.Printf("No access token for https://%s/ found. Please create a new access token with api scopes.\n", missingToken.Host)
		fmt.Println("Press enter to open your browser to your access tokens page.")

		reader.ReadString('\n')

		err = open.Start(fmt.Sprintf("https://%s/profile/personal_access_tokens", missingToken.Host))
		if err != nil {
			fmt.Printf("Error trying open personal access token page.\n")
			return remote, missingToken
		}

		if err != nil {
			fmt.Printf("Error trying to read string.\n")
			return remote, missingToken
		}

		fmt.Print("Please paste the access token here: ")

		text, err := reader.ReadString('\n')

		err = service.Git.SetAccessToken(remote, strings.Trim(text, "\n"))
		return remote, err
	}

	return remote, err
}
