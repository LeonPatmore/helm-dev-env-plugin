package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/google/go-github/v47/github"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
	"gopkg.in/yaml.v3"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/cli/values"
	"helm.sh/helm/v3/pkg/getter"
)

var namespace string
var tags []string

type CIConfig struct {
	version   string
	namespace string
	chart     struct {
		name    string "olympus-service"
		version string
	}
}

type DevConfig struct {
	namespace string
}

func getDeploymentValueFileUrl(repo string, client github.Client) (string, error) {
	ctx := context.Background()
	fileContent, _, _, err := client.Repositories.GetContents(ctx, "nexmoinc", repo, "deployment.yaml", &github.RepositoryContentGetOptions{Ref: "OTT-XXX/Fix-Values"})
	if err != nil {
		return "", err
	}
	return fileContent.GetDownloadURL(), nil
}

func getCiYaml(repo string, client github.Client) (*CIConfig, error) {
	ctx := context.Background()
	fileContent, _, _, err := client.Repositories.GetContents(ctx, "nexmoinc", repo, "ci.yaml", &github.RepositoryContentGetOptions{Ref: "master"})
	if err != nil {
		return nil, err
	}
	yamlStr, err := fileContent.GetContent()
	if err != nil {
		return nil, err
	}
	fmt.Println(yamlStr)
	config := &CIConfig{}
	err = yaml.Unmarshal([]byte(yamlStr), config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

func getGithubClient() github.Client {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "token"},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)
	return *client
}

func getReposForDevEnv(devEnv string, client github.Client) []string {
	ctx := context.Background()
	res, _, err := client.Search.Repositories(ctx, fmt.Sprintf("topic:%s org:nexmoinc", devEnv), &github.SearchOptions{})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Found %d repos for this dev env: ", *res.Total)

	var repos []string
	for _, x := range res.Repositories {
		repos = append(repos, *x.Name)
	}

	fmt.Println(repos)
	return repos
}

func installService(chartName string, releaseName string, namespace string, imageTag string, opts *values.Options) error {
	actionConfig := new(action.Configuration)
	settings := cli.New()
	client := action.NewInstall(actionConfig)

	ecrRepo := fmt.Sprintf("nexmo-%s", releaseName)
	extraValues := []string{"awsRegion=eu-west-1",
		"global.awsRegion=eu-west-1",
		fmt.Sprintf("image.tag=%s", imageTag),
		fmt.Sprintf("global.image.tag=%s", imageTag),
		fmt.Sprintf("ecrRepoName=%s", ecrRepo)}
	opts.Values = extraValues

	chartLocation, err := client.LocateChart(fmt.Sprintf("%s/%s", chartName, chartName), settings)
	if err != nil {
		return err
	}

	if err := actionConfig.Init(settings.RESTClientGetter(), namespace,
		os.Getenv("HELM_DRIVER"), log.Printf); err != nil {
		log.Printf("%+v", err)
		os.Exit(1)
	}
	chart, err := loader.Load(chartLocation)
	if err != nil {
		return err
	}
	client.Namespace = namespace
	client.CreateNamespace = true
	client.ReleaseName = releaseName
	client.IsUpgrade = true

	// install the chart here
	p := getter.All(settings)
	optsAsMap, err := opts.MergeValues(p)
	if err != nil {
		return err
	}
	rel, err := client.Run(chart, optsAsMap)
	if err != nil {
		return err
	}

	log.Printf("Installed Chart from path: %s in namespace: %s\n", rel.Name, rel.Namespace)
	// this will confirm the values set during installation
	log.Println(rel.Config)
	return nil
}

func runDevInstall(namespace string) {
	fmt.Println("Hello, World!")

	fmt.Printf("namespace is %s\n", namespace)

	tagMap := make(map[string]string)
	for _, tagString := range tags {
		tagSplit := strings.Split(tagString, "=")
		fmt.Printf("Setting service %s to version %s", tagSplit[0], tagSplit[1])
		tagMap[tagSplit[0]] = tagSplit[1]
	}

	githubClient := getGithubClient()
	repos := getReposForDevEnv("messages", githubClient)
	for _, repo := range repos {
		ciConfig, err := getCiYaml(repo, githubClient)
		if err != nil {
			os.Exit(1)
		}
		var chartName string
		if len(ciConfig.chart.name) > 0 {
			chartName = ciConfig.chart.name
		} else {
			chartName = "olympus-service"
		}
		valueFileUrl, err := getDeploymentValueFileUrl(repo, githubClient)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		opts := values.Options{ValueFiles: []string{valueFileUrl,
			"value files"}}

		var imageTag string
		if val, ok := tagMap[repo]; ok {
			imageTag = val
		} else {
			imageTag = "latest"
		}
		errr := installService(chartName, repo, namespace, imageTag, &opts)
		if errr != nil {
			fmt.Println(errr)
			os.Exit(1)
		}
	}
}

func main() {
	var rootCmd = &cobra.Command{
		Use:   "helm olympus-dev",
		Short: "For creating a dev env",
		Run: func(cmd *cobra.Command, args []string) {
			runDevInstall(namespace)
		},
	}
	rootCmd.Flags().StringVarP(&namespace, "devname", "d", "", "Namespace for the dev env")
	rootCmd.MarkFlagRequired("devname")

	rootCmd.Flags().StringArrayVarP(&tags, "tag", "t", []string{}, "Tags for the services you want to install on a branch")

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
