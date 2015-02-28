package vcapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Base URL for API calls
func (v *VeracrossSettings) BaseURL() string {
	url := fmt.Sprintf("https://%v:%v@api.veracross.com/%v/v1", v.Username, v.Password, v.Client)
	return url
}

type Fetcher interface {
	Fetch()
}

// Fetches JSON from server and Decodes into type.
func fetch(url string, into interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(into); err != nil {
		return err
	}

	return nil
}
