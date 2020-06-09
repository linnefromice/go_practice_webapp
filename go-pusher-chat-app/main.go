package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	pusher "github.com/pusher/pusher-http-go"
)

var client = pusher.Client {
	AppID: "1016122",
	Key: "14cf1c590d1ebf75c15c",
	Secret: "005dbca55fb85bc5a1cf",
	Cluster: "ap3",
	Secure: true,
}

type user struct {
	Name string `json:"name" xml:"name" form:"name" query:"name"`
	Email string `json:"email" xml:"email" form:"email" query:"email"`
}

func registerNewUser(rw http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	var newUser user
	err = json.Unmarshal(body, &newUser)
	if err != nil {
		panic(err)
	}

	client.Trigger("update", "new-user", newUser)
	json.NewEncoder(rw).Encode(newUser)
}

func pusherAuth(res http.ResponseWriter, req *http.Request) {
	params, _ := ioutil.ReadAll(req.Body)
	response, err := client.AuthenticatePrivateChannel(params)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(res, string(response))
}

func main()  {
	http.Handle("/", http.FileServer(http.Dir("./public")))

	http.HandleFunc("/new/user", registerNewUser)
	http.HandleFunc("pusher/auth", pusherAuth)
	log.Fatal(http.ListenAndServe(":8090", nil))
}