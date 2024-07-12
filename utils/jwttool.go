package utils

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/nsplnpbjy/bbs/config"
	"github.com/nsplnpbjy/bbs/datamod"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	jwtSecret []byte = []byte("comradegenrr")
	issuer    string = "comradegenrr"
)

type Claims struct {
	Username string
	Password string
	jwt.StandardClaims
}

// 根据用户的用户名和密码参数token
func GenerateToken(username string, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Minute * 15).Unix()

	claims := Claims{
		Username: username,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime, // 过期时间
			Issuer:    issuer,     //指定发行人
		},
	}
	// 该方法内部生成签名字符串，再用于获取完整、已签名的token
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

// 根据传入的token值获取到Claims对象信息(进而获取其中的用户名和密码)
func ParseToken(token string) (*Claims, error) {
	// 用于解析鉴权的声明，方法内部主要是具体的解码和校验的过程，最终返回*Token
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		// 从tokenClaims中获取到Claims对象，并使用断言，将该对象转换为我们自己定义的Claims
		// 要传入指针，项目结构体都是用指针传递，节省空间
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid { // Valid()验证基于时间的声明
			return claims, nil
		}
	}
	return nil, err
}

func RefreashToken(token string) string {
	claims, _ := ParseToken(token)
	if (claims.ExpiresAt - time.Now().Unix()) < 300 {
		println("in")
		user := ParseTokenGetUserInfo(token)
		newToken, _ := GenerateToken(user.Username, user.Password)
		return newToken
	}
	return token
}

func ParseTokenGetUserInfo(token string) *datamod.User {
	clams, _ := ParseToken(token)
	username := clams.Username
	collection := config.GetUserCollection()
	result := collection.FindOne(context.TODO(), bson.M{"username": username})
	finduser := new(datamod.User)
	result.Decode(&finduser)
	return finduser.DePassword()
}
