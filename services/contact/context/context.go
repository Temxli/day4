package context

import (
    "context"
    "github.com/google/uuid"
)

type ContextKey string

const (
    ContextKeyID ContextKey = "id"
)

func WithID(ctx context.Context, id string) context.Context {
    return context.WithValue(ctx, ContextKeyID, id)
}

func GetID(ctx context.Context) string {
    id, _ := ctx.Value(ContextKeyID).(string)
    return id
}
