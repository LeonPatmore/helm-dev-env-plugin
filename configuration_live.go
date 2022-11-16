package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/google/go-github/v47/github"
	"golang.org/x/oauth2"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/release"
)

func GetGithubClient() (*github.Client, error) {
	ctx := context.Background()
	token, err := getSecret("github_token")
	if err != nil {
		return nil, err
	}
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)
	return client, nil
}

type LiveConfiguration struct {
	Client    *github.Client
	GithubOrg string
}

func (r LiveConfiguration) GetContents(ctx context.Context, owner string, repo string, path string, ref *string) (*string, error) {
	var refAfterDefault *string = ref
	if ref == nil {
		defaultBranch, err := r.GetDefaultBranch(repo)
		if err != nil {
			return nil, err
		}
		refAfterDefault = &defaultBranch
	}
	fileContent, _, _, err := r.Client.Repositories.GetContents(ctx, owner, repo, path, &github.RepositoryContentGetOptions{Ref: *refAfterDefault})
	if err != nil {
		return nil, err
	}
	content, err := fileContent.GetContent()
	return &content, err
}

func (r LiveConfiguration) GetDevRepos(devEnv string) ([]string, error) {
	ctx := context.Background()
	res, _, err := r.Client.Search.Repositories(ctx, fmt.Sprintf("topic:%s org:%s", devEnv, r.GetOrg()), &github.SearchOptions{})
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
	if err != nil {
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
	return r.GithubOrg
}

func (r LiveConfiguration) GetDefaultImageRepo(repo string, ciConfig CIConfig) (string, error) {
	return getSecret("default-image-repo")
}

func (r LiveConfiguration) Install(install *action.Install, chrt *chart.Chart, vals map[string]interface{}) (*release.Release, error) {
	return install.Run(chrt, vals)
}

func (r LiveConfiguration) ActionConfiguration() *action.Configuration {
	settings := cli.New()
	actionConfig := new(action.Configuration)
	if err := actionConfig.Init(settings.RESTClientGetter(), namespace, os.Getenv("HELM_DRIVER"), log.Printf); err != nil {
		log.Printf("%+v", err)
		os.Exit(1)
	}
	return actionConfig
}

func (r LiveConfiguration) SearchRepos(devEnv string) (*github.RepositoriesSearchResult, *github.Response, error) {
	ctx := context.Background()
	return r.Client.Search.Repositories(ctx, fmt.Sprintf("topic:%s org:%s", devEnv, r.GithubOrg), &github.SearchOptions{})
}

func (r LiveConfiguration) GetDefaultChatName() (string, error) {
	return getSecret("default-chart-name")
}

func (r LiveConfiguration) GetDownloadURL(ctx context.Context, owner string, repo string, path string, ref *string) (*string, error) {
	var refAfterDefault *string = ref
	if ref == nil {
		defaultBranch, err := r.GetDefaultBranch(repo)
		if err != nil {
			return nil, err
		}
		refAfterDefault = &defaultBranch
	}
	fileContent, _, _, err := r.Client.Repositories.GetContents(ctx, owner, repo, path, &github.RepositoryContentGetOptions{Ref: *refAfterDefault})
	if err != nil {
		return nil, err
	}
	downloadUrl := fileContent.GetDownloadURL()
	return &downloadUrl, nil
}

func (r LiveConfiguration) GetDefaultBranch(repo string) (string, error) {
	ctx := context.Background()
	rep, _, err := r.Client.Repositories.Get(ctx, r.GetOrg(), repo)
	return *rep.DefaultBranch, err
}
