package main

import (
	"context"
	"fmt"

	"gopkg.in/yaml.v2"
)

type CIConfig struct {
	Version   string
	Namespace string
	ImageRepo string `yaml:"imageRepo"`
	Chart     struct {
		Name    string
		Version string
	}
}

func GetCiYaml(repo string, configuration Configuration) (*CIConfig, error) {
	ctx := context.Background()
	fileContent, err := configuration.GetContents(ctx, configuration.GetOrg(), repo, "ci.yaml", nil)
	if err != nil {
		return nil, err
	}
	fmt.Println(*fileContent)
	config := CIConfig{}
	err = yaml.Unmarshal([]byte(*fileContent), &config)
	if err != nil {
		return nil, err
	}
	if config.Chart.Name == "" {
		defaultChartName, err := configuration.GetDefaultChatName()
		if err != nil {
			return nil, err
		}
		config.Chart.Name = defaultChartName
	}
	if config.ImageRepo == "" {
		imageRepo, err := configuration.GetDefaultImageRepo(repo, config)
		if err != nil {
			return nil, err
		}
		config.ImageRepo = imageRepo
	}
	return &config, nil
}
