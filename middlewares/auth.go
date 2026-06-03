package middlewares

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func VerifyAuthorization() gin.HandlerFunc {

	return func(context *gin.Context) {

		// Get jwt string from the header
		rawJwtString := context.GetHeader("Authorization")
		if rawJwtString == "" {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized Access."})
			return
		}

		// Strip Bearer from the jwt string
		jwtString := strings.TrimPrefix(rawJwtString, "Bearer ")

		// Parse jwt string into jwt token
		jwtToken, err := jwt.Parse(
			jwtString,
			func(token *jwt.Token) (any, error) {
				return []byte(os.Getenv("JWT_SECRET")), nil
			},
			jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}),
		)
		if err != nil {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
			return
		}

		// Get the token content
		claims, ok := jwtToken.Claims.(jwt.MapClaims)
		if !ok {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized Access."})
			return
		}

		// Set context to be used by subsequent routes
		context.Set("userId", claims["id"])
		context.Set("userEmail", claims["email"])

		// Continue to the next handler
		context.Next()
	}

}
