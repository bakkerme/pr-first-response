package main

import (
	"context"
	"fmt"

	hfutils "github.com/bakkerme/hyperfocus-utils"
	"github.com/google/go-github/v40/github"
	"golang.org/x/oauth2"
)

func main() {
	ctx := context.Background()

	envRead := hfutils.EnvRead{}
	token, found := envRead.LookupEnv("GITHUB_ACCESS_TOKEN")

	if !found {
		panic("Could not load GITHUB_ACCESS_TOKEN from env")
	}

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	searchResult, _, err := client.Search.Issues(ctx, "user-review-requested:@me", nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(searchResult)

	// notifications, _, err := client.Activity.ListNotifications(ctx, nil)
	// if err != nil {
	// panic(err)
	// }

	// for _, notification := range notifications {
	// // fmt.Println(notification)
	// if notification.Unread != nil && !(*notification.Unread) {
	// continue
	// }

	// fmt.Println(notification.GetURL())
	// fmt.Println(notification.GetReason())
	// fmt.Println(notification.GetRepository().GetName())
	// fmt.Println(notification.GetSubject().GetTitle())
	// fmt.Println("                      ")
	// }

	// list all repositories for the authenticated user
	// repos, _, err := client.Repositories.List(ctx, "", nil)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(repos)
}
