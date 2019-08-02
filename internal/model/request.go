package model

type InsertUserReq struct {
	Name   string `bson:"name"`
	Age    int    `bson:"age"`
	UserId string `bson:"userId"`
}

type GetUserReq struct {
	UserId string `json:"UserId" bson:"userId"`
}

type UpdateUserReq struct {
	Name   string `bson:"name"`
	Age    int    `bson:"age"`
	UserId string `bson:"userId"`
}
