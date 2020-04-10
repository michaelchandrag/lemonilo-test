package model


type UserProfileInterface interface {
	Insert() (UserProfile, error)
	Update(int) (error)
	Delete(int) (error)
}