package github_interaction

import (
	"fmt"
	"github.com/google/go-github/github"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
)

func ListOrgs(token string, user string) {

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token})
	authClient := oauth2.NewClient(ctx, ts)

	client := github.NewClient(authClient)
	// list organizations for my user for now...
	repos, _, err := client.Repositories.List(ctx, user, nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Repos")
	fmt.Printf("List of organizations: %v \n", github.Stringify(repos))

}

func CreteRepository(name string, token string, user string) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token})
	authClient := oauth2.NewClient(ctx, ts)
	client := github.NewClient(authClient)

	repo := &github.Repository{
		Name:    github.String(name),
		Private: github.Bool(true),
	}
	client.Repositories.Create(ctx, "", repo)
	fmt.Printf("repo %v created", name)

}
