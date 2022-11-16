package main

import (
	"fmt"
	"os"

	"github.com/google/go-github/v47/github"
	"github.com/spf13/cobra"
)

var namespace string
var devEnv string
var tags []string

func initCommands(githubClient *github.Client, org string) (*cobra.Command, error) {
	config := LiveConfiguration{Client: *githubClient, GithubOrg: org}
	var rootCmd = &cobra.Command{
		Use:   "helm dev",
		Short: "For creating a dev env",
		Run: func(cmd *cobra.Command, args []string) {
			err := RunDevInstall(devEnv, namespace, tags, config)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
		},
	}
	rootCmd.Flags().StringVarP(&namespace, "devname", "d", "", "Namespace for the dev env")
	err := rootCmd.MarkFlagRequired("devname")
	if err != nil {
		return nil, err
	}

	rootCmd.Flags().StringVarP(&devEnv, "type", "t", "", "Type fo the dev env")
	err = rootCmd.MarkFlagRequired("type")
	if err != nil {
		return nil, err
	}

	rootCmd.Flags().StringArrayVar(&tags, "tag", []string{}, "Tags for the services you want to install on a branch")
	return rootCmd, nil
}

func main() {
	org, err := getSecret("org")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	githubClient, err := GetGithubClient()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	rootCmd, err := initCommands(githubClient, org)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if err = rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
