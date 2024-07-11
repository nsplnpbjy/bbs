package service

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/nsplnpbjy/bbs/config"
	"github.com/nsplnpbjy/bbs/datamod"
	"github.com/nsplnpbjy/bbs/utils"
	"go.mongodb.org/mongo-driver/bson"
)

func PostIdea(c *gin.Context) {
	token := c.PostForm("token")
	user := utils.ParseTokenGetUserInfo(token)
	text := c.PostForm("ideaText")
	idea := datamod.Idea{Id: uuid.New().String(), Post_time: time.Now().Unix(), Post_user_id: user.Id, Comments_id: nil, Text: text}
	collecion := config.GetIdeaCollection()
	_, error := collecion.InsertOne(context.TODO(), idea)
	if error != nil {
		c.JSON(http.StatusOK, idea.IdeaInsertedFailed())
		return
	} else {
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
			c.JSON(http.StatusOK, idea.IdeaDeleteSuccess(token))
			return
		}
	} else {
		c.JSON(http.StatusOK, idea.IdeaDeleteFailed())
		return
	}
}
