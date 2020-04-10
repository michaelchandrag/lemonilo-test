package main

import (
	"os"
	"fmt"

	model "github.com/michaelchandrag/lemonilo-test/model"
	database "github.com/michaelchandrag/lemonilo-test/database"
)

func main () {
	err := database.Connect()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}
	fmt.Println("Database connected")

	// sample insert logic
	/* userProfile := model.UserProfile{
		Email: "canzinzzzide@yahoo.co.id",
		Password: "password",
		Address: "address",
	}

	var opt model.UserProfileContractor
	
	var _userProfile model.UserProfileInterface
	_userProfile = userProfile

	result, err := opt.InsertContractor(_userProfile)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result) */

	// sample update logic
	userID := 1
	userProfile := model.UserProfile{
		Email: "canzinzzzide@yahoo.co.id",
		Password: "newPassword",
		Address: "newAddress",
	}

	var opt model.UserProfileContractor

	var _userProfile model.UserProfileInterface = userProfile

	err = opt.UpdateContractor(userID, _userProfile)
	if err != nil {
		fmt.Println(err)
	}
	return


}