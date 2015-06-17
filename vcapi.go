package vcapi

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"net/url"
	"strconv"
)

const (
	libraryVersion = "0.1.0"
	defaultBaseURL = "https://api.veracross.com/"
	userAgent      = "vcapi/" + libraryVersion
	mediaType      = "application/json"
	format         = "json"

	headerRateLimit     = "X-Rate-Limit-Limit"
	headerRateRemaining = "X-Rate-Limit-Remaining"
	headerRateReset     = "X-Rate-Limit-Reset"
	headerCountTotal    = "X-Total-Count"
)

type Params map[string]string
type ListOptions struct {
	Page     int
	NextPage int
	Params   Params
}

type Rate struct {
	// The number of request per hour the client is currently limited to.
	Limit int `json:"limit"`

	// The number of remaining requests the client can make this hour.
	Remaining int `json:"remaining"`

	// The time at w\hic the current rate limit will reset.
	Reset int `json:"reset"`
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

	Students StudentService
	Parents  ParentService
	FacStaff FacStaffService
}

func NewClient(config *Config) *Client {

	// Default to API Version 2
	if config.APIVersion == "" {
		config.APIVersion = "v2"
	}

	baseURL, _ := url.Parse(defaultBaseURL)
	// add Version and SchoolID to URL Path
	baseURL.Path = config.SchoolID + "/" + config.APIVersion + "/"

	c := &Client{client: http.DefaultClient, BaseURL: baseURL, UserAgent: userAgent, Config: config}

	c.Students = StudentService{client: c}
	c.Parents = ParentService{client: c}
	c.FacStaff = FacStaffService{client: c}
	return c
}

// NewRequest creates an API request. A relative URL can be provided in urlStr, which will be resolved to the
// BaseURL of the Client. Relative URLS should always be specified without a preceding slash.
func (c *Client) NewRequest(urlStr string) (*http.Request, error) {
	method := "GET"
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)

	req, err := http.NewRequest(method, u.String(), nil)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.Config.Username, c.Config.Password)

	req.Header.Add("Accept", mediaType)
	req.Header.Add("User-Agent", userAgent)
	return req, nil
}

// Do sends an API request and returns the API response.
func (c *Client) Do(req *http.Request, into interface{}) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	// Set rate limits
	if limit := resp.Header.Get(headerRateLimit); limit != "" {
		c.Rate.Limit, _ = strconv.Atoi(limit)
	}

	if remaining := resp.Header.Get(headerRateRemaining); remaining != "" {
		c.Rate.Remaining, _ = strconv.Atoi(remaining)
	}

	if reset := resp.Header.Get(headerRateReset); reset != "" {
		c.Rate.Reset, _ = strconv.Atoi(reset)
	}

	if err := json.NewDecoder(resp.Body).Decode(into); err != nil {
		return nil, err
	}

	return resp, nil
}

func addOptions(basePath, format string, opt *ListOptions) string {
	// Specify URL Parameters
	params := url.Values{}
	for k, v := range opt.Params {
		params.Add(k, v)
	}
	// only set format if not already specified by options
	if _, ok := opt.Params["format"]; !ok {
		params.Set("format", format)
	}

	// Sets the page which should be retrieved.
	if page := opt.Page; opt.Page != 0 {
		params.Set("page", fmt.Sprintf("%v", page))
	}

	path := basePath + "?" + params.Encode()
	return path
}

func paginate(resp *http.Response, opt *ListOptions) {
	if recordCount := resp.Header.Get(headerCountTotal); recordCount != "" {
		count, _ := strconv.Atoi(recordCount)

		// number of pages, rounded up
		pages := math.Floor((float64(count) / 100.0) + .9)
		// update NextPage number
		if pages != 1 {
			opt.NextPage = opt.Page + 1
		}
		if float64(opt.Page) >= pages {
			opt.NextPage = 0
		}

	}
}
