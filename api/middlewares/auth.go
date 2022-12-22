package middlewares

// import (
// 	"net/http"
// 	"strings"

// 	"github.com/gin-gonic/gin"
// 	"github.com/mitchellh/mapstructure"
// 	"qualifighting.backend.de/lib"
// 	"qualifighting.backend.de/models"
// )

// // Checks if the user is authenticated.
// // If not, it will return a 401 HTTP status code.
// // If the user is authenticated, it will set the user information to the context.
// func Auth(c *gin.Context) {
// 	authHeader := c.Request.Header.Get("Authorization")
// 	if authHeader == "" {
// 		c.Writer.Header().Set("WWW-Authenticate", "Bearer")
// 		c.AbortWithStatus(http.StatusUnauthorized)
// 		return
// 	}

// 	authHeaderParts := strings.Split(authHeader, " ")
// 	if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
// 		c.Writer.Header().Set("WWW-Authenticate", "Bearer")
// 		c.AbortWithStatus(http.StatusUnauthorized)
// 		return
// 	}

// 	jwt := authHeaderParts[1]

// 	auth := lib.GetFirebaseAuth()

// 	user, err := auth.VerifyIDToken(c.Request.Context(), jwt)

// 	if err != nil {
// 		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
// 		return
// 	}

// 	var jwtData *models.JWTUserMiddleware

// 	// map important values to struct, see models.JwTUser struct for more
// 	err = mapstructure.Decode(user.Claims, &jwtData)
// 	if err != nil {
// 		c.AbortWithStatusJSON(http.StatusUnauthorized, err)
// 		return
// 	}

// 	c.Set("user", jwtData)
// 	c.Next()
// }
