package model

type UserModel struct {
	ID             uint64 `gorm:"column:id;primaryKey"`
	Username       string `gorm:"column:username;type:varchar(64);unique"`
	PasswordSha256 string `gorm:"column:password;type:varchar(128)"`
}
