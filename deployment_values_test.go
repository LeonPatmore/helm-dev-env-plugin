package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDeploymentValuesDownloadUrl(t *testing.T) {
	values, err := GetDeploymentValuesDownloadUrl("repo", MockConfiguration{})
	assert.Equal(t, err, nil)
	assert.Equal(t, values, "some-url")
}

// TODO: Test errors are handled.
