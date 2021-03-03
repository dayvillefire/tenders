package auth

import (
	"log"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/dayvillefire/tenders/config"
	"github.com/dayvillefire/tenders/models"
	"github.com/gin-gonic/gin"
)

var (
	AuthMiddleware        *jwt.GinJWTMiddleware
	middlewareInitialized bool
	IdentityKey           = "authUserID"
)

type login struct {
	ShortCode string `form:"shortcode" json:"shortcode" binding:"required"`
	PIN       string `form:"pin" json:"pin" binding:"required"`
}

func helloHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	user, _ := c.Get(IdentityKey)
	c.JSON(200, gin.H{
		"userID":   claims[IdentityKey],
		"userName": user.(*models.User).FirstName + " " + user.(*models.User).LastName,
	})
}

func GetAuthMiddleware() *jwt.GinJWTMiddleware {
	if middlewareInitialized {
		return AuthMiddleware
	}
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "Tenders",
		Key:         []byte(config.Config.Auth.StoreSecret),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: IdentityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*models.User); ok {
				return jwt.MapClaims{
					IdentityKey: v.ShortCode,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			user, _ := models.UserByShortCode(claims[IdentityKey].(string))
			return &user
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals login
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}

			log.Printf("auth.Authenticator: %#v", loginVals)

			userID := loginVals.ShortCode
			password := loginVals.PIN

			user, err := models.UserAuth(userID, password)
			log.Printf("auth.Authenticator: %#v, %#v", user, err)
			if err == nil {
				return user, nil
			}

			log.Printf("auth.Authenticator: ERR: %s", err.Error())

			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if v, ok := data.(*models.User); ok && v.ShortCode != "" {
				log.Printf("auth.Authorizator: TRUE")
				return true
			}

			log.Printf("auth.Authorizator: FALSE")
			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	// Cache for next go-around
	AuthMiddleware = authMiddleware
	middlewareInitialized = true

	return authMiddleware
}

/*
	r.POST("/login", authMiddleware.LoginHandler)

	r.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	auth := r.Group("/auth")
	// Refresh time can be longer than token timeout
	auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/hello", helloHandler)
	}
*/
