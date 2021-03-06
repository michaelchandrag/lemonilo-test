package model

import (
	"fmt"
)

type UserProfileContractor struct {
	profile UserProfileInterface
}

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
		return err
	}
	return nil
}

func (us *UserProfileContractor) FindContractor(upi UserProfileInterface) (results []UserProfile, err error) {
	results, err = upi.Find()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return results, err
}

func (us *UserProfileContractor) LoginContractor(email string, upi UserProfileInterface) (result UserProfile, err error) {
	result, err = upi.FindByEmail(email)
	if err != nil {
		fmt.Println(err)
		return result, err
	}
	return result, nil
}