package app

import (
	"awesomeProject/internal/app/ds"
	"awesomeProject/internal/app/role"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
)

const jwtPrefix = "Bearer"

func (app *Application) WithAuthCheck(assignedRoles ...role.Role) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		jwtStr := c.GetHeader("Authorization")
		if !strings.HasPrefix(jwtStr, jwtPrefix) {
			for _, oneOfAssignedRole := range assignedRoles {
				if role.NotAuthorized == oneOfAssignedRole {
					c.Set("userRole", role.NotAuthorized)
					return
				}
			}
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		//log.Println(jwtStr)
		//log.Println(jwtPrefix)
		jwtStr = jwtStr[len(jwtPrefix):]
		//log.Println(jwtStr)
		err := app.redisClient.CheckJWTInBlacklist(c.Request.Context(), jwtStr)
		if err == nil {
			c.AbortWithError(http.StatusForbidden, fmt.Errorf("Галя отмена!")) // по этому токену разлогинивались, не пускаем хакера
			log.Println("AAAAAAAAA")
			return
		}
		if !errors.Is(err, redis.Nil) { // значит что это не ошибка отсуствия - внутренняя ошибка
			c.AbortWithError(http.StatusInternalServerError, err)
			log.Println("BBBBBBBBB")
			return
		}

		claims := &ds.JWTClaims{}
		token, err := jwt.ParseWithClaims(jwtStr, claims, func(token *jwt.Token) (interface{}, error) {
			log.Println("CCCCCCCCCCc")
			return []byte(app.config.JWT.Token), nil
		})
		if err != nil || !token.Valid {
			c.AbortWithError(http.StatusForbidden, fmt.Errorf("Галя отмена!")) // токен устарел или глупый хакер ввел рандомные символы
			return
		}

		for _, oneOfAssignedRole := range assignedRoles {
			if claims.Role == oneOfAssignedRole {
				c.Set("userId", claims.UserUUID)
				c.Set("userRole", claims.Role)
				return
			}
		}
		c.AbortWithStatus(http.StatusForbidden)
		log.Println("role is not assigned")
	}

}
