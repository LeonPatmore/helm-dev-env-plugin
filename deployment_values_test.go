package main

import (
	"context"
	"testing"

	"github.com/google/go-github/v47/github"
)

type MockConfiguration struct {}

func (r MockConfiguration) GetContents(ctx context.Context, owner string, repo string, path string, opts *github.RepositoryContentGetOptions) (string, error) {
	return "a",  nil
}

func TestGetDeploymentValuesFromRepo(t *testing.T) {
	getDeploymentValuesFromRepo("a", nil)
}	
