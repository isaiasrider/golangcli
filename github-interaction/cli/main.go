package main

import (
	"flag"
	"fmt"
	"github_interaction"
	"os"
)

var token = os.Getenv("GH_ACC_TOKEN")
var user = os.Getenv("GH_USER")

func main() {
	// print usage
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "%s tool. Developed for provide autonomy to developers\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "Maintainer: github.com/isaiasrider \n")
		fmt.Fprintf(flag.CommandLine.Output(), "Copyright 2022\n")
		fmt.Fprintln(flag.CommandLine.Output(), "Usage information:")
		flag.PrintDefaults()
	}

	repos := flag.Bool("repos", false, "List repos")
	create := flag.Bool("create", false, "Create resource, require --reponame flag")
	reponame := flag.String("reponame", "", "Repository Name")
	org := flag.String("org", "", "Pass organization to create a repository")

	// Parse flags
	flag.Parse()

	switch {
	case *repos:
		github_interaction.ListRepos(token, user)
	case *create:
		github_interaction.CreateRepository(*reponame, token, *org)
	default:
		fmt.Fprintln(os.Stderr, "Invalid Option!")
		flag.Usage()
		os.Exit(1)
	}
}
