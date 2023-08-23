package jira

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/url"
	"testing"
)

type MockHttpClient struct {
	t             *testing.T
	CalledMethod  string
	CalledWith    []string
	CalledTimes   int
	CalledHeaders http.Header
	statusCode    int
}

func NewMockHttpClient(t *testing.T, statusCode int) *MockHttpClient {
	return &MockHttpClient{t: t, statusCode: statusCode, CalledTimes: 0, CalledWith: []string{}}
}

func (m *MockHttpClient) Do(req *http.Request) (*http.Response, error) {
	m.CalledTimes += 1
	data, err := io.ReadAll(req.Body)
	defer req.Body.Close()

	if err != nil {
		m.t.Fatal("error occurred while doing mock request")
	}

	m.CalledWith = append(m.CalledWith, string(data))
	m.CalledMethod = req.Method
	m.CalledHeaders = req.Header

	if m.statusCode == http.StatusBadRequest {
		body := bytes.NewReader([]byte("{\"errorMessages\":[],\"errors\":{\"name\":\"A version with this name already exists in this project.\"}}"))
		return &http.Response{Status: "Bad request", StatusCode: http.StatusBadRequest, Body: io.NopCloser(body)}, nil
	}
	body := bytes.NewReader([]byte("{\"hello\": \"world\"}"))
	return &http.Response{Status: "Created", StatusCode: http.StatusCreated, Body: io.NopCloser(body)}, nil
}

func TestNewClient(t *testing.T) {
	m := NewMockHttpClient(t, 200)
	parsedHost, _ := url.Parse("https://test.nu")

	c, err := NewClient("https://test.nu", "marcel@test.nl", "c0ffee", m)

	assert.NoError(t, err)

	expected := &Client{
		host:           parsedHost,
		httpClient:     m,
		authentication: nil,
	}

	expected.authentication = &authenticationService{
		client: expected,
		email:  "marcel@test.nl",
		token:  "c0ffee",
	}

	assert.Equal(t, expected, c)
}

func TestNewClient_missingHost(t *testing.T) {
	c, err := NewClient("", "marcel@test.nl", "c0ffee", nil)
	assert.Nil(t, c)
	assert.EqualError(t, err, "could not create jira client: hostname cannot be empty")
}

func TestNewClient_missingEmail(t *testing.T) {
	c, err := NewClient("https://test.nu", "", "c0ffee", nil)
	assert.Nil(t, c)
	assert.EqualError(t, err, "could not create jira client: email cannot be empty")
}

func TestNewClient_missingToken(t *testing.T) {
	c, err := NewClient("https://test.nu", "marcel@test.nl", "", nil)
	assert.Nil(t, c)
	assert.EqualError(t, err, "could not create jira client: token cannot be empty")
}
