package main

import (
	"fmt"
	"strings"

	"helm.sh/helm/v3/pkg/cli/values"
)

func GetReposForDevEnv(devEnv string, configuration Configuration) ([]string, error) {
	res, _, err := configuration.SearchRepos(devEnv)
	if err != nil {
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

func RunDevInstall(devEnv string, namespace string, tags []string, configuration Configuration) error {
	fmt.Printf("Installing dev to namespace %s\n", namespace)

	tagMap := make(map[string]string)
	for _, tagString := range tags {
		tagSplit := strings.Split(tagString, "=")
		fmt.Printf("Setting service %s to version %s", tagSplit[0], tagSplit[1])
		tagMap[tagSplit[0]] = tagSplit[1]
	}

	repos, err := GetReposForDevEnv(devEnv, configuration)
	if err != nil {
		return err
	}
	for _, repo := range repos {
		fmt.Printf("Installing using repo [ %s ]\n", repo)
		ciConfig, err := GetCiYaml(repo, configuration)
		if err != nil {
			return err
		}
		valueFileUrl, err := GetDeploymentValuesDownloadUrl(repo, configuration)
		if err != nil {
			return err
		}
		opts := values.Options{ValueFiles: []string{*valueFileUrl, "value files"}}

		var imageTag string
		if val, ok := tagMap[repo]; ok {
			imageTag = val
		} else {
			imageTag = "latest"
		}
		errr := InstallService(ciConfig.Chart.Name, repo, namespace, imageTag, &opts, *ciConfig, configuration)
		if err != nil {
			return errr
		}
	}
	return nil
}
