package main

import (
	"testing"
)

func TestGetSecret(t *testing.T) {
	t.Skip("Only used for e2e testing")
    _, err := getSecret("example")
    if err != nil {
        t.Fatal(err)
    }
}
