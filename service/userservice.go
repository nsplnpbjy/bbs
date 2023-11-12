package service

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/nsplnpbjy/bbs/config"
	"github.com/nsplnpbjy/bbs/datamod"
	"github.com/nsplnpbjy/bbs/utils"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
)

func RegistUser(c *gin.Context) {
	username := string(c.PostForm("username"))
	password := string(c.PostForm("password"))
	encodedPassword := string(utils.PasswordEncrypt(password))
	collection := config.GetUserCollection()
	insertUser := datamod.User{}
	insertUser.Id = uuid.New().String()
	insertUser.Username = username
	insertUser.Password = encodedPassword
	insertUser.Regist_time = time.Now().String()
	if !datamod.CheckBlank(insertUser) {
		c.JSON(200, gin.H{"msg": "fail"})
		return
	}
	_, err := collection.InsertOne(context.TODO(), insertUser)
	if err != nil {
		log.Err(err)
		c.JSON(200, gin.H{"msg": "fail"})
	} else {
		c.JSON(200, gin.H{"msg": "done"})
	}
}

func LoginUser(c *gin.Context) {
	username := string(c.PostForm("username"))
	password := string(c.PostForm("password"))
	collection := config.GetUserCollection()
	resultUser := datamod.User{}
	result := collection.FindOne(context.TODO(), bson.M{"username": username})
	result.Decode(&resultUser)
	if utils.PasswordCompare(password, resultUser.Password) {
		c.JSON(200, gin.H{"msg": "login done"})
	} else {
		c.JSON(200, gin.H{"msg": "login fail"})
	}
}
