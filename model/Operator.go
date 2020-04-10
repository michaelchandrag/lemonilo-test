package model

import (
	"fmt"
)

type UserProfileContractor struct {
	profile UserProfileInterface
}

/* func (us *UserProfileContractor) AddContractor(upi UserProfileInterface) {
	us.profiles = append(us.profiles, upi)
} */

/* func (us *UserProfileContractor) GetUserProfiles() UserProfileResult {
	result := make(UserProfileResult, len(us.profiles))
	for i := 0; i < len(us.profiles); i++ {
		result[i] = us.profiles[i].GetUsername()
	}
	return result
} */

func (us *UserProfileContractor) InsertContractor(upi UserProfileInterface) (result UserProfile, err error){
	result, err = upi.Insert()
	if err != nil {
		fmt.Println(err)
		return result, err
	}
	return result, err
}

func (us *UserProfileContractor) UpdateContractor(id int, upi UserProfileInterface) (error) {
	err := upi.Update(id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (us *UserProfileContractor) DeleteContractor(id int, upi UserProfileInterface) (error) {
	err := upi.Delete(id)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}