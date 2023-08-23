package jira

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// Client allows the program to interact with the JIRA API
type Client struct {
	host           *url.URL
	httpClient     HttpClient
	authentication *authenticationService
}

// HttpClient is the http client interface used by the JIRA client
type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

func NewClient(host, email, token string, httpClient HttpClient) (*Client, error) {
	if host == "" {
		return nil, errors.New("could not create jira client: hostname cannot be empty")
	}

	if email == "" {
		return nil, errors.New("could not create jira client: email cannot be empty")
	}

	if token == "" {
		return nil, errors.New("could not create jira client: token cannot be empty")
	}

	u, err := url.Parse(host)

	if err != nil {
		return nil, fmt.Errorf("could not create jira client, invalid host provided: %w", err)
	}

	client := &Client{host: u, httpClient: httpClient}
	client.authentication = &authenticationService{
		client: client,
		email:  email,
		token:  token,
	}
	return client, nil
}

func (c *Client) GetIssueDescription(key string) (*IssueDescriptionResponse, error) {
	endpoint, _ := url.Parse(fmt.Sprintf("/rest/api/3/issue/%v?fields=description", key))
	u := c.host.ResolveReference(endpoint).String()

	req, err := http.NewRequest("GET", u, nil)

	if err != nil {
		return nil, err
	}

	c.authentication.setBasicAuth(req)
	req.Header.Add("Content-Type", "application/json")

	res, err := c.httpClient.Do(req)

	if err != nil {
		return nil, err
	}

	var response IssueDescriptionResponse

	data, err := io.ReadAll(res.Body)
	defer res.Body.Close()

	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
