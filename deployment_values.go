package main

import (
	"context"

	"github.com/google/go-github/v47/github"
)

var org, err = getSecret("GITHUB_ORG")

type Configuration interface {
	GetContents(ctx context.Context, owner string, repo string, path string, opts *github.RepositoryContentGetOptions) (string, error)
}

type GithubConfiguration struct {
	client github.Client
}

func (r GithubConfiguration) GetContents(ctx context.Context, owner string, repo string, path string, opts *github.RepositoryContentGetOptions) (string, error) {
	fileContent, _, _, err := r.client.Repositories.GetContents(ctx, owner, repo, path, opts)
	if err != nil {
		return "", err
	}
	return fileContent.GetDownloadURL(), nil
}

func init() {
	if err != nil {
		panic("Error while trying to get the github org")
	}
}

func getDeploymentValuesFromRepo(repo string, configuration Configuration) (string, error) {
	ctx := context.Background()
	return configuration.GetContents(ctx, org, repo, "deployment.yaml", &github.RepositoryContentGetOptions{Ref: "OTT-XXX/Fix-Values"})
}
