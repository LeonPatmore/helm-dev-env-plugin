package main

import (
	"context"
	"fmt"

	"github.com/google/go-github/v47/github"
	"gopkg.in/yaml.v2"
)

type CIConfig struct {
	Version   string
	Namespace string
	Chart     struct {
		Name    string
		Version string
	}
}

func GetCiYaml(repo string, configuration Configuration) (*CIConfig, error) {
	ctx := context.Background()
	fileContent, err := configuration.GetContents(ctx, configuration.GetOrg(), repo, "ci.yaml", &github.RepositoryContentGetOptions{Ref: "master"})
	if err != nil {
		return nil, err
	}
	fmt.Println(fileContent)
	config := CIConfig{}
	err = yaml.Unmarshal([]byte(fileContent), &config)
	if err != nil {
		return nil, err
	}
	if config.Chart.Name == "" {
		config.Chart.Name = configuration.GetDefaultChatName()
	}
	return &config, nil
}
