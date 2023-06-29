package handler

import (
	"context"
	"net/http"

	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
	github_endpoint "golang.org/x/oauth2/github"
)

var (
	oauthConf = &oauth2.Config{
		ClientID:     "8560c194400c5c463823",
		ClientSecret: "678e55364c21e8b823de404b5f3005904f1f9797",
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
