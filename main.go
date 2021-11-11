package main

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	hfutils "github.com/bakkerme/hyperfocus-utils"
	"github.com/google/go-github/v40/github"
	"github.com/hexops/valast"
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

	searchResult, _, err := client.Search.Issues(ctx, "user-review-requested:@me", &github.SearchOptions{ListOptions: github.ListOptions{PerPage: 100}})
	if err != nil {
		panic(err)
	}

	for _, result := range searchResult.Issues {
		fmt.Println(*result.Title)
		fmt.Println(*result.HTMLURL)
		fmt.Println(result.CreatedAt.Local())

		repoData := getPRDataFromURL(*result.HTMLURL)
		pr, _, err := client.PullRequests.Get(ctx, repoData.owner, repoData.repo, repoData.id)
		if err != nil {
			panic(err)
		}

		fmt.Println(valast.String(pr))
	}

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

type ownerAndRepo struct {
	owner string
	repo  string
	id    int
}

func getPRDataFromURL(url string) ownerAndRepo {
	// https://github.com/bakkerme/pr-first-response/pull/1
	splite := strings.Split(url, "/")
	fmt.Println(splite)

	owner := splite[3]
	repo := splite[4]
	idStr := splite[6]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		panic(err)
	}

	return ownerAndRepo{
		owner: owner,
		repo:  repo,
		id:    id,
	}
}
