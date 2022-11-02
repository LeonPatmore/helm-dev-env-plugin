package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-secretsmanager-caching-go/secretcache"
)

var secretCache, _ = secretcache.New()
var secretPrefix = "helm-dev-plugin"

func getSecret(id string) (string, error) {
	env := os.Getenv(id)
	if len(env) > 0 {
		return env, nil
	}
	res, err := secretCache.GetSecretString(fmt.Sprintf("%s/%s", secretPrefix, id))
	return res, err
}
