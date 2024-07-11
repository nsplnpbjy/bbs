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

// 登录逻辑
func RegistUser(c *gin.Context) {
	username := string(c.PostForm("username"))
	password := string(c.PostForm("password"))
	encodedPassword := string(utils.PasswordEncrypt(password))
	collection := config.GetUserCollection()
	insertUser := datamod.User{Id: uuid.New().String(), Username: username, Password: encodedPassword, Regist_time: time.Now().Unix(), Ideas_id: nil, Comments_id: nil}
	if !insertUser.CheckBlank() {
		c.JSON(http.StatusOK, insertUser.FailReturner())
		return
	}
	result := collection.FindOne(context.TODO(), bson.M{"username": username})
	finduser := new(datamod.User)
	result.Decode(&finduser)
	if finduser.Id != "" {
		c.JSON(http.StatusOK, insertUser.UsernameAlreadyExist())
		return
	}
	_, err := collection.InsertOne(context.TODO(), insertUser)
	if err != nil {
		log.Err(err)
		c.JSON(http.StatusOK, insertUser.DePassword().FailReturner())
	} else {
		token, _ := utils.GenerateToken(insertUser.Username, insertUser.Password)
		c.JSON(http.StatusOK, insertUser.DePassword().SuccessReturner(token))
	}
}

// 注册逻辑
func LoginUser(c *gin.Context) {
	username := string(c.PostForm("username"))
	password := string(c.PostForm("password"))
	collection := config.GetUserCollection()
	resultUser := datamod.User{}
	result := collection.FindOne(context.TODO(), bson.M{"username": username})
	result.Decode(&resultUser)
	if utils.PasswordCompare(password, resultUser.Password) {
		token, _ := utils.GenerateToken(resultUser.Username, resultUser.Password)
		c.JSON(http.StatusOK, resultUser.DePassword().SuccessReturner(token))
	} else {
		c.JSON(http.StatusOK, resultUser.DePassword().FailReturner())
	}
}
