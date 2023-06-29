package handler

import (
	"context"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
	github_endpoint "golang.org/x/oauth2/github"
)

var (
	oauthConf = &oauth2.Config{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		Endpoint:     github_endpoint.Endpoint,
		Scopes: []string{
			"repo",
		},
	}
	oauthStateString = "thisshouldberandom"
)

const htmlIndex = `<html><body>
Logged in with <a href="/login">GitHub</a>
</body></html>
`

func HandleMain(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(htmlIndex))
}

// /login
func HandleGitHubLogin(w http.ResponseWriter, r *http.Request) {
	url := oauthConf.AuthCodeURL(oauthStateString, oauth2.AccessTypeOnline)
	logrus.Info("github login rul :: ", url)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func HandleGitHubCallback(w http.ResponseWriter, r *http.Request) {
	state := r.FormValue("state")
	if state != oauthStateString {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	code := r.FormValue("code")
	logrus.Info("service_code ::: ", code)
	ctx := context.Background()
	authToken, err := oauthConf.Exchange(ctx, code)
	if err != nil {
		return
	}
	logrus.Info("Token :: ", authToken.AccessToken)
	// http.Redirect(w, r, "https://console.test.01cloud.dev", http.StatusTemporaryRedirect)
}
