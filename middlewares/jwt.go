package middlewares

import (
	"encoding/json"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go-learning/utils/common"
	"go-learning/utils/logger"
	"go-learning/utils/redis"
	"strings"
	"time"
)

type Claims struct {
	jwt.StandardClaims
}

func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			redisSessionStr string
			ok              bool
		)

		tokenString, err := GetJwtTokenFromHeader(c)
		if err != nil {
			common.GenerateErrorResponse(c, err.Error())
			return
		}

		// select mode app
		if viper.GetString("app.mode") == "development" {
			redisSessionStr, ok = DummyRedis[tokenString]
			if !ok {
				common.GenerateErrorResponse(c, "token invalid, please log in again")
				return
			}
		} else {
			redisSessionStr, err = redis.RedisClient.Get(c, tokenString).Result()
			if err != nil {
				err = errors.New("redis error")
				logger.ErrorWithCtx(c, nil, err)
				common.GenerateErrorResponse(c, err.Error())
				return
			}
		}

		var redisSession RedisSession
		err = json.Unmarshal([]byte(redisSessionStr), &redisSession)
		if err != nil {
			common.GenerateErrorResponse(c, "failed unmarshal redis session")
			return
		}

		if time.Now().After(redisSession.ExpiredAt) {
			common.GenerateErrorResponse(c, "token expired, please log in again")
			return
		}

		c.Next()
	}
}

func GetJwtTokenFromHeader(c *gin.Context) (tokenString string, err error) {
	authHeader := c.Request.Header.Get("Authorization")

	if common.IsEmptyField(authHeader) {
		return tokenString, errors.New("authorization header is required")
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return tokenString, errors.New("invalid Authorization header format")
	}

	return parts[1], nil
}

func GenerateJwtToken() (token string, err error) {
	expirationTime := time.Now().Add(1 * time.Minute)

	claims := &Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	generatedTokenJwt := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err = generatedTokenJwt.SignedString([]byte(viper.GetString("jwt_secret_key")))
	if err != nil {
		return
	}

	return
}
