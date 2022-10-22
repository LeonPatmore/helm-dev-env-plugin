package main

import (
	"context"

	"github.com/google/go-github/v47/github"
)

type MockConfiguration struct {
	getContentsRes string "someCoolValues"
	getContentsErr error
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
