package main

import (
	"context"

	"github.com/joho/godotenv"
	"github.com/sugaml/github-oauth/src"

	"github.com/sirupsen/logrus"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		logrus.Fatalf("error getting env, not coming through %v", err)
	}
	logrus.Info("successfully load env values...")

	// code := "c0839352ef6fc8d15e82"
	ctx := context.Background()
	token := "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"
	// token, err := utils.CodeToAuthToken(code)
	// if err != nil {
	// 	logrus.Error("token get issue :: ", err)
	// 	return
	// }
	logrus.Info("access token :: ", token)
	client := src.NewGitHub(token)

	// Get Login User

	user, err := client.GetUser(ctx)
	if err != nil {
		logrus.Error("user err :: ", err)
		return
	}
	logrus.Info("Username :: ", *user.Login)

	// List repositoriies

	// repositories, err := client.GetRepositories(ctx)
	// if err != nil {
	// 	logrus.Error("token get issue :: ", err)
	// 	return
	// }
	// for _, repository := range repositories {
	// 	logrus.Infof("Repository Name :: %s || ID :: %d ", *repository.Name, *repository.ID)
	// }

	// Get Repository by ID = 505312609 Repo=simple-go

	repository, err := client.GetRepositoryByID(ctx, 505312609)
	if err != nil {
		logrus.Error("repo error :: ", err)
		return
	}
	logrus.Info("Repo :: ", *repository.Name)

	// Get Repository Conetent

	content, err := client.GetGithubContent(ctx, &src.GitContent{
		Owner:    *user.Login,
		RepoName: *repository.Name,
		Branch:   "main",
		Path:     "",
		FileName: "Dockerfile",
	})
	if err != nil {
		logrus.Error("git content error :: ", err)
		return
	}
	logrus.Info(content)
}
