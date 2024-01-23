package models

type User struct {
	Id          int64  `gorm:"primaryKey" json:"id"`
	NamaLengkap string `gorm:"varchar(300)" json:"nama_lengkap"`
	Email       string `gorm:"varchar(300)" json:"email"`
	Password    string `gorm:"varchar(300)" json:"password"`
}
