package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/whitby/vcapi"
)

var (
	username = "api.whitby"
	password string
)

func init() {
	password = os.Getenv("VCAPI_PASSWORD")
	if password == "" {
		log.Fatal("VCAPI_PASSWORD not set")
	}

}

func main() {
	config := &vcapi.Config{Username: username, Password: password, SchoolID: "whitby", APIVersion: "v2"}
	client := vcapi.NewClient(config)
	req, err := client.NewRequest("students.json")
	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))

	fmt.Println(resp.Header)
	fmt.Println(client.Rate)
}
