package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/quality-dashboard-api/api/services"
	"log"
	"net/http"
)


func handleLogin(w http.ResponseWriter, r *http.Request){

	fmt.Println("Inside handleLogin method")
}

func handleGetAllUsers(w http.ResponseWriter, r *http.Request)  {

	fmt.Println("Inside handleGetAllUsers method")

	w.Header().Set("Content-Type", "application/json")
	userBytes := services.GetAllUsers()

	w.Write(userBytes)
}

// RunServer - Runs the server on the given port
func RunServer(){

	const port string = ":8888"

	fmt.Println("Running the api now... ...")

	fmt.Println("Inside RunServer method")

	router := mux.NewRouter()

	//router.HandleFunc("/api/v1/login", handleLogin).Methods("GET")

	router.HandleFunc("/api/v1/getAllUsers", handleGetAllUsers).Methods("GET")

	log.Fatal(http.ListenAndServe(port, router))
}