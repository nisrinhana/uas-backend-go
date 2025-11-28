package helper

import (
    "time"
    "github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("SECRET_KEY")

type CustomClaims struct {
    UserID      string   `json:"user_id"`
    RoleID      string   `json:"role_id"`
    Permissions []string `json:"permissions"`
    jwt.RegisteredClaims
}

func GenerateJWT(userID string, roleID string, perms []string) (string, error) {

    claims := CustomClaims{
        UserID:      userID,
        RoleID:      roleID,
        Permissions: perms,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtSecret)
}

func VerifyJWT(tokenStr string) (*CustomClaims, error) {

    token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
        return jwtSecret, nil
    })
    if err != nil {
        return nil, err
    }

    claims, ok := token.Claims.(*CustomClaims)
    if !ok || !token.Valid {
        return nil, jwt.ErrSignatureInvalid
    }

    return claims, nil
}
