package service

import (
	"errors"
	"github.com/asavt7/nixedu/backend/internal/model"
	"github.com/asavt7/nixedu/backend/internal/storage"
	"github.com/asavt7/nixedu/backend/internal/tokenstorage"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

const (
	jwtSecretKey        = "GET_ME_FROM_ENV"   //todo
	jwtRefreshSecretKey = "GET_ME_FROM_ENV_1" //todo
	refreshTime         = 15 * time.Minute
	accessTokenTTL      = 1 * time.Hour
	refreshTokenTTL     = 12 * time.Hour
)

// GetRefreshJWTSecret returns secret for refresh tokens
func GetRefreshJWTSecret() string {
	return jwtRefreshSecretKey
}

// GetJWTSecret returns secret for access tokens
func GetJWTSecret() string {
	return jwtSecretKey
}

type Claims struct {
	UserID int    `json:"userId"`
	UID    string `json:"uid"`
	jwt.StandardClaims
}

type AuthorizationServiceImpl struct {
	repo       storage.UserStorage
	tokenStore tokenstorage.TokenKeeper
}

func (s *AuthorizationServiceImpl) Logout(accessTokenClaims *Claims) error {
	userID := accessTokenClaims.UserID
	return s.tokenStore.Delete(userID)
}

func (s *AuthorizationServiceImpl) ValidateRefreshToken(accessTokenClaims *Claims) error {
	return s.validateToken(accessTokenClaims, func(tokens model.CachedTokens) string {
		return tokens.RefreshUID
	})
}

func (s *AuthorizationServiceImpl) ValidateAccessToken(accessTokenClaims *Claims) error {
	return s.validateToken(accessTokenClaims, func(tokens model.CachedTokens) string {
		return tokens.AccessUID
	})
}

func (s *AuthorizationServiceImpl) validateToken(accessTokenClaims *Claims, tokenFromCashedFunc func(model.CachedTokens) string) error {
	userID := accessTokenClaims.UserID
	clientTokenUID := accessTokenClaims.UID

	cached, err := s.tokenStore.Get(userID)
	if err != nil {
		return err
	}
	if tokenFromCashedFunc(cached) != clientTokenUID {
		return errors.New("invalid token")
	}
	return nil
}

// CheckUserCredentials
func (s *AuthorizationServiceImpl) CheckUserCredentials(username string, password string) (model.User, error) {
	user, err := s.repo.FindByUsername(username)
	if err != nil {
		return user, err
	}
	err = checkPassword(user.PasswordHash, password)
	return user, err
}

func (s *AuthorizationServiceImpl) IsNeedToRefresh(claims *Claims) bool {
	return time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) < refreshTime
}

func checkPassword(passwordHash string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
}

func (s *AuthorizationServiceImpl) ParseAccessTokenToClaims(token string) (*Claims, error) {
	return s.parseTokenToClaims(token, []byte(GetJWTSecret()))
}

func (s *AuthorizationServiceImpl) ParseRefreshTokenToClaims(token string) (*Claims, error) {
	return s.parseTokenToClaims(token, []byte(GetRefreshJWTSecret()))
}

func (s *AuthorizationServiceImpl) parseTokenToClaims(token string, key []byte) (*Claims, error) {
	tkn, err := jwt.ParseWithClaims(token, Claims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return &Claims{}, err
	}
	if tkn == nil || !tkn.Valid {
		return &Claims{}, errors.New("Token is incorrect")
	}
	return tkn.Claims.(*Claims), nil
}

// GenerateTokens GenerateTokens
func (s *AuthorizationServiceImpl) GenerateTokens(userID int) (accessToken, refreshToken string, accessExp, refreshExp time.Time, err error) {
	accessToken, accessUID, accessExp, err := generateToken(userID, time.Now().Add(accessTokenTTL), []byte(GetJWTSecret()))
	refreshToken, refreshUID, refreshExp, err := generateToken(userID, time.Now().Add(refreshTokenTTL), []byte(GetRefreshJWTSecret()))

	if err != nil {
		return
	}

	cashedTokens := model.CachedTokens{
		AccessUID:  accessUID,
		RefreshUID: refreshUID,
	}
	_, err = s.tokenStore.Save(userID, cashedTokens)

	return
}

func generateToken(userID int, expirationTime time.Time, secret []byte) (string, string, time.Time, error) {
	tokenUID := uuid.New().String()
	claims := &Claims{
		UserID: userID,
		UID:    tokenUID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", tokenUID, time.Now(), err
	}

	return tokenString, tokenUID, expirationTime, nil
}

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
