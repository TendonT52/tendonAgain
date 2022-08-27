package services

import (
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"golang.org/x/xerrors"
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
	Id string `json:"id" bson:"_id,omitempty"`
	jwt.StandardClaims
}

func (service *JwtServices) GenerateAccessToken(userId string,  jwtId string, duration time.Duration) string {
	claims := &authCustomClaims{
		Id: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(duration)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, _ := token.SignedString([]byte(service.secretKey))
	return t
}

func (service *JwtServices) GenerateRefreshToken(userId string) string {
	claims := &authCustomClaims{
		Id: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(time.Hour * 72)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, _ := token.SignedString([]byte(service.secretKey))
	return t
}

func (service *JwtServices) ValidateToken(encodedToken string) (string, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(service.secretKey), nil
	})
	var uErr *jwt.UnverfiableTokenError
	var expErr *jwt.TokenExpiredError
	var nbfErr *jwt.TokenNotValidYetError
	if token.Valid {
		claim := token.Claims.(jwt.MapClaims)
		return claim["id"].(string), nil
	} else if xerrors.As(err, &uErr) {
		return "", UnverfiableToken.From(err)
	} else if xerrors.As(err, &expErr) {
		return "", TokenExpired.From(err)
	} else if xerrors.As(err, &nbfErr) {
		return "", TokenNotValidYet.From(err)
	} else {
		return "", JwtError.From(err)
	}
}
