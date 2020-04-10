package model

type UserProfile struct {
	Username 	string
	Password 	string
}

func (this *UserProfile) Create(username string, password string) {
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
}