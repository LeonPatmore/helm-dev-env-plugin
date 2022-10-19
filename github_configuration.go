package main

import (
	"context"

	"github.com/google/go-github/v47/github"
)

type GithubConfiguration struct {
	client github.Client
	githubOrg string
}

func (r GithubConfiguration) GetContents(ctx context.Context, owner string, repo string, path string, opts *github.RepositoryContentGetOptions) (string, error) {
	fileContent, _, _, err := r.client.Repositories.GetContents(ctx, owner, repo, path, opts)
	if err != nil {
		return "", err
	}
	return fileContent.GetDownloadURL(), nil
}
