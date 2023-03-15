package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"os"
)

func main() {
	// Parse flags
	repos := flag.Bool("repos", false, "List repos")
	create := flag.Bool("create", false, "create resource, require --reponame flag")
	reponame := flag.String("reponame", "", "Repository Name")

	flag.Parse()

	switch {
	case *repos:
		listOrgs()
	case *create:
		creteRepository(*reponame)
	default:
		fmt.Fprintln(os.Stderr, "Invalid Option!")
		flag.Usage()
		os.Exit(1)
	}
}

func listOrgs() {

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: ""})
	authClient := oauth2.NewClient(ctx, ts)

	client := github.NewClient(authClient)
	// list organizations for my user for now...
	repos, _, err := client.Repositories.List(ctx, "isaiasrider", nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Repos")
	fmt.Printf("List of organizations: %v \n", github.Stringify(repos))

}

func creteRepository(name string) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: ""})
	authClient := oauth2.NewClient(ctx, ts)
	client := github.NewClient(authClient)

	repo := &github.Repository{
		Name:    github.String(name),
		Private: github.Bool(true),
	}
	client.Repositories.Create(ctx, "", repo)
	fmt.Printf("repo %v created", name)

}
