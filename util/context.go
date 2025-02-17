package util

import (
	"context"
)

type contextKey string

const userIDContextKey contextKey = "userID"

func SetUserIDToContext(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, userIDContextKey, userID)
}

func GetUserIDFromContext(ctx context.Context) string {
	userID := ctx.Value(userIDContextKey).(string)
	return userID
}
