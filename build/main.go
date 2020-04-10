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

	fmt.Println("tes")

	var userProfile model.UserProfile
	var userProfile2 model.UserProfile

	var opt model.UserProfileContractor

	userProfile.Create("michaelchandrag", "password")
	userProfile2.Create("michaelchandrag2", "password")
	
	var _userProfile model.UserProfileInterface
	var _userProfile2 model.UserProfileInterface
	_userProfile = userProfile
	_userProfile2 = userProfile2

	opt.AddContractor(_userProfile)
	opt.AddContractor(_userProfile2)

	fmt.Println(opt.GetUserProfiles())
}