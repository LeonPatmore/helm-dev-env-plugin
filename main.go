package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var namespace string
var devEnv string
var tags []string

func main() {
	org, err := getSecret("org")
	githubClient := GetGithubClient()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	config := LiveConfiguration{Client: githubClient, GithubOrg: org}
	var rootCmd = &cobra.Command{
		Use:   "helm dev",
		Short: "For creating a dev env",
		Run: func(cmd *cobra.Command, args []string) {
			RunDevInstall(devEnv, namespace, tags, config)
		},
	}
	rootCmd.Flags().StringVarP(&namespace, "devname", "d", "", "Namespace for the dev env")
	rootCmd.MarkFlagRequired("devname")

	rootCmd.Flags().StringVarP(&devEnv, "type", "t", "", "Type fo the dev env")
	rootCmd.MarkFlagRequired("type")

	rootCmd.Flags().StringArrayVar(&tags, "tag", []string{}, "Tags for the services you want to install on a branch")

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
