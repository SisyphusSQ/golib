package base_do

import "time"

type ConfigKV struct {
	ID        int64     `gorm:"primaryKey;column:id" json:"id"`
	K         string    `gorm:"column:config_key" json:"configKey"`
	V         string    `gorm:"column:config_value" json:"configValue"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (ConfigKV) TableName() string {
	return "common_config_kv"
}
