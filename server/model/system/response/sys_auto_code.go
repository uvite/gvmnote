package response

import "time"

type Db struct {
	Database string `json:"database" gorm:"column:database"`
}

type Table struct {
	TableName  string    `json:"table_name" gorm:"column:table_name"`
	TableComment    string    `json:"table_comment" gorm:"column:table_comment"`
	CreateTime time.Time `json:"create_time" gorm:"column:create_time"`
}

type Column struct {
	DataType      string `json:"dataType" gorm:"column:data_type"`
	ColumnName    string `json:"columnName" gorm:"column:column_name"`
	DataTypeLong  string `json:"dataTypeLong" gorm:"column:data_type_long"`
	ColumnComment string `json:"columnComment" gorm:"column:column_comment"`
}
