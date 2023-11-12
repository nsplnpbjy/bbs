package config

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// 链接数据库
var (
	client            *mongo.Client
	err               error
	db                *mongo.Database
	userCollection    *mongo.Collection
	ideaCollection    *mongo.Collection
	commentCollection *mongo.Collection
)

// 1.建立连接
func DbInit() {

	if client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://comradegenrr:bjy653300@107.172.86.242:27017").SetConnectTimeout(5*time.Second)); err != nil {
		log.Err(err).Msg("链接数据库失败")
		panic("链接数据库失败")
	}
	//2.选择数据库
	db = client.Database("bbs")

	//3.选择表 user
	userCollection = db.Collection("user")

	//4.选择表 idea
	ideaCollection = db.Collection("idea")

	//5.选择标commentCollection
	commentCollection = db.Collection("comment")

}

func GetUserCollection() *mongo.Collection {
	return userCollection
}

func GetIdeaCollection() *mongo.Collection {
	return ideaCollection
}

func GetCommentCollection() *mongo.Collection {
	return commentCollection
}

func GetError() error {
	return err
}
