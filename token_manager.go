package main

import (
    "context"
)

type contextKey string

const payloadKey = contextKey("jwtPayload")

func StorePayload(ctx context.Context, payload *JWTPayload) context.Context {
    return context.WithValue(ctx, payloadKey, payload)
}

func GetPayload(ctx context.Context) *JWTPayload {
    if payload, ok := ctx.Value(payloadKey).(*JWTPayload); ok {
        return payload
    }
    return nil
}
