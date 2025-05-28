package base_do

import "time"

type LarkMsgLog struct {
	ID          int64     `gorm:"primaryKey;column:id" json:"id"`
	UUID        string    `gorm:"column:uuid" json:"uuid"`
	MessageID   string    `gorm:"column:message_id" json:"message_id"`
	Contacts    string    `gorm:"column:contacts" json:"contacts"`
	Status      Status    `gorm:"column:send_status" json:"send_status"`
	FromMethod  string    `gorm:"column:from_method" json:"from_method"`
	SendMessage string    `gorm:"column:send_message" json:"send_message"`
	Response    string    `gorm:"column:response" json:"response"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (LarkMsgLog) TableName() string {
	return "lark_msg_log"
}

type Status int8

const (
	Success Status = iota + 1
	Failure
)
