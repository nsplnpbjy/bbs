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
	if !insertUser.CheckBlank() {
		c.JSON(http.StatusOK, insertUser.FailReturner())
		return
	}
	_, err := collection.InsertOne(context.TODO(), insertUser)
	if err != nil {
		log.Err(err)
		c.JSON(http.StatusOK, insertUser.FailReturner())
	} else {
		insertUser.Password = ""
		c.JSON(http.StatusOK, insertUser.SuccessReturner())
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
		resultUser.Password = ""
		c.JSON(http.StatusOK, resultUser.SuccessReturner())
	} else {
		c.JSON(http.StatusOK, resultUser.FailReturner())
	}
}
