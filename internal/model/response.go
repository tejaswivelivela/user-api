package model

type InsertUserRes struct {
	Status string
}

type GetUserRes struct {
	Name   string `json:"name" bson:"name"`
	Age    int    `json:"age" bson:"age"`
	UserId string `json:"userId" bson:"userId"`
}

type UpdateUserRes struct {
	Status string
}
