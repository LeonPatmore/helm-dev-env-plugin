package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCiYaml(t *testing.T) {
	ciYaml, err := GetCiYaml("repo", MockConfiguration{getContentsRes: `version: '1.0'
namespace: mynamespace
chart:
    name: coolChart
    version: 1.0.0`})
	assert.Equal(t, err, nil)
	assert.Equal(t, "1.0", ciYaml.Version)
	assert.Equal(t, "mynamespace", ciYaml.Namespace)
	assert.Equal(t, "coolChart", ciYaml.Chart.Name)
	assert.Equal(t, "1.0.0", ciYaml.Chart.Version)
}

func TestGetCiYamlWithDefaultChartName(t *testing.T) {
	ciYaml, err := GetCiYaml("repo", MockConfiguration{getContentsRes: `version: '1.0'
namespace: mynamespace
chart:
    version: 1.0.0`})
	assert.Equal(t, err, nil)
	assert.Equal(t, "1.0", ciYaml.Version)
	assert.Equal(t, "mynamespace", ciYaml.Namespace)
	assert.Equal(t, "myChartName", ciYaml.Chart.Name)
	assert.Equal(t, "1.0.0", ciYaml.Chart.Version)
}

func TestGetCiYamlErr(t *testing.T) {
	_, err := GetCiYaml("repo", MockConfiguration{getContentsErr: errors.New("Some error")})
	assert.NotNil(t, err)
}
