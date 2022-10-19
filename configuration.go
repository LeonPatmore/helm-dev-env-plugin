package main

import (
	"context"

	"github.com/google/go-github/v47/github"
)

type Configuration interface {
	GetContents(ctx context.Context, owner string, repo string, path string, opts *github.RepositoryContentGetOptions) (string, error)
	GetOrg() string
}
