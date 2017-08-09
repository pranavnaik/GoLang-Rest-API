//Author: Pranav Naik
//Subject: GoLang API test

package main

import (
	"fmt"
	"log"
	"net/http"  //allow to access core tools to handle http requests
	"github.com/gorilla/mux"  //URL router and dispatcher for golang
	"encoding/json"
	"strconv"
)

type Item struct {
	Id 	int	`json:"Id"`
	Name string `json:"Name"`
	Quantity int `json:"quantity"`
	IsAvailable int `json:"IsAvailable"`
}

type Items []Item

func CheckForStock(w http.ResponseWriter, r *http.Request){

	//External package from "https://github.com/gorilla/mux" used for url parsing
	vars := mux.Vars(r)
	//Storing the url parameters in a map "vars"
	Name := vars["ItemName"]
	Id, _ := strconv.Atoi(vars["ItemId"])
	qty, _ := strconv.Atoi(vars["Quantity"])

	//Checking the quantity against the hardcoded 100 value..
	if qty <= 100 {
		Items := Items{
			Item{Id: Id, Name: Name, IsAvailable: 1},
		}
		json.NewEncoder(w).Encode(Items) //Encoding response in json format
	}else{
		Items := Items{
			Item{Id: Id, Name: Name, IsAvailable: 0},
		}
		json.NewEncoder(w).Encode(Items)
	}

}

func handleRequests() {

	//Http router created using mux library
	myRouter := mux.NewRouter().StrictSlash(true)
	//Defining the handler
	myRouter.HandleFunc("/check-inventory-levels/{ItemName}/{ItemId}/{Quantity}/", CheckForStock)
	//listen to API requests on port
	log.Fatal(http.ListenAndServe(":8081", myRouter))

}


func main() {

	fmt.Println("Rest API to check for a qantity from availabe stock")
	handleRequests()

}
