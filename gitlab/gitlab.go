package gitlab

import (
	"fmt"

	"github.com/joshbohde/lab"
	gitlab "github.com/xanzy/go-gitlab"
)

type Gitlab struct{}

func New() *Gitlab {
	return &Gitlab{}
}

func (g *Gitlab) getClient(remote lab.RemoteProject) *gitlab.Client {
	c := gitlab.NewClient(nil, remote.Token)
	c.SetBaseURL(fmt.Sprintf("https://%s", remote.Host))
	return c
}

func (g *Gitlab) Project(remote lab.RemoteProject) (project lab.Project, err error) {
	client := g.getClient(remote)
	p, _, err := client.Projects.GetProject(remote.Path)
	if err != nil {
		return
	}

	project = lab.Project{
		DefaultBranch: p.DefaultBranch,
	}
	return
}

func (g *Gitlab) CreateMergeRequest(remote lab.RemoteProject, opts *lab.CreateMergeRequestOptions) (lab.MergeRequest, error) {
	client := g.getClient(remote)

	options := gitlab.CreateMergeRequestOptions{
		Title:        &opts.Title,
		Description:  &opts.Description,
		SourceBranch: &opts.SourceBranch,
		TargetBranch: &opts.TargetBranch,
	}

	mr, _, err := client.MergeRequests.CreateMergeRequest(remote.Path, &options)

	ret := lab.MergeRequest{
		URL:         mr.WebURL,
		Title:       mr.Title,
		Description: mr.Description,
	}

	return ret, err
}
