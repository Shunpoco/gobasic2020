package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var url = "https://api.github.com/repos/"

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

func search(repo string, issueNum int) (*Issue, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s/issues/%d", url, repo, issueNum), nil)
	if err != nil {
		return nil, err
	}
	// TODO: tokenを別口で渡す
	req.Header.Add("Authorization", fmt.Sprintf("token %s", ""))
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}
