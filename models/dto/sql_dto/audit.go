package sql_dto

type AuditSQLRecord struct {
	Cluster         string   `gorm:"column:cluster" json:"cluster"`
	ClusterVendor   string   `gorm:"column:cluster_vendor" json:"cluster_vendor"`
	Digest          string   `gorm:"column:digest" json:"digest"`
	DB              string   `gorm:"column:db" json:"db"`
	Cmd             string   `gorm:"column:cmd" json:"cmd"`
	SQLExample      string   `gorm:"column:sql_example" json:"sql_example"`
	NoIndexUsed     int8     `gorm:"column:no_index_used" json:"no_index_used"`
	NoGoodIndexUsed int8     `gorm:"column:no_good_index_used" json:"no_good_index_used"`
	SortScan        int8     `gorm:"column:sort_scan" json:"sort_scan"`
	AvgExecTime     int64    `gorm:"column:avg_exec_time" json:"avg_exec_time"`
	MaxExecTime     int64    `gorm:"column:max_exec_time" json:"max_exec_time"`
	AvgSentRows     int64    `gorm:"column:avg_sent_rows" json:"avg_sent_rows"`
	MaxSentRows     int64    `gorm:"column:max_sent_rows" json:"max_sent_rows"`
	AvgRowsExamined int64    `gorm:"column:avg_rows_examined" json:"avg_rows_examined"`
	MaxRowsExamined int64    `gorm:"column:max_rows_examined" json:"max_rows_examined"`
	AvgAffectedRows int64    `gorm:"column:avg_affected_rows" json:"avg_affected_rows"`
	MaxAffectedRows int64    `gorm:"column:max_affected_rows" json:"max_affected_rows"`
	AvgLockTime     int64    `gorm:"column:avg_lock_time" json:"avg_lock_time"`
	MaxLockTime     int64    `gorm:"column:max_lock_time" json:"max_lock_time"`
	AvgSortRows     int64    `gorm:"column:avg_sort_rows" json:"avg_sort_rows"`
	MaxSortRows     int64    `gorm:"column:max_sort_rows" json:"max_sort_rows"`
	BadReason       []string `gorm:"column:bad_reason" json:"bad_reason"`
	IsCut           bool     `gorm:"column:is_cut" json:"is_cut"`
}
