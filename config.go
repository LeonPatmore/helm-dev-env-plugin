package main

import (
	"os"

	"github.com/aws/aws-secretsmanager-caching-go/secretcache"
)

var secretCache, _ = secretcache.New()

func getSecret(id string) (string, error) {
	env := os.Getenv(id)
	if len(env) > 0 {
		return env, nil
	}
	res, err := secretCache.GetSecretString(id)
	return res, err
}
