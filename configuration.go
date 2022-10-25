package main

import (
	"context"

	"github.com/google/go-github/v47/github"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart"
	"helm.sh/helm/v3/pkg/release"
)

type Configuration interface {
	GetContents(ctx context.Context, owner string, repo string, path string, opts *github.RepositoryContentGetOptions) (string, error)
	GetDownloadUrl(ctx context.Context, owner string, repo string, path string, opts *github.RepositoryContentGetOptions) (string, error)
	GetOrg() string
	GetDefaultChatName() (string, error)
	GetDefaultImageRepo(repo string, ciConfig CIConfig) (string, error)
	GetDevRepos(devEnv string) ([]string, error)
	GetRegion() (string, error)
	LocateChart(name string, client *action.Install) (string, error)
	LoadChart(location string) (*chart.Chart, error)
	ActionConfiguration() *action.Configuration
	Install(install *action.Install, chrt *chart.Chart, vals map[string]interface{}) (*release.Release, error)
	SearchRepos(devEnv string) (*github.RepositoriesSearchResult, *github.Response, error)
}
