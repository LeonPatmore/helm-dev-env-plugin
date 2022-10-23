package main

import (
	"fmt"
	"log"
	"os"

	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/cli/values"
	"helm.sh/helm/v3/pkg/getter"
)

func LoadChartFromLocation(name string, client *action.Install, configuration Configuration) (*chart.Chart, error) {
	chartLocation, err := configuration.LocateChart(name, client)
	if err != nil {
		return nil, err
	}
	// if err := actionConfig.Init(cliSettings.RESTClientGetter(), namespace,
	// 	os.Getenv("HELM_DRIVER"), log.Printf); err != nil {
	// 	log.Printf("%+v", err)
	// 	os.Exit(1)
	// }
	return configuration.LoadChart(chartLocation)
}

// func installService(chartName string, releaseName string, namespace string, imageTag string, opts *values.Options, configuration Configuration) error {
// 	ecrRepo := fmt.Sprintf("nexmo-%s", releaseName)
// 	region, err := configuration.GetRegion()
// 	if err != nil {
// 		return err
// 	}
// 	extraValues := []string{
// 		fmt.Sprintf("awsRegion=%s", region),
// 		fmt.Sprintf("global.awsRegion=%s", region),
// 		fmt.Sprintf("image.tag=%s", imageTag),
// 		fmt.Sprintf("global.image.tag=%s", imageTag),
// 		fmt.Sprintf("ecrRepoName=%s", ecrRepo),
// 	}

// 	opts.Values = extraValues

// 	chartLocation := fmt.Sprintf("%s/%s", chartName, chartName)

// 	actionConfig := new(action.Configuration)
// 	settings := cli.New()
// 	client := action.NewInstall(actionConfig)

// }

func installServiceld(chartName string, releaseName string, namespace string, imageTag string, opts *values.Options) error {
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