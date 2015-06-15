package vcapi

import (
	"net/http"
	"net/url"
)

const (
	libraryVersion = "0.1.0"
	defaultBaseURL = "https://api.veracross.com/"
	userAgent      = "vcapi/" + libraryVersion
	mediaType      = "application/json"
)

type Config struct {
	Username   string
	Password   string
	SchoolID   string
	APIVersion string
}

type Client struct {
	// HTTP client used to communicate with the Veracross API.
	client *http.Client

	// Base URL for API requests.
	BaseURL *url.URL

	// User agent for client
	UserAgent string

	// Username, Password and Client
	Config *Config
}

func NewClient(config *Config) *Client {

	// Default to API Version 2
	if config.APIVersion == "" {
		config.APIVersion = "v2"
	}

	baseURL, _ := url.Parse(defaultBaseURL)
	// add Version and SchoolID to URL Path
	baseURL.Path = config.APIVersion + "/" + config.SchoolID + "/"

	c := &Client{client: http.DefaultClient, BaseURL: baseURL, UserAgent: userAgent, Config: config}

	return c
}
