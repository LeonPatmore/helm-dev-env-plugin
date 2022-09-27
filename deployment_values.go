package main

import (
	"context"

	"github.com/google/go-github/v47/github"
)

var org, err = getSecret("GITHUB_ORG")
if err != nil {
	
}

func getDeploymentValuesFromRepo(repo string, client github.Client) (string, error) {
	ctx := context.Background()
	fileContent, _, _, err := client.Repositories.GetContents(ctx, org, repo, "deployment.yaml", &github.RepositoryContentGetOptions{Ref: "OTT-XXX/Fix-Values"})
	if err != nil {
		return "", err
	}
	return fileContent.GetDownloadURL(), nil
}
