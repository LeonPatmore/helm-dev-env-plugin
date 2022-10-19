package main

import (
	"context"
	"fmt"

	"github.com/google/go-github/v47/github"
	"gopkg.in/yaml.v2"
)

type CIConfig struct {
	version   string
	namespace string
	chart     struct {
		name    string "olympus-service"
		version string
	}
}

func GetCiYaml(repo string, configuration Configuration) (*CIConfig, error) {
	ctx := context.Background()
	fileContent, err := configuration.GetContents(ctx, "nexmoinc", repo, "ci.yaml", &github.RepositoryContentGetOptions{Ref: "master"})
	if err != nil {
		return nil, err
	}
	fmt.Println(fileContent)
	config := &CIConfig{}
	err = yaml.Unmarshal([]byte(fileContent), config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
