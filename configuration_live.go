package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/google/go-github/v47/github"
	"golang.org/x/oauth2"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/release"
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

type LiveConfiguration struct {
	client github.Client
	githubOrg string

}

func (r LiveConfiguration) GetContents(ctx context.Context, owner string, repo string, path string, opts *github.RepositoryContentGetOptions) (string, error) {
	fileContent, _, _, err := r.client.Repositories.GetContents(ctx, owner, repo, path, opts)
	if err != nil {
		return "", err
	}
	return fileContent.GetDownloadURL(), nil
}

func (r LiveConfiguration) GetDevRepos(devEnv string) ([]string, error) {
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

func (r LiveConfiguration) GetRegion() (string, error) {
	session, err := session.NewSession()
	if err != nil{
		return "", err
	}
	return *session.Config.Region, nil
}

func (r LiveConfiguration) LocateChart(name string, client *action.Install) (string, error) {
	return client.LocateChart(name, cli.New())
}

func (r LiveConfiguration) LoadChart(location string) (*chart.Chart, error) {
	return loader.Load(location)
}

func (r LiveConfiguration) GetOrg() string {
	return r.githubOrg 
}

func (r LiveConfiguration) GetDefaultImageRepo(repo string, ciConfig CIConfig) (string, error) {
	return getSecret("default-image-repo")
}

func (r LiveConfiguration) Install(install *action.Install, chrt *chart.Chart, vals map[string]interface{}) (*release.Release, error) {
	return install.Run(chrt, vals)
}
