package entity

type Users struct {
	UserName string `json:"user_name" gorm:"primaryKey; type:varchar(50); not null"`
	Password string `json:"password" gorm:"type:varchar(100); not null"`
}
