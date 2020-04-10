package model


type UserProfileInterface interface {
	// GetUsername() 	string
	Insert() (UserProfile, error)
}