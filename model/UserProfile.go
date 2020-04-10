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

func (this UserProfile) Update(id int) error {
	query := fmt.Sprintf(`
			UPDATE 
				user_profile
			SET
				email = ?,
				password = ?,
				address = ?
			WHERE
				user_id = ?
		`)
		_, err := db.Engine.Exec(query,
				this.Email, this.Password, this.Address,
				id)
		if err != nil {
			fmt.Println(err)
			return err
		}
		return nil
}

func (this UserProfile) Delete(id int) error {
	query := fmt.Sprintf(`
			DELETE FROM
				user_profile
			WHERE
				user_id = ?
		`)
	_, err := db.Engine.Exec(query,
				id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (this UserProfile) Find() (results []UserProfile, err error) {
	query := `
			SELECT
				user_id,
				email,
				password,
				address
			FROM
				user_profile
		`
	err = db.Engine.Select(&results, query)
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (this UserProfile) FindByEmail(email string) (result UserProfile, err error) {
	query := `
		SELECT
			user_id,
			email,
			password,
			address
		FROM
			user_profile
		WHERE
			email = ?
	`
	if err := db.Engine.Get(&this, query, email); err != nil {
		return this, err
	}
	return this, nil
}