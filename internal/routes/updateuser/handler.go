package updateuser

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"user-api/internal/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func UpdateUser(client *mongo.Client) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("content-type", "application/json")
		var request model.UpdateUserReq
		err := json.NewDecoder(req.Body).Decode(&request)
		if err != nil {
			fmt.Println("Unable to decode Request", err)
		}
		//fmt.Println(request)
		collection := client.Database("company").Collection("users")
		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		query := bson.M{"userId": request.UserId}
		update := bson.M{"age": request.Age}
		UpdateResult, err := collection.UpdateOne(ctx, query, bson.M{"$set": update})
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			fmt.Println("err:", err)
			res.Write([]byte(fmt.Sprintf(`{"message":"%v"}`, err)))
			return
		}
		var response model.UpdateUserRes
		fmt.Println(UpdateResult)
		if UpdateResult.ModifiedCount > 0 {
			response.Status = "Succesfully Updated the document"
		} else {
			response.Status = "Unable to Update"
		}
		json.NewEncoder(res).Encode(response)

	}
}
