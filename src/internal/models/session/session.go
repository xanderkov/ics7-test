package session

import (
	"context"
)

type Session struct {
	SessionID string
	UserId    int
}

type sessionCtx struct{}

func GetSessionFromCtx(ctx context.Context) (session Session, ok bool) {
	session, ok = ctx.Value(sessionCtx{}).(Session)
	return session, ok
}

func SetSessionToCtx(ctx context.Context, session Session) context.Context {
	return context.WithValue(ctx, sessionCtx{}, session)
}
