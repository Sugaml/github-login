package src

import (
	"context"
	"io/ioutil"
	"net/http"

	"github.com/google/go-github/github"
	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
)

type GitHubClient struct {
	client *github.Client
}

type GitContent struct {
	Owner    string
	RepoName string
	Branch   string
	Path     string
	FileName string
}

func NewGitHub(token string) *GitHubClient {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	ctx := context.Background()

	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	return &GitHubClient{client: client}
}

func (git *GitHubClient) GetUser(ctx context.Context) (*github.User, error) {
	user, _, err := git.client.Users.Get(ctx, "")
	if err != nil {
		return nil, err
	}
	return user, err
}

func (git *GitHubClient) GetRepositories(ctx context.Context) ([]*github.Repository, error) {
	repos, _, err := git.client.Repositories.List(ctx, "", &github.RepositoryListOptions{
		Affiliation: "owner,collaborator",
		ListOptions: github.ListOptions{
			PerPage: 1000,
		},
	})
	return repos, err
}

func (git *GitHubClient) GetRepositoryByID(ctx context.Context, id int64) (*github.Repository, error) {
	repo, _, err := git.client.Repositories.GetByID(ctx, id)
	return repo, err
}

func (git *GitHubClient) GetBranches(ctx context.Context, owner string, repo string) ([]*github.Branch, *github.Response, error) {
	branches, res, err := git.client.Repositories.ListBranches(context.Background(), owner, repo, &github.ListOptions{PerPage: 1000})
	return branches, res, err
}

func (git *GitHubClient) GetBranch(ctx context.Context, owner string, repo string, bname string) (*github.Branch, *github.Response, error) {
	branch, res, err := git.client.Repositories.GetBranch(ctx, owner, repo, bname)
	return branch, res, err
}

func (git *GitHubClient) GetOrganizatons(ctx context.Context, user string) ([]*github.Organization, error) {
	org, _, err := git.client.Organizations.List(ctx, user, &github.ListOptions{})
	return org, err
}

func (git *GitHubClient) GetGithubContent(ctx context.Context, req *GitContent) (string, error) {
	body := ""
	_, s, _, err := git.client.Repositories.GetContents(ctx, req.Owner, req.RepoName, req.Path, &github.RepositoryContentGetOptions{
		Ref: req.Branch,
	})
	for _, st := range s {
		if *st.Type == "file" {
			logrus.Debugf("File Name :: %s Downlaod URL :: %s ", *st.Name, st.GetDownloadURL())
			if *st.Name == req.FileName {
				body, _ = getFileContent(st.GetDownloadURL())
			}
		}
	}
	return body, err
}

func getFileContent(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
