package controllers

import (
	"log"
	"project_koperasi/models"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

var identitas = "id"

// coba handler
func HelloHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	User, _ := c.Get(identitas)
	c.JSON(200, gin.H{
		"userID":   claims[identitas],
		"UserName": User.(*models.User).UserName,
		"text":     "Hello World",
	})
}

// auth midlleware
func AuthMidlleware(c *gin.Context) {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("Secret Key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identitas,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*models.User); ok {
				return jwt.MapClaims{
					identitas: v.UserName,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &models.User{
				UserName: claims[identitas].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals models.Login
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			UserID := loginVals.Username
			Password := loginVals.Password

			if UserID == "admin" && Password == "admin" || UserID == "test" && Password == "test" {
				return &models.User{
					UserName:  UserID,
					LastName:  "Bo-Yi",
					FirstName: "Wu",
				}, nil
			}
			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if v, ok := data.(*models.User); ok && v.UserName == "admin" {
				return true
			}
			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		TokenLookup:   "header: Authorization, query:token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})

	if err != nil {
		log.Fatal("JWT Error : " + err.Error())
	}

	errInit := authMiddleware.MiddlewareInit()

	if errInit != nil {
		log.Fatal("authMiddlewareInit Error : " + errInit.Error())
	}
	c.JSON(200, gin.H{"Status": "Oke"})
}
