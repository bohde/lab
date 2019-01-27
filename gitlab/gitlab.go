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

func (g *Gitlab) CreateMergeRequest(remote lab.RemoteProject, opts *lab.MergeRequest) error {
	client := g.getClient(remote)

	removeSource := !opts.KeepSource

	options := gitlab.CreateMergeRequestOptions{
		Title:              &opts.Title,
		Description:        &opts.Description,
		SourceBranch:       &opts.SourceBranch,
		TargetBranch:       &opts.TargetBranch,
		RemoveSourceBranch: &removeSource,
	}

	mr, _, err := client.MergeRequests.CreateMergeRequest(remote.Path, &options)
	if err != nil {
		return err
	}

	opts.URL = mr.WebURL

	return err
}

func (g *Gitlab) CreateIssue(remote lab.RemoteProject, opts *lab.Issue) error {
	client := g.getClient(remote)

	options := gitlab.CreateIssueOptions{
		Title:       &opts.Title,
		Description: &opts.Description,
	}

	issue, _, err := client.Issues.CreateIssue(remote.Path, &options)
	if err != nil {
		return err
	}

	opts.URL = issue.WebURL

	return err
}
