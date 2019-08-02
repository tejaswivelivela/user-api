package getuser

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"user-api/internal/model"

	"github.com/gorilla/mux"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetUserEndpoint(client *mongo.Client) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("content-type", "application/json")
		vars := mux.Vars(req)
		id := vars["id"]
		//	var request model.GetUserReq
		collection := client.Database("company").Collection("users")
		ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
		var response model.GetUserRes

		err := collection.FindOne(ctx, bson.M{"userId": id}).Decode(&response)
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			fmt.Println("err:", err)
			res.Write([]byte(fmt.Sprintf(`{"message":"%v"}`, err)))
			return
		}
		json.NewEncoder(res).Encode(response)
	}
}
