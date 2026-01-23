package auth

import "context"

type CheckUserFunc func(ctx context.Context, phone string) (string, error)
type CheckUserPasswordFunc func(ctx context.Context, phone, password string) (string, error)
