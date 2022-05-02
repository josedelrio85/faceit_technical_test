package main

import (
	"log"
	"net/rpc"
	"time"

	"github.com/google/uuid"
	faceit_cc_client "github.com/josedelrio85/faceit_technical_test/client/pkg"
)

func main() {
	log.Println("Starting FaceIT rpc client ...")
	hostname := "localhost"
	port := ":9001"

	client, err := rpc.DialHTTP("tcp", hostname+port)
	if err != nil {
		log.Fatal("dialing: ", err)
	}
	time.Sleep(2 * time.Second)

	log.Println("Status of server -> HealthCheck")
	_, err = faceit_cc_client.HealthCheck(client)
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Println("-------------------------------")
	time.Sleep(2 * time.Second)

	log.Println("Adding new user -> user.Add")
	id := uuid.New().String()
	newuser := faceit_cc_client.User{
		Id:        id,
		FirstName: "zzzzzzzz",
		LastName:  "asdfasdf",
		NickName:  "asdfasdf",
		Password:  "asdfasdf",
		Email:     "asdfasdf@asdfas.es",
		Country:   "ES",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	faceit_cc_client.PrettyPrint(newuser)
	resp, err := faceit_cc_client.Add(newuser, client)
	if err != nil {
		log.Fatalf(err.Error())
	}
	faceit_cc_client.PrettyPrint(resp)
	log.Println("-------------------------------")
	time.Sleep(2 * time.Second)

	log.Println("Updating user -> user.Update")
	updateuser := faceit_cc_client.User{
		Id:        id,
		FirstName: "zzzzzzzz2",
		LastName:  "asdfasdf2",
		NickName:  "asdfasdf2",
		Password:  "asdfasdf2",
		Email:     "asdfasdf@asdfas.es",
		Country:   "ES",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	faceit_cc_client.PrettyPrint(updateuser)
	resp, err = faceit_cc_client.Update(updateuser, client)
	if err != nil {
		log.Fatalf(err.Error())
	}
	faceit_cc_client.PrettyPrint(resp)
	log.Println("-------------------------------")
	time.Sleep(2 * time.Second)

	log.Println("Deleting user -> user.Delete")
	deleteuser := faceit_cc_client.User{
		Id: id,
	}
	faceit_cc_client.PrettyPrint(deleteuser)
	resp, err = faceit_cc_client.Delete(deleteuser, client)
	if err != nil {
		log.Fatalf(err.Error())
	}
	faceit_cc_client.PrettyPrint(resp)
	log.Println("-------------------------------")
	time.Sleep(2 * time.Second)

	log.Println("Listing users by `country=UK`, 5 results per page and offset 0 -> user.List")
	// List users by country = UK, 5 results per page and offset 0
	listuser := faceit_cc_client.User{
		Pagination: faceit_cc_client.Pagination{
			SearchBy:    "country",
			SearchValue: "UK",
			ResultsPage: 5,
			Offset:      0,
		},
	}
	faceit_cc_client.PrettyPrint(listuser)
	resp, err = faceit_cc_client.List(listuser, client)
	if err != nil {
		log.Fatalf(err.Error())
	}
	faceit_cc_client.PrettyPrint(resp)
	log.Printf("Length of retrieved list => %d", len(resp.Data.UserList))
	log.Println("-------------------------------")
	time.Sleep(2 * time.Second)
	log.Println("Closing client... Goodbye!")
}
