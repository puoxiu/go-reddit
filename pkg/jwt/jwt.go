package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)


const (
	TokenExpireDuration = time.Hour * 2
	SecretKey = "xingxingisme"
	JWTIssuer = "reddit"
)

type MyClaims struct {
	UserID int64 `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// GenToken 生成JWT
func GetToken(userID int64, username string) (string, error) {
	// 创建一个我们自己的声明
	c := MyClaims{
		userID,
		username,
		jwt.StandardClaims{
			// 过期时间
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer: JWTIssuer,
		},
	}

	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString([]byte(SecretKey))
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