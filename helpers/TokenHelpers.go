package helpers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"os"
	"strings"
)

func Middleware(c *fiber.Ctx) error {
	splitToken := strings.Split(c.Get("Authorization"), "Bearer ")
	tokenString := splitToken[1]
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT_KEY")), nil
	})
	if err != nil {
		res := GetResponse(fiber.StatusUnauthorized, nil, err)
		return c.Status(res.Status).JSON(res)
	}
	//c.Locals("SiteId", claims["siteid"].(string))
	//c.Locals("UserId", claims["userid"].(string))
	//c.Locals("OrgId", claims["orgid"].(string))
	c.Locals("RoleId", claims["userroleid"])
	return c.Next()
}
