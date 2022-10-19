package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO: Can I override mock configuration struct and make it return a valid yaml?

func TestGetCiYaml(t *testing.T) {
	_, err := GetCiYaml("repo", MockConfiguration{})
	assert.Equal(t, err, nil)
}
