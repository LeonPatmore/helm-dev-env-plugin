package main

import (
	"errors"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chartutil"
	kubefake "helm.sh/helm/v3/pkg/kube/fake"
	"helm.sh/helm/v3/pkg/registry"
	"helm.sh/helm/v3/pkg/storage"
	"helm.sh/helm/v3/pkg/storage/driver"
)

// See https://github.com/helm/helm/blob/d79ae9f927c00069ec39571501e82a9a8d0697ef/pkg/action/action_test.go#L37
func ActionConfigFixture(t *testing.T) *action.Configuration {
	t.Helper()

	registryClient, err := registry.NewClient()
	if err != nil {
		t.Fatal(err)
	}

	return &action.Configuration{
		Releases:       storage.Init(driver.NewMemory()),
		KubeClient:     &kubefake.FailingKubeClient{PrintingKubeClient: kubefake.PrintingKubeClient{Out: ioutil.Discard}},
		Capabilities:   chartutil.DefaultCapabilities,
		RegistryClient: registryClient,
	}
}

func TestLoadChartFromLocation(t *testing.T) {
	installAction := action.NewInstall(ActionConfigFixture(t))
	_, err := LoadChartFromLocation("chart", installAction, MockConfiguration{})
	assert.Equal(t, err, nil)
}

func TestLoadChartFromLocationWhenLocateChartError(t *testing.T) {
	installAction := action.NewInstall(ActionConfigFixture(t))
	_, err := LoadChartFromLocation("chart", installAction, MockConfiguration{locateChartErr: errors.New("Some big error")})
	assert.NotNil(t, err)
}
