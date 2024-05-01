package model

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserBasic struct {
	Identity  string `bson:"_id"`
	Account   string `bson:"account"`
	Password  string `bson:"password"`
	Nickname  string `bson:"nickname"`
	Sex       int    `bson:"sex"`
	Email     string `bson:"email"`
	Avatar    string `bson:"avatar"`
	CreatedAt int64  `bson:"created_at"`
	UpdatedAt int64  `bson:"updated_at"`
}

func (UserBasic) CollectionName() string {
	return "user_basic"
}

func GetAccountAndPassword(account, password string) (*UserBasic, error) {
	ub := new(UserBasic)
	err := Mongo.Collection(UserBasic{}.CollectionName()).
		FindOne(context.Background(), bson.D{{Key: "account", Value: account}, {Key: "password", Value: password}}).
		Decode(ub)
	return ub, err
}

func GetUserDetail(identity primitive.ObjectID) (*UserBasic, error) {
	ub := new(UserBasic)
	err := Mongo.Collection(UserBasic{}.CollectionName()).
		FindOne(context.Background(), bson.D{{Key: "_id", Value: identity}}).
		Decode(ub)
	return ub, err
}

func GetEmailCnt(email string) (int64, error) {
	return Mongo.Collection(UserBasic{}.CollectionName()).
		CountDocuments(context.Background(), bson.D{{Key: "email"}})
}
