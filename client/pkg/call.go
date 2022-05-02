package faceit_cc_client

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/rpc"
	"time"
)

type User struct {
	Id         string     `json:"id"`
	FirstName  string     `json:"first_name"`
	LastName   string     `json:"last_name"`
	NickName   string     `gorm:"column:nickname" json:"nickname"`
	Password   string     `json:"password"`
	Email      string     `json:"email"`
	Country    string     `json:"country"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	Pagination Pagination `json:"pagination"`
}

type Pagination struct {
	SearchBy    string `json:"search_by"`
	SearchValue string `json:"search_value"`
	ResultsPage int    `json:"search_results_per_page"`
	Offset      int    `json:"offset"`
}

type Data struct {
	UserId   string `json:"userid,omitempty"`
	UserList []User `json:"userlist,omitempty"`
}

type Response struct {
	Status      int    `json:"status"`
	Description string `json:"description,omitempty"`
	Success     bool   `json:"success"`
	Data        Data   `json:"data,omitempty"`
}

func HealthCheck(client *rpc.Client) (int, error) {
	var healthresp int
	err := client.Call("status.HealthCheck", 1, &healthresp)
	if err != nil {
		return http.StatusNotFound, err
	}
	log.Printf("Test HealthCheck reply => %v\n", healthresp)
	return healthresp, nil
}

func Add(user User, client *rpc.Client) (Response, error) {
	reply := Response{}
	log.Println("Calling user.Add method")
	err := client.Call("user.Add", user, &reply)
	if err != nil {
		return reply, err
	}
	return reply, nil
}

func Update(user User, client *rpc.Client) (Response, error) {
	reply := Response{}
	log.Println("Calling user.Update method")
	err := client.Call("user.Update", user, &reply)
	if err != nil {
		return reply, err
	}
	return reply, nil
}

func Delete(user User, client *rpc.Client) (Response, error) {
	reply := Response{}
	log.Println("Calling user.Delete method")
	err := client.Call("user.Delete", user, &reply)
	if err != nil {
		return reply, err
	}
	return reply, nil
}

func List(user User, client *rpc.Client) (Response, error) {
	reply := Response{}
	log.Println("Calling user.List method")
	err := client.Call("user.List", user, &reply)
	if err != nil {
		return reply, err
	}
	return reply, nil
}

// PrettyPrint is a helper function to print structs
func PrettyPrint(input interface{}) {
	empJSON, err := json.MarshalIndent(input, "", "  ")
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Printf("PrettyPrint output \n %s\n", string(empJSON))
}
