package main

import (
	"net/http"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/sugaml/github-oauth/src/handler"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		logrus.Fatalf("error getting env, not coming through %v", err)
	}
	logrus.Info("successfully load env values...")

	http.HandleFunc("/", handler.HandleMain)
	http.HandleFunc("/login", handler.HandleGitHubLogin)
	http.HandleFunc("/github_oauth_cb", handler.HandleGitHubCallback)
	logrus.Info("Started running on http://127.0.0.1:8000\n")
	logrus.Info(http.ListenAndServe(":8000", nil))
}
