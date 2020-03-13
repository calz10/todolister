package models

import "github.com/lib/pq"

type Profile struct {
	UserId    string `"type:varchar(100)[]"`
	Name      string
	Birthday  int64
	TodoIds   pq.StringArray `gorm:"type:varchar(64)[]"`
	GroupdIds pq.StringArray `gorm:"type:varchar(64)[]"`
}
