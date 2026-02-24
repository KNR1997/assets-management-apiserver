package store

import (
	"context"
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type AuditLog struct {
	ID int64 `gorm:"primaryKey"`

	// What was changed
	TableName string `gorm:"column:table_name;type:varchar(100);not null"`
	RecordID  string `gorm:"column:record_id;type:varchar(100);not null"`

	// What changed
	FieldName *string `gorm:"column:field_name;type:varchar(100)"`
	OldValue  *string `gorm:"column:old_value;type:text"`
	NewValue  *string `gorm:"column:new_value;type:text"`
	Diff      *string `gorm:"column:diff;type:text"`

	// Who and when
	Operation string    `gorm:"column:operation;type:varchar(10);not null"`
	ChangedAt time.Time `gorm:"column:changed_at;type:timestamptz;not null;default:now()"`
	ChangedBy string    `gorm:"column:changed_by;type:varchar(100);not null"`

	// Context
	IPAddress *string `gorm:"column:ip_address;type:inet"`
	UserAgent *string `gorm:"column:user_agent;type:text"`
	SessionID *string `gorm:"column:session_id;type:varchar(100)"`
	RequestID *string `gorm:"column:request_id;type:varchar(100)"`

	// Additional metadata
	// Metadata datatypes.JSON `gorm:"column:metadata;type:jsonb"`
}

type AuditEntry struct {
	TableName string
	RecordID  string
	Operation string

	FieldName *string
	OldValue  json.RawMessage
	NewValue  json.RawMessage
	Diff      json.RawMessage

	ChangedBy string
	IPAddress *string
	UserAgent *string
	SessionID *string
	RequestID *string

	Metadata json.RawMessage
}

type AuditRepository interface {
	Log(ctx context.Context, tx *gorm.DB, entry AuditEntry) error
}

type auditRepository struct {
}

func NewAuditRepository() AuditRepository {
	return &auditRepository{}
}

func (r *auditRepository) Log(
	ctx context.Context,
	tx *gorm.DB,
	entry AuditEntry,
) error {

	log := AuditLog{
		TableName: entry.TableName,
		RecordID:  entry.RecordID,
		Operation: entry.Operation,

		FieldName: entry.FieldName,
		ChangedAt: time.Now(),
		ChangedBy: entry.ChangedBy,

		IPAddress: entry.IPAddress,
		UserAgent: entry.UserAgent,
		SessionID: entry.SessionID,
		RequestID: entry.RequestID,
	}

	if entry.OldValue != nil {
		old := string(entry.OldValue)
		log.OldValue = &old
	}

	if entry.NewValue != nil {
		newVal := string(entry.NewValue)
		log.NewValue = &newVal
	}

	if entry.Diff != nil {
		diff := string(entry.Diff)
		log.Diff = &diff
	}

	// if entry.Metadata != nil {
	// 	log.Metadata = datatypes.JSON(entry.Metadata)
	// }

	return tx.WithContext(ctx).Create(&log).Error
}
