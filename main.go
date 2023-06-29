package main

import (
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/sugaml/github-oauth/src/handler"
)

func main() {
	http.HandleFunc("/", handler.HandleMain)
	http.HandleFunc("/login", handler.HandleGitHubLogin)
	http.HandleFunc("/github_oauth_cb", handler.HandleGitHubCallback)
	logrus.Info("Started running on http://127.0.0.1:8000\n")
	logrus.Info(http.ListenAndServe(":8000", nil))
}
