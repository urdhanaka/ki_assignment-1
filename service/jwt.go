package service

import (
	"crypto/rand"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type JWTService interface {
	GenerateToken(userID uuid.UUID) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
	FindUserIDByToken(token string) (uuid.UUID, error)
}

type jwtService struct {
	secretKey string
	issuer    string
}

func NewJWTService() JWTService {
	return &jwtService{
		secretKey: getSecretKey(),
		issuer:    "go-rest-api",
	}
}

func getSecretKey() string {
	JWTsecretKey := make([]byte, 32)
	if _, err := rand.Read(JWTsecretKey); err != nil {
		panic(err)
	}

	return string(JWTsecretKey)
}

func (j *jwtService) GenerateToken(userID uuid.UUID) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = userID
	claims["issuer"] = j.issuer

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secretKey))
}

func (j *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secretKey), nil
	})
}

func (j *jwtService) FindUserIDByToken(token string) (uuid.UUID, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secretKey), nil
	})
	if err != nil {
		return uuid.Nil, err
	}

	userID, err := uuid.Parse(claims["user_id"].(string))
	if err != nil {
		return uuid.Nil, err
	}

	return userID, nil
}
