package main

import (
	"context"

	"github.com/google/go-github/v47/github"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart"
)

type Configuration interface {
	GetContents(ctx context.Context, owner string, repo string, path string, opts *github.RepositoryContentGetOptions) (string, error)
	GetDownloadUrl(ctx context.Context, owner string, repo string, path string, opts *github.RepositoryContentGetOptions) (string, error)
	GetOrg() string
	GetDefaultChatName() string
	GetDevRepos(devEnv string) ([]string, error)
	GetRegion() (string, error)
	LocateChart(name string, client *action.Install) (string, error)
	LoadChart(location string) (*chart.Chart, error)
}
