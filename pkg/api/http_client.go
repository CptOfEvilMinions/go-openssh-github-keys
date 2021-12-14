package api

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/CptOfEvilMinions/go-openssh-github-keys/pkg/config"
)

type Client struct {
	httpClient   http.Client
	GithubToken  string
	GithubAPIurl *url.URL
}

// Create global HTTP client
var HttpClient *Client

func (c *Client) Get(path string) (*http.Response, error) {
	// Combine Teamserver URL and path
	rel := &url.URL{Path: path}
	u := c.GithubAPIurl.ResolveReference(rel)

	// Create HTTP GET request
	req, err := http.NewRequest(
		http.MethodGet,
		u.String(),
		nil,
	)
	if err != nil {
		return nil, err
	}

	// Specify headers
	req.Header.Add("Accept", "application/vnd.github.v3+json")
	req.Header.Add("Authorization", fmt.Sprintf("token %s", c.GithubToken))

	// HTTP GET to Github API
	resp, err := c.httpClient.Do(req)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return resp, err
}

func InitHTTPclient(cfg *config.Config) {
	// Set TLS config
	tlsConfig := &tls.Config{
		InsecureSkipVerify: false,
		MinVersion:         tls.VersionTLS12,
	}

	// Create HTTP client
	http.DefaultTransport.(*http.Transport).TLSClientConfig = tlsConfig
	c := &Client{}

	// Convert string to URL
	var err error
	if c.GithubAPIurl, err = url.Parse("https://api.github.com"); err != nil {
		panic(err)
	}

	// Set the JWT token
	c.GithubToken = cfg.Token

	// Set HTTP client
	HttpClient = c
}
