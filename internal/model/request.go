package model

type InsertUserReq struct {
	Name   string `bson:"name"`
	Age    int    `bson:"age"`
	UserId string `bson:"userId"`
}
