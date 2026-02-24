package store

import (
	"context"
	"encoding/json"
	"fmt"

	repository "github.com/knr1997/assets-management-apiserver/internal/utils"
	"gorm.io/gorm"
)

type AuditService interface {
	LogCreate(ctx context.Context, tx *gorm.DB, table string, id any, obj any) error
	LogUpdate(ctx context.Context, tx *gorm.DB, table string, id any, oldObj, newObj any) error
	LogDelete(ctx context.Context, tx *gorm.DB, table string, id any, obj any) error
}

type auditService struct {
	repo AuditRepository
}

func NewAuditService(repo AuditRepository) AuditService {
	return &auditService{repo: repo}
}

func extractAuditContext(ctx context.Context) AuditContext {
	if val, ok := GetAuditContext(ctx); ok {
		return val
	}
	return AuditContext{}
}

func (a *auditService) LogCreate(
	ctx context.Context,
	tx *gorm.DB,
	table string,
	id any,
	obj any,
) error {

	newValue, _ := json.Marshal(obj)
	auditCtx := extractAuditContext(ctx)

	return a.repo.Log(ctx, tx, AuditEntry{
		TableName: table,
		RecordID:  fmt.Sprintf("%v", id),
		Operation: "CREATE",
		NewValue:  newValue,

		ChangedBy: auditCtx.UserID,
		IPAddress: &auditCtx.IPAddress,
		RequestID: &auditCtx.RequestID,
		UserAgent: &auditCtx.UserAgent,
	})
}

func (a *auditService) LogUpdate(
	ctx context.Context,
	tx *gorm.DB,
	table string,
	id any,
	oldObj, newObj any,
) error {

	oldValue, _ := json.Marshal(oldObj)
	newValue, _ := json.Marshal(newObj)

	diff := repository.CalculateDiff(oldValue, newValue)

	return a.repo.Log(ctx, tx, AuditEntry{
		TableName: table,
		RecordID:  fmt.Sprintf("%v", id),
		Operation: "UPDATE",
		OldValue:  oldValue,
		NewValue:  newValue,
		Diff:      json.RawMessage(diff),
	})
}

func (a *auditService) LogDelete(
	ctx context.Context,
	tx *gorm.DB,
	table string,
	id any,
	obj any,
) error {

	oldValue, _ := json.Marshal(obj)

	return a.repo.Log(ctx, tx, AuditEntry{
		TableName: table,
		RecordID:  fmt.Sprintf("%v", id),
		Operation: "DELETE",
		OldValue:  oldValue,
	})
}
