package jira

import "net/http"

// authenticationService is used to authenticate requests
type authenticationService struct {
	client *Client
	email  string
	token  string
}

// setBasisAuth sets the username and password on the provided request
func (a *authenticationService) setBasicAuth(req *http.Request) {
	req.SetBasicAuth(a.email, a.token)
}
