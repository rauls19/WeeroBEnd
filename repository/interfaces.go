package repository

import (
	model "weeroBE/model"
)

//IUserRepository is the interface of the User repository
type IUserRepository interface {
	LoginUser(model.User) string
	GetUser(model.User) model.User
	RemoveUser()
	CreateUser(model.User) bool
	ModifyProfile(model.User)
	UploadPhoto(model.User)
	RemovePhoto(model.User)
	//UpdateField(field []string, value []string, key string)
	InsertPathImage(id string, imageid string)
}

//IAuthRepository is the interface of the Token repository
type IAuthRepository interface {
	RenewToken()
	SaveClaim(clientid string, clientsecret string, scope string, userid int)
	RemoveClaim(userid int)
	RevokeToken()
	CheckClaim(clientid string, clientsecrect string) bool
	CheckClientId(clientid string) (bool, string)
}
