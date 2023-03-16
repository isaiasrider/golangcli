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
	// Parse flags
	repos := flag.Bool("repos", false, "List repos")
	create := flag.Bool("create", false, "create resource, require --reponame flag")
	reponame := flag.String("reponame", "", "Repository Name")

	flag.Parse()

	switch {
	case *repos:
		github_interaction.ListRepos(token, user)
	case *create:
		github_interaction.CreateRepository(*reponame, token, user)
	default:
		fmt.Fprintln(os.Stderr, "Invalid Option!")
		flag.Usage()
		os.Exit(1)
	}
}
