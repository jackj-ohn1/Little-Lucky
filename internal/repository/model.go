package repository

type User struct {
	Account string `gorm:"account,primaryKey"`
	Email   string `gorm:"email"`
}
