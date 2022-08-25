package services

import (
	"time"

	"github.com/TendonT52/tendon-api/controllers"
	"github.com/dgrijalva/jwt-go/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	Id           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	jwt.StandardClaims
}

func (service *JwtServices) GenerateAccessToken(user controllers.SignInUser) string {
	claims := &authCustomClaims{
		Id:           user.Id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(time.Minute*10)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, _ := token.SignedString([]byte(service.secretKey))
	return t
}

func (service *JwtServices) GenerateRefreshToken(user controllers.SignInUser) string {
	claims := &authCustomClaims{
		Id:           user.Id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(time.Hour*72)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, _ := token.SignedString([]byte(service.secretKey))
	return t
}
//TODO write test for validateToken
func (service *JwtServices) ValidateToken(encodedToken string) (*jwt.Token, error) {
  token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
     if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
        return nil, jwt.ErrSignatureInvalid
     }
     return []byte(service.secretKey), nil
  })
  if err != nil {
     return nil, err
  }
  return token, nil
}