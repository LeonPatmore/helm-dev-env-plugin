package main

import (
	"context"
	"fmt"

	"github.com/google/go-github/v47/github"
	"golang.org/x/oauth2"
)

func getGithubClient() github.Client {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "token"},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)
	return *client
}

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

func (r GithubConfiguration) GetDevRepos(devEnv string) ([]string, error) {
	ctx := context.Background()
	res, _, err := r.client.Search.Repositories(ctx, fmt.Sprintf("topic:%s org:%s", devEnv, r.GetOrg()), &github.SearchOptions{})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Printf("Found %d repos for this dev env: ", *res.Total)

	var repos []string
	for _, x := range res.Repositories {
		repos = append(repos, *x.Name)
	}

	fmt.Println(repos)
	return repos, nil
}
