package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/quality-dashboard-api/api"
	"github.com/quality-dashboard-api/api/config"
	"net/http"
)

func handleLogin(w http.ResponseWriter, r *http.Request){

	fmt.Println("Inside handleLogin method")
}

func main()  {

	config.ConnectToDB()
	//config.DB = DB

	//if(err != nil){
	//	fmt.Print("Error")
	//	panic(err.Error())
	//}

	// Added a new comment
	api.RunServer()
}