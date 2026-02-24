package store

import "context"

type AuditContext struct {
	UserID    string
	IPAddress string
	RequestID string
	UserAgent string
}

type auditCtxKey struct{}

var auditContextKey = auditCtxKey{}

func SetAuditContext(ctx context.Context, val AuditContext) context.Context {
	return context.WithValue(ctx, auditContextKey, val)
}

func GetAuditContext(ctx context.Context) (AuditContext, bool) {
	val, ok := ctx.Value(auditContextKey).(AuditContext)
	return val, ok
}
