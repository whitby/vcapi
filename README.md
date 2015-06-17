[Go](http://golang.org/) client library for Veracross API

Documentation: http://godoc.org/github.com/whitby/vcapi

# Example usage

```Go
package main

import (
	"fmt"

	"github.com/whitby/vcapi"
)

func main() {
	// Configuration
	config := &vcapi.Config{
		Username:   "api.username",  // API Username
		Password:   "myAPIPassword", // API Password
		SchoolID:   "whitby",        // Client, school name
		APIVersion: "v2",            // Not a necessary field, API Version defaults to v2
	}

	// Create a new client with the above configuration.
	client := vcapi.NewClient(config)

	// Params are URL Parameters
	opt := &vcapi.ListOptions{Params: vcapi.Params{
		"option":        "0",
		"updated_after": "2015-06-01",
	}}

	// Query and Pagination
	// create a loop
	for {
		// request all parents
		parents, err := client.Parents.List(opt)
		if err != nil {
			// Handle any errors
			panic(err)
		}

		for _, parent := range parents {
			// Do something with individual result
			// For example, print the parent first name
			fmt.Println(parent.FirstName)

		}
		// Pagination. In this case increment the page by +1 as long as there's another
		// page in the results.
		opt.Page++

		if opt.NextPage == 0 {
			// break out of the loop when we reach the last page
			break
		}
	}
}
```

