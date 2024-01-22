package middleware

import (
	//"4096Tiktok/controller"
	"4096Tiktok/dao"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

func JwtMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		TString := c.Query("token")
		if TString == "" {
			TString = c.PostForm("token")
		}
		// token为空
		if TString == "" {
			c.JSON(http.StatusOK, Response{StatusCode: 101, StatusMsg: "token为空"})
			c.Abort()
			return
		}

		token, claims, err := TokenParse(TString)

		if err != nil {
			if !token.Valid && strings.Contains(err.Error(), "invalid"){
				c.JSON(http.StatusOK, Response{StatusCode: 102, StatusMsg: "token错误"})
				c.Abort()
				return
			}else if !token.Valid && strings.Contains(err.Error(), "expired") {
				c.JSON(http.StatusOK, Response{StatusCode: 103, StatusMsg: "token过期"})
				c.Abort()
				return
			}
		}



		UserId := claims.UserId
		DB := dao.GetDB()
		var user dao.User
		DB.First(&user, UserId)

		// not registered
		if user.UserID == 0 {
			c.JSON(http.StatusOK, Response{StatusCode: 104, StatusMsg: "用户未注册"})
			c.Abort()
			return
		}

		// write
		c.Set("user", user)
		c.Next()
	}
}


func JwtMiddleWarePass() gin.HandlerFunc {
	return func(c *gin.Context) {
		TString := c.Query("token")
		if TString == "" {
			TString = c.PostForm("token")
		}
		// token为空
		if TString == "" {
			return
		}

		token, claims, err := TokenParse(TString)

		if err != nil {
			if !token.Valid && strings.Contains(err.Error(), "invalid"){
				c.JSON(http.StatusOK, Response{StatusCode: 102, StatusMsg: "token错误"})
				c.Abort()
				return
			}else if !token.Valid && strings.Contains(err.Error(), "expired") {
				c.JSON(http.StatusOK, Response{StatusCode: 103, StatusMsg: "token过期"})
				c.Abort()
				return
			}
		}



		UserId := claims.UserId
		DB := dao.GetDB()
		var user dao.User
		DB.First(&user, UserId)

		// not registered
		if user.UserID == 0 {
			c.JSON(http.StatusOK, Response{StatusCode: 104, StatusMsg: "用户未注册"})
			c.Abort()
			return
		}

		// write
		c.Set("user", user)
		c.Next()
	}
}


var JwtKey = []byte("114514")

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

func TokenRelease(user dao.User) (string, error) {
	ExpTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserId: user.UserID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: ExpTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "4096Tiktok",
			Subject:   "user token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	TokenString, err := token.SignedString(JwtKey)

	if err != nil {
		return "", err
	}

	return TokenString, nil
}

func TokenParse(TokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(TokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	fmt.Println("err is ", err)
	return token, claims, err

}