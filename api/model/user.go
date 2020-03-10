package model

type User struct {
	ID          int `json:"-"`
	Hid         string
	Name        string
	Surname     string
	Birthday    string
	Email       string
	Password    string
	Age         int
	Interested  int
	Location    string
	Description string
	Mobilephone int
	Languages   []string
}

type ProfilePicture struct {
	ID         int
	UserId     int
	Base64Data string
}
