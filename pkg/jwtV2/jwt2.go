package jwtV2

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)


const (
	TokenExpireDuration = time.Hour * 2
	TokenExpireDurationLong = time.Hour * 24 * 30
	SecretKey = "xingxingisme"
	JWTIssuer = "reddit"
)

type MyClaims struct {
	UserID int64 `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// GenToken 生成JWT
func GetToken(userID int64, username string) (accessToken string,refreshToken string, err error) {
	// 创建声明 访问token
	accessClaims := MyClaims{
		userID,
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer: JWTIssuer,
		},
	}
	// 创建声明 刷新token
	refreshClaims := MyClaims{
		userID,
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDurationLong).Unix(),
			Issuer: JWTIssuer,
		},
	}

	accessToken,err = jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims).SignedString([]byte(SecretKey))
	if err != nil {
		return "", "", err
	}
	refreshToken,err = jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SecretKey))
	if err != nil {
		return "", "", err
	}

	return accessToken,refreshToken, nil
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, err
}
