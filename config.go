package main

import (
	"github.com/aws/aws-secretsmanager-caching-go/secretcache"
)

var secretCache, _ = secretcache.New()

func getSecret(id string) (string, error) {
	res, err := secretCache.GetSecretString(id)
	return res, err
}
