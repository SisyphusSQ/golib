package sql_dto

type Version struct {
	Version string `json:"version" gorm:"column:@@version"`
}

type TableInfo struct {
	Table       string `json:"TABLE" gorm:"column:TABLE"`
	CreateTable string `json:"Create Table" gorm:"column:Create Table"`
}

type IndexInfo struct {
	Table        string `json:"Table" gorm:"column:Table"`
	NonUnique    int    `json:"Non_unique" gorm:"column:Non_unique"`
	KeyName      string `json:"Key_name" gorm:"column:Key_name"`
	SeqInIndex   int    `json:"Seq_in_index" gorm:"column:Seq_in_index"`
	ColumnName   string `json:"Column_name" gorm:"column:Column_name"`
	Collation    string `json:"Collation" gorm:"column:Collation"`
	Cardinality  int    `json:"Cardinality" gorm:"column:Cardinality"`
	SubPart      int    `json:"Sub_part" gorm:"column:Sub_part"`
	Packed       int    `json:"Packed" gorm:"column:Packed"`
	Null         string `json:"Null" gorm:"column:Null"`
	IndexType    string `json:"Index_type" gorm:"column:Index_type"`
	Comment      string `json:"Comment" gorm:"column:Comment"`
	IndexComment string `json:"Index_comment" gorm:"column:Index_comment"`
}

type ExplainInfo struct {
	Id           int     `json:"id" gorm:"column:id"`
	SelectType   string  `json:"select_type" gorm:"column:select_type"`
	Table        string  `json:"table" gorm:"column:table"`
	Partitions   string  `json:"partitions" gorm:"column:partitions"`
	Type         string  `json:"type" gorm:"column:type"`
	PossibleKeys string  `json:"possible_keys" gorm:"column:possible_keys"`
	Key          string  `json:"key" gorm:"column:key"`
	KeyLen       string  `json:"key_len" gorm:"column:key_len"`
	Ref          string  `json:"ref" gorm:"column:ref"`
	Rows         int     `json:"rows" gorm:"column:rows"`
	Filtered     float64 `json:"filtered" gorm:"column:filtered"`
	Extra        string  `json:"Extra" gorm:"column:Extra"`
}

type BadRecord struct {
	ID                int64  `gorm:"primaryKey;column:id" json:"id"`
	Digest            string `gorm:"column:digest" json:"digest"`
	ClientUser        string `gorm:"column:client_user" json:"client_user"`
	DB                string `gorm:"column:db" json:"db"`
	Cmd               string `gorm:"column:cmd" json:"cmd"`
	SQLExample        string `gorm:"column:sql_example" json:"sql_example"`
	NoIndexUsed       int8   `gorm:"column:no_index_used" json:"no_index_used"`
	NoGoodIndexUsed   int8   `gorm:"column:no_good_index_used" json:"no_good_index_used"`
	SortScan          int8   `gorm:"column:sort_scan" json:"sort_scan"`
	AvgExecTime       int64  `gorm:"column:avg_exec_time" json:"avg_exec_time"`
	MaxExecTime       int64  `gorm:"column:max_exec_time" json:"max_exec_time"`
	AvgSentRows       int64  `gorm:"column:avg_sent_rows" json:"avg_sent_rows"`
	MaxSentRows       int64  `gorm:"column:max_sent_rows" json:"max_sent_rows"`
	AvgRowsExamined   int64  `gorm:"column:avg_rows_examined" json:"avg_rows_examined"`
	MaxRowsExamined   int64  `gorm:"column:max_rows_examined" json:"max_rows_examined"`
	AvgAffectedRows   int64  `gorm:"column:avg_affected_rows" json:"avg_affected_rows"`
	MaxAffectedRows   int64  `gorm:"column:max_affected_rows" json:"max_affected_rows"`
	AvgLockTime       int64  `gorm:"column:avg_lock_time" json:"avg_lock_time"`
	MaxLockTime       int64  `gorm:"column:max_lock_time" json:"max_lock_time"`
	AvgSortRows       int64  `gorm:"column:avg_sort_rows" json:"avg_sort_rows"`
	MaxSortRows       int64  `gorm:"column:max_sort_rows" json:"max_sort_rows"`
	ExecTimeCount     int64  `gorm:"-" json:"exec_time_count"`
	SentRowsCount     int64  `gorm:"-" json:"sent_rows_count"`
	RowsExaminedCount int64  `gorm:"-" json:"rows_examined_count"`
	AffectedRowsCount int64  `gorm:"-" json:"affected_rows_count"`
	LockTimeCount     int64  `gorm:"-" json:"lock_time_count"`
	SortRowsCount     int64  `gorm:"-" json:"sort_rows_count"`
}
