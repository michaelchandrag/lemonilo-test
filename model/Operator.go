package model

type UserProfileResult []string

type UserProfileContractor struct {
	profiles []UserProfileInterface
}

func (us *UserProfileContractor) AddContractor(upi UserProfileInterface) {
	us.profiles = append(us.profiles, upi)
}

func (us *UserProfileContractor) GetUserProfiles() UserProfileResult {
	result := make(UserProfileResult, len(us.profiles))
	for i := 0; i < len(us.profiles); i++ {
		result[i] = us.profiles[i].GetUsername()
	}
	return result
}