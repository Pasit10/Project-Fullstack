package middlewares

import (
	templateError "backend/error"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(os.Getenv("SECRET_KEY"))

func SetJWTHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenStr := c.Cookies("jwt")
		if tokenStr == "" {
			httpstatuscode, errorResponse := templateError.GetErrorResponse(templateError.MissingOrMalformedToken)
			return c.Status(httpstatuscode).JSON(errorResponse)
		}

		token, err := ParseJWT(tokenStr)
		if err != nil || !token.Valid {
			httpstatuscode, errorResponse := templateError.GetErrorResponse(templateError.InvalidOrExpiredToken)
			return c.Status(httpstatuscode).JSON(errorResponse)
		}

		claims := token.Claims.(jwt.MapClaims)
		c.Locals("uid", claims["uid"])
		c.Locals("name", claims["name"])
		c.Locals("role", claims["role"])

		return c.Next()
	}
}
func GenerateJWT(uid string, name string, role string) (string, error) {
	claims := jwt.MapClaims{
		"uid":  uid,
		"name": name,
		"role": role,
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ParseJWT(tokenStr string) (*jwt.Token, error) {
	return jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
}
