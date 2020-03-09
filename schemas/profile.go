package schemas

type Profile struct {
	UserId    string
	TodoIds   []string `gorm:"type:varchar(64)[]"`
	GroupdIds []string `gorm:"type:varchar(64)[]"`
}
