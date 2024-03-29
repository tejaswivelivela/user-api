package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
	"user-api/internal/config"
	"user-api/internal/routes/getuser"
	"user-api/internal/routes/insertuser"
	"user-api/internal/routes/updateuser"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	con, err := config.ReadConfig("config.json")
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	//"mongodb://localhost:27017"
	connectionString := fmt.Sprintf("mongodb://%v:%v", con.Database.HostName, con.Database.Port)
	clientOptions := options.Client().ApplyURI(connectionString)
	client, _ := mongo.Connect(ctx, clientOptions)
	//fmt.Println("My Configuration details:", con)
	router := mux.NewRouter()
	//	subpath := router.PathPrefix("/v1")
	router.HandleFunc("/insertuser", insertuser.CreateNewUser(client)).Methods(http.MethodPost)
	router.HandleFunc("/getuser/{id}", getuser.GetUserEndpoint(client)).Methods(http.MethodGet)
	router.HandleFunc("/updateuser", updateuser.UpdateUser(client)).Methods(http.MethodPut)
	http.ListenAndServe(":8000", router)
}
