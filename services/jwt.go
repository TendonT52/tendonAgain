package services

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
)

type JwtServices struct {
	secretKey string
}

func NewJwtServices(secret string) *JwtServices {
	return &JwtServices{
		secretKey: secret,
	}
}

type authCustomClaims struct {
	Name          string `json:"name" bson:"name"`
	Surname       string `json:"surname" bson:"surname"`
	Email         string `json:"email" bson:"email"`
	Curriculum_id []int  `json:"curriculum" bson:"curriculum_id"`
	jwt.StandardClaims
}

func (service *JwtServices) GenerateToken(user User) string {
	claims := &authCustomClaims{
		Email:     user.Email,
		Name:      user.Name,
		Surname:   user.Surname,
		Curriculum_id: user.Curriculum_id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(service.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func (service *JwtServices) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("Invalid token %s", token.Header["alg"])
		}
		return []byte(service.secretKey), nil
	})

}
