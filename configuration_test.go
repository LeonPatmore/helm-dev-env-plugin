package main

import (
	"context"
	"testing"

	"github.com/google/go-github/v47/github"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart"
	"helm.sh/helm/v3/pkg/release"
)

type MockConfiguration struct {
	getContentsRes string
	getContentsErr error
	locateChartErr error
	getDefaultImageRepoErr error
	installErr error
	t *testing.T
}


func (r MockConfiguration) GetContents(ctx context.Context, owner string, repo string, path string, opts *github.RepositoryContentGetOptions) (string, error) {
	return r.getContentsRes, r.getContentsErr
}

func (r MockConfiguration) GetDownloadUrl(ctx context.Context, owner string, repo string, path string, opts *github.RepositoryContentGetOptions) (string, error) {
	return "some-url", nil
}

func (r MockConfiguration) GetOrg() string {
	return "myOrg"
}

func (r MockConfiguration) GetDefaultChatName() string {
	return "myChartName"
}

func (r MockConfiguration) GetDevRepos(devEnv string) ([]string, error) {
	return []string{"repo1", "repo2"}, nil
}

func (r MockConfiguration) GetRegion() (string, error) {
	return "eu-west-1", nil
}

func (r MockConfiguration) LocateChart(name string, client *action.Install) (string, error) {
	return "location", r.locateChartErr
}

func (r MockConfiguration) LoadChart(location string) (*chart.Chart, error) {
	return &chart.Chart{}, nil
}

func (r MockConfiguration) GetDefaultImageRepo(repo string, ciConfig CIConfig) (string, error) {
	return "cool-repo-" + repo + "-" + ciConfig.Namespace, r.getDefaultImageRepoErr
}

func (r MockConfiguration) ActionConfiguration() *action.Configuration {
	return ActionConfigFixture(r.t)
}

func (r MockConfiguration) Install(install *action.Install, chrt *chart.Chart, vals map[string]interface{}) (*release.Release, error) {
	return &release.Release{Namespace: "namespace"}, r.installErr
}
