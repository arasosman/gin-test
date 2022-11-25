package auth

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"os"
	"strings"
)

func TokenValid(c *gin.Context) error {
	tokenString := extractToken(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return err
	}
	Claims(c, token)
	return nil
}

func Claims(c *gin.Context, token *jwt.Token) {

	result := Profile{}
	// convert map to json
	jsonString, _ := json.Marshal(token.Claims)
	json.Unmarshal(jsonString, &result)
	c.Set("profile", result)
}

func extractToken(c *gin.Context) string {
	token := c.Query("token")
	if token != "" {
		return token
	}
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func GetProfile(c *gin.Context) Profile {
	var profile Profile
	val, _ := c.Get("profile")
	profile = val.(Profile)
	return profile
}

type Profile struct {
	OrganizationType string `json:"organization_type"`
	OrganizationId   int    `json:"organization_id"`
	UserId           int    `json:"sub"`
}
