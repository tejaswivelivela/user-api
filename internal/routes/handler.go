package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"user-api/internal/model"

	"go.mongodb.org/mongo-driver/mongo"
)

func CreateNewUser(client *mongo.Client) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("content-type", "application/json")
		var request model.InsertUserReq
		err := json.NewDecoder(req.Body).Decode(&request)
		if err != nil {
			fmt.Println("Unable to decode Request", err)
		}
		fmt.Println(request)
		collection := client.Database("company").Collection("users")
		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		result, _ := collection.InsertOne(ctx, request)
		var response model.InsertUserRes
		if result != nil {
			response.Status = "New Document Inserted"
		} else {
			response.Status = "Unable to Insert"
		}
		json.NewEncoder(res).Encode(response)

	}
}
