package main

import (
	"context"

	"github.com/google/go-github/v47/github"
)


func (r GithubConfiguration) GetOrg() string {
	return r.githubOrg 
}

func GetDeploymentValuesDownloadUrl(repo string, configuration Configuration) (string, error) {
	ctx := context.Background()
	return configuration.GetContents(ctx, configuration.GetOrg(), repo, "deployment.yaml", &github.RepositoryContentGetOptions{Ref: "OTT-XXX/Fix-Values"})
}
