package main

import (
	"fmt"

	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart"
	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/cli/values"
	"helm.sh/helm/v3/pkg/getter"
)

func LoadChart(name string, client *action.Install, configuration Configuration) (*chart.Chart, error) {
	chartLocation, err := configuration.LocateChart(name, client)
	if err != nil {
		return nil, err
	}
	return configuration.LoadChart(chartLocation)
}

type LocalOptions struct {
	*values.Options
}

func (r LocalOptions) WithDefaultValues(imageTag string, ciConfig CIConfig, configuration Configuration) error {
	region, err := configuration.GetRegion()
	if err != nil {
		return err
	}
	extraValues := []string{
		fmt.Sprintf("awsRegion=%s", region),
		fmt.Sprintf("global.awsRegion=%s", region),
		fmt.Sprintf("image.tag=%s", imageTag),
		fmt.Sprintf("global.image.tag=%s", imageTag),
		fmt.Sprintf("imageRepo=%s", ciConfig.ImageRepo),
	}
	r.Values = extraValues
	return nil
}

func InstallService(chartName string, releaseName string, namespace string, imageTag string, opts *values.Options, ciConfig CIConfig, configuration Configuration) error {
	fmt.Printf("Installing chart [ %s ] with release name [ %s ] to namespace [ %s ], tag [ %s ]\n", chartName, releaseName, namespace, imageTag)
	client := action.NewInstall(configuration.ActionConfiguration())
	client.Namespace = namespace
	client.CreateNamespace = true
	client.ReleaseName = releaseName
	client.IsUpgrade = true

	localOptions := &LocalOptions{opts}
	err := localOptions.WithDefaultValues(imageTag, ciConfig, configuration)
	if err != nil {
		return err
	}

	chart, err := LoadChart(chartName, client, configuration)

	if err != nil {
		fmt.Printf("Failed to load chart [ %s ]\n", chartName)
		return err
	}

	p := getter.All(cli.New())
	optsAsMap, err := localOptions.MergeValues(p)
	if err != nil {
		return err
	}

	rel, err := configuration.Install(client, chart, optsAsMap)
	if err != nil {
		return err
	}
	fmt.Printf("Installed Chart from path: %s in namespace: %s\n", rel.Name, rel.Namespace)
	return nil
}
