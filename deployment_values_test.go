package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDeploymentValuesDownloadURL(t *testing.T) {
	values, err := GetDeploymentValuesDownloadURL("repo", MockConfiguration{})
	assert.Equal(t, err, nil)
	assert.Equal(t, *values, "some-url")
}

// TODO: Test errors are handled.
