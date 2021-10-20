package service

import (
	"errors"
	"github.com/asavt7/nixEducation/pkg/model"
	"github.com/asavt7/nixEducation/pkg/storage"
	"github.com/asavt7/nixEducation/pkg/tokenstorage"
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

	salt = "some_salt_from_configs" //todo
)

func GetRefreshJWTSecret() string {
	return jwtRefreshSecretKey
}

func GetJWTSecret() string {
	return jwtSecretKey
}

type Claims struct {
	UserID int    `json:"userId"`
	Uuid   string `json:"uuid"`
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
	userId := accessTokenClaims.UserID
	clientTokenUuid := accessTokenClaims.Uuid

	cached, err := s.tokenStore.Get(userId)
	if err != nil {
		return err
	}
	if tokenFromCashedFunc(cached) != clientTokenUuid {
		return errors.New("invalid token")
	}
	return nil
}

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

func (s *AuthorizationServiceImpl) GenerateTokens(userID int) (accessToken, refreshToken string, accessExp, refreshExp time.Time, err error) {
	accessToken, accessUuid, accessExp, err := generateToken(userID, time.Now().Add(accessTokenTTL), []byte(GetJWTSecret()))
	refreshToken, refreshUuid, refreshExp, err := generateToken(userID, time.Now().Add(refreshTokenTTL), []byte(GetRefreshJWTSecret()))

	if err != nil {
		return
	}

	cashedTokens := model.CachedTokens{
		AccessUID:  accessUuid,
		RefreshUID: refreshUuid,
	}
	_, err = s.tokenStore.Save(userID, cashedTokens)

	return
}

func generateToken(userID int, expirationTime time.Time, secret []byte) (string, string, time.Time, error) {
	tokenUuid := uuid.New().String()
	claims := &Claims{
		UserID: userID,
		Uuid:   tokenUuid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", tokenUuid, time.Now(), err
	}

	return tokenString, tokenUuid, expirationTime, nil
}

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
