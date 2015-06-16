package main

import (
	"fmt"
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
	/*
		students, err := client.Students.List()
		if err != nil {
			log.Fatal(err)
		}

		for _, student := range students {
			fmt.Println(student.FirstName + " " + student.LastName)
		}
	*/
	student, err := client.Students.ID("288")
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("%+v", student.FirstName)
	fmt.Println(student.FirstName + " " + student.LastName)
	fmt.Println(student.Birthday)
	fmt.Println(student.UpdateDate)

}
