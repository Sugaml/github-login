package utils

import (
	"context"
	"encoding/json"

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
	//oauthStateString = "thisshouldberandom"
)

func TokenToJSON(token *oauth2.Token) (string, error) {
	d, err := json.Marshal(token)
	if err != nil {
		return "", err
	}
	return string(d), nil

}

func TokenFromJSON(jsonStr string) (*oauth2.Token, error) {
	var token oauth2.Token
	if err := json.Unmarshal([]byte(jsonStr), &token); err != nil {
		return nil, err
	}
	return &token, nil
}

func CodeToAuthToken(code string) (string, error) {
	ctx := context.Background()
	authToken, err := oauthConf.Exchange(ctx, code)
	if err != nil {
		return "", err
	}
	return authToken.AccessToken, nil
}
