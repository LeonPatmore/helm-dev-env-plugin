package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO: Can I override mock configuration struct and make it return a valid yaml?

func TestGetCiYaml(t *testing.T) {
	ciYaml, err := GetCiYaml("repo", MockConfiguration{getContentsRes: `version: '1.0'
namespace: mynamespace
chart:
    version: 1.0.0`})
	assert.Equal(t, err, nil)
	assert.Equal(t, "1.0", ciYaml.version)
	assert.Equal(t, "namespace", &ciYaml.namespace)
	assert.Equal(t, "olympus-service", ciYaml.chart.name)
	assert.Equal(t, "1.0.0", ciYaml.chart.name)
}

func TestGetCiYamlErr(t *testing.T) {
	_, err := GetCiYaml("repo", MockConfiguration{getContentsErr: errors.New("Some error")})
	assert.NotNil(t, err)
}
