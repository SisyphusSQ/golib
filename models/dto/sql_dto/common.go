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
	KeyLen       int     `json:"key_len" gorm:"column:key_len"`
	Ref          string  `json:"ref" gorm:"column:ref"`
	Rows         int     `json:"rows" gorm:"column:rows"`
	Filtered     float64 `json:"filtered" gorm:"column:filtered"`
	Extra        string  `json:"Extra" gorm:"column:Extra"`
}
