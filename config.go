package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-secretsmanager-caching-go/secretcache"
)

var secretCache, _ = secretcache.New()
var secretPrefix = "helm-dev-plugin"

func getSecret(id string) (string, error) {
	env := os.Getenv(strings.ToUpper(id))
	if len(env) > 0 {
		return env, nil
	}
	res, err := secretCache.GetSecretString(fmt.Sprintf("%s/%s", secretPrefix, id))
	if err != nil{
		fmt.Printf("Failed to get AWS secret with id %s\n",id)
	}
	return res, err
}
