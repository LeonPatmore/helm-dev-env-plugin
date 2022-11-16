package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSecret(t *testing.T) {
	t.Skip("Only used for e2e testing")
	_, err := getSecret("example")
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetSecret_usesEnvVar(t *testing.T) {
	os.Setenv("SOME_EXAMPLE", "val")
	val, err := getSecret("SOME_EXAMPLE")
	assert.Equal(t, "val", val)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetSecret_capitalisesForEnvVar(t *testing.T) {
	os.Setenv("SOME_EXAMPLE", "val")
	val, err := getSecret("some_example")
	assert.Equal(t, "val", val)
	if err != nil {
		t.Fatal(err)
	}
}
