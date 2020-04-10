package model

import (
	"fmt"

	db "github.com/michaelchandrag/lemonilo-test/database"
)

type UserProfile struct {
	UserID 		int 		`json:"user_id" db:"user_id"`
	Email 		string 		`json:"email" db:"email"`
	Password 	string 		`json:"password" db:"password"`
	Address 	string 		`json:"address" db:"address"`
}

/* func (this *UserProfile) Create(username string, password string) {
	this.Username = username
	this.Password = password
}

func (this UserProfile) GetUsername() string {
	return this.Username
}

func (this *UserProfile) GetPassword() string {
	return this.Password
}

func (this *UserProfile) GetContents() UserProfile {
	return UserProfile{
		Username: this.GetUsername(),
		Password: this.GetPassword(),
	}
} */

func (this UserProfile) Insert() (result UserProfile, err error) {
	query := fmt.Sprintf(`
			INSERT into user_profile (
				email, address, password
			) VALUES (
				?, ?, ?
			)
		`)

	resp, err := db.Engine.Exec(query,
				this.Email, this.Address, this.Password)
	if err != nil {
		fmt.Println(err)
		return result, err
	}
	lastId, _ := resp.LastInsertId()
	result = UserProfile{
		UserID: int(lastId),
		Email: this.Email,
		Address: this.Address,
		Password: this.Password,
	}
	return result, nil
}