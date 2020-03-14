package model

type Auth struct {
	Token    int
	user     string
	password string
}

type Token struct {
	Access_token string
	Grand_type   string
	TokenType    string
}

type AuthClient struct {
	CliendId     int
	ClientSecret string
	Scope        string
}
