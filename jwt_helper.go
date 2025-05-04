package main

import (
    "errors"
    "github.com/golang-jwt/jwt/v4"
)

var jwtSecret = []byte("my-secret-key")

func ValidateToken(tokenString string) (*JWTPayload, error) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, errors.New("unexpected signing method")
        }
        return jwtSecret, nil
    })

    if err != nil {
        return nil, err
    }

    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        return &JWTPayload{
            MerchantId: claims["merchantId"].(string),
            BranchId:   claims["branchId"].(string),
            DeviceId:   claims["deviceId"].(string),
            BranchSk:   claims["branchSk"].(string),
            LoginRefId: claims["loginRefId"].(string),
        }, nil
    } else {
        return nil, errors.New("invalid token claims")
    }
}
