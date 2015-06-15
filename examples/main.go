package main

import (
	"fmt"

	"github.com/groob/vcapi"
)

var (
	username = "user"
	password = "password"
)

func main() {
	config := &vcapi.Config{Username: username, Password: password, SchoolID: "whitby"}
	client := vcapi.NewClient(config)
	fmt.Println(client)
}
