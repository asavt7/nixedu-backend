package service

import (
	"errors"
	"github.com/asavt7/nixEducation/pkg/model"
	"github.com/asavt7/nixEducation/pkg/storage"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

const (
	jwtSecretKey        = "GET_ME_FROM_ENV"   //todo
	jwtRefreshSecretKey = "GET_ME_FROM_ENV_1" //todo
)

func GetRefreshJWTSecret() string {
	return jwtRefreshSecretKey
}

func GetJWTSecret() string {
	return jwtSecretKey
}

type Claims struct {
	UserId int `json:"userId"`
	jwt.StandardClaims
}

type AuthorizationServiceImpl struct {
	repo storage.UserStorage
}

func (s *AuthorizationServiceImpl) CheckUserCredentials(username string, password string) (model.User, error) {
	user, err := s.repo.GetByUsername(username)
	if err != nil {
		return user, err
	}
	err = checkPassword(user, password)
	return user, err
}

func (s *AuthorizationServiceImpl) IsNeedToRefresh(claims Claims) bool {
	return time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) < 15*time.Minute
}

func checkPassword(user model.User, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
}

func (s *AuthorizationServiceImpl) ParseTokenToClaims(token string) (Claims, error) {
	tkn, err := jwt.ParseWithClaims(token, Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(GetRefreshJWTSecret()), nil
	})
	if err != nil {
		return Claims{}, err
	}
	if tkn == nil || !tkn.Valid {
		return Claims{}, errors.New("Token is incorrect")
	}
	return tkn.Claims.(Claims), nil
}

func (s *AuthorizationServiceImpl) GenerateTokens(userId int) (accessToken, refreshToken string, accessExp, refreshExp time.Time, err error) {
	accessToken, accessExp, err = generateToken(userId, time.Now().Add(1*time.Hour), []byte(GetJWTSecret()))
	refreshToken, refreshExp, err = generateToken(userId, time.Now().Add(24*time.Hour), []byte(GetRefreshJWTSecret()))
	return
}

func generateToken(userId int, expirationTime time.Time, secret []byte) (string, time.Time, error) {
	claims := &Claims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", time.Now(), err
	}

	return tokenString, expirationTime, nil
}

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
