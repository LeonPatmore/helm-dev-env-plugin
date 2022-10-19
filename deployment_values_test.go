package main

import (
	"context"
	"testing"

	"github.com/google/go-github/v47/github"
	"github.com/stretchr/testify/assert"
)

type MockConfiguration struct {}

func (r MockConfiguration) GetContents(ctx context.Context, owner string, repo string, path string, opts *github.RepositoryContentGetOptions) (string, error) {
	return "someCoolValues",  nil
}

func (r MockConfiguration) GetOrg() string {
	return "myOrg"
}
 
func TestGetDeploymentValuesDownloadUrl(t *testing.T) {
	values, err := GetDeploymentValuesDownloadUrl("repo", MockConfiguration{})
	assert.Equal(t, err, nil)
	assert.Equal(t, values, "someCoolValues")
}
