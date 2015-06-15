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

	headerRateLimit     = "X-RateLimit-Limit"
	headerRateRemaining = "X-RateLimit-Remaining"
	headerRateReset     = "X-RateLimit-Reset"
)

type Rate struct {
	// The number of request per hour the client is currently limited to.
	Limit int `json:"limit"`

	// The number of remaining requests the client can make this hour.
	Remaining int `json:"remaining"`

	// The time at w\hic the current rate limit will reset.
	Reset Timestamp `json:"reset"`
}

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

	// Rate contains the current rate limit for the client as determined by the most recent
	// API call.
	Rate Rate

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
