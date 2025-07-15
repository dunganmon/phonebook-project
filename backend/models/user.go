package models

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	FullName string `json:"full_name"`
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
}

func CreateUser(Username string, Password string, FullName string) (newUser User) {
	return User{
		FullName: FullName,
		Username: Username,
		Password: Password,
	}
}
