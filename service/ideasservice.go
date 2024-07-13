package service

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/nsplnpbjy/bbs/config"
	"github.com/nsplnpbjy/bbs/datamod"
	"github.com/nsplnpbjy/bbs/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func PostIdea(c *gin.Context) {
	token := c.PostForm("token")
	user := utils.ParseTokenGetUserInfo(token)
	title := c.PostForm("ideaTitle")
	text := c.PostForm("ideaText")
	idea := datamod.Idea{Id: uuid.New().String(), Post_time: time.Now().Unix(), Post_user_id: user.Id, Comments_id: nil, Title: title, Text: text}
	collecion := config.GetIdeaCollection()
	_, error := collecion.InsertOne(context.TODO(), idea)
	if error != nil {
		c.JSON(http.StatusOK, idea.IdeaInsertedFailed())
		return
	} else {
		userinfo := new(datamod.User)
		config.GetUserCollection().FindOne(context.TODO(), bson.M{"id": user.Id}).Decode(userinfo)
		config.GetUserCollection().UpdateOne(context.TODO(), bson.M{"id": userinfo.Id}, bson.D{{Key: "$set", Value: bson.D{{Key: "ideas_id", Value: append(userinfo.Ideas_id, idea.Id)}}}})
		c.JSON(http.StatusOK, idea.IdeaInsertedSuccess(token))
		return
	}
}

func DeleteIdea(c *gin.Context) {
	token := c.PostForm("token")
	user := utils.ParseTokenGetUserInfo(token)
	idea_id := c.PostForm("ideaId")
	if idea_id == "" {
		c.JSON(http.StatusOK, new(datamod.Idea).IdeaDeleteFailed())
		return
	}
	idea := new(datamod.Idea)
	result := config.GetIdeaCollection().FindOne(context.TODO(), bson.M{"id": idea_id})
	result.Decode(&idea)
	if idea.Post_user_id == user.Id {
		_, err := config.GetIdeaCollection().DeleteOne(context.TODO(), bson.M{"id": idea_id})
		if err != nil {
			c.JSON(http.StatusOK, idea.IdeaDeleteFailed())
			return
		} else {
			userinfo := new(datamod.User)
			config.GetUserCollection().FindOne(context.TODO(), bson.M{"id": user.Id}).Decode(userinfo)
			config.GetUserCollection().UpdateOne(context.TODO(), bson.M{"id": userinfo.Id}, bson.D{{Key: "$set", Value: bson.D{{Key: "ideas_id", Value: utils.DeleteSlice(userinfo.Ideas_id, idea_id)}}}})
			c.JSON(http.StatusOK, idea.IdeaDeleteSuccess(token))
			return
		}
	} else {
		c.JSON(http.StatusOK, idea.IdeaDeleteFailed())
		return
	}
}

func ShowNewestIdeas(c *gin.Context) {
	token := c.PostForm("token")
	ideas := new(datamod.Ideas)
	startNum, _ := strconv.ParseInt(c.PostForm("startNum"), 10, 64)
	filter := bson.D{}
	opts := options.Find().SetSort(bson.D{{Key: "post_time", Value: -1}}).SetSkip(startNum).SetLimit(config.IdeasNum)
	cursor, err := config.GetIdeaCollection().Find(context.TODO(), filter, opts)
	if err != nil {
		c.JSON(http.StatusOK, ideas.IdeasSelectFailed())
		return
	}
	err = cursor.All(context.TODO(), ideas)
	if err != nil {
		c.JSON(http.StatusOK, ideas.IdeasSelectFailed())
		return
	}
	c.JSON(http.StatusOK, ideas.IdeasSelectSuccess(token))
}

func ShowAllIdeasByUserInfo(c *gin.Context) {
	token := c.PostForm("token")
	userid := utils.ParseTokenGetUserInfo(token).Id
	ideas := new(datamod.Ideas)
	filter := bson.M{"post_user_id": userid}
	opts := options.Find().SetSort(bson.D{{Key: "post_time", Value: -1}})
	cursor, err := config.GetIdeaCollection().Find(context.TODO(), filter, opts)
	if err != nil {
		c.JSON(http.StatusOK, ideas.IdeasSelectFailed())
		return
	}
	err = cursor.All(context.TODO(), ideas)
	if err != nil {
		c.JSON(http.StatusOK, ideas.IdeasSelectFailed())
		return
	}
	c.JSON(http.StatusOK, ideas.IdeasSelectSuccess(token))
}

func SearchIdeaByTitle(c *gin.Context) {
	token := c.PostForm("token")
	title := c.PostForm("ideaTitle")
	filter := bson.D{{Key: "title", Value: bson.D{{Key: "$regex", Value: title}}}}
	results, err := config.GetIdeaCollection().Find(context.TODO(), filter)
	if err != nil {
		println(err.Error())
		c.JSON(http.StatusOK, new(datamod.Ideas).IdeasSelectFailed())
		return
	}
	ideas := new(datamod.Ideas)
	if results.All(context.TODO(), ideas) != nil {
		c.JSON(http.StatusOK, new(datamod.Ideas).IdeasSelectFailed())
		return
	}
	c.JSON(http.StatusOK, ideas.IdeasSelectSuccess(token))
}

func SearchIdeaByText(c *gin.Context) {
	token := c.PostForm("token")
	text := c.PostForm("ideaText")
	filter := bson.D{{Key: "text", Value: bson.D{{Key: "$regex", Value: text}}}}
	results, err := config.GetIdeaCollection().Find(context.TODO(), filter)
	if err != nil {
		c.JSON(http.StatusOK, new(datamod.Ideas).IdeasSelectFailed())
		return
	}
	ideas := new(datamod.Ideas)
	if results.All(context.TODO(), ideas) != nil {
		c.JSON(http.StatusOK, new(datamod.Ideas).IdeasSelectFailed())
		return
	}
	c.JSON(http.StatusOK, ideas.IdeasSelectSuccess(token))
}
