package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"weeroBE/model"
	"weeroBE/repository"
	"weeroBE/service"
	"weeroBE/utils"
)

//UserI variable with all instances
type UserI struct {
	userrepository repository.IUserRepository
	authrepository repository.IAuthRepository
}

//InstanceUser return the instances of the User
//Params: instance User, instance Authentication
//Interface added
func InstanceUser(instance repository.IUserRepository, authinstance repository.IAuthRepository) *UserI {
	return &UserI{instance, authinstance}
}

//SignUp returns a new Authentication Model
//Params: mobilephone, password & HashId
//Generate the clientID, clientSecret & Scope and save it in DB
//Returns: Authentication Model, False (if the user exists in the DB)
func (instance *UserI) SignUp(w http.ResponseWriter, r *http.Request) {
	data := json.NewDecoder(r.Body)
	var userMod model.User
	//var response model.Response
	err := data.Decode(&userMod)
	if err != nil {
		panic(err)
	}
	clientid := utils.GenerateClientID(userMod.Mobilephone)
	result := instance.userrepository.CreateUser(userMod)
	if result {
		clientSecret := utils.GenerateCLientSecret(userMod.Mobilephone, userMod.Password)
		instance.authrepository.SaveClaim(clientid, clientSecret, "Test Scope", userMod.Mobilephone)
		tk := &model.Token{}
		tk.Access_token = utils.GenerateToken(clientid, clientSecret)
		tk.Grand_type = utils.GenerateRefreshToken(clientid)
		tk.TokenType = "Bearer"
		//response.Response.Response{tk}
		json.NewEncoder(w).Encode(tk)
		return
	}
	json.NewEncoder(w).Encode(result)
}

//LoginUser returns a new Authentication Model
//Params: mobilephone, password & HashId
//Generate the clientID, clientSecret & Scope and save it in DB
//Returns: Authentication Model, Empty (if the user doesn't exist in the DB)
func (instance *UserI) LoginUser(w http.ResponseWriter, r *http.Request) {
	data := json.NewDecoder(r.Body)
	var userMod model.User
	err := data.Decode(&userMod)
	if err != nil {
		panic(err)
	}
	//userMod.Age = service.GetAge(userMod.Birthday)
	result := instance.userrepository.LoginUser(userMod)
	if result != "" {
		clientid := utils.GenerateClientID(userMod.Mobilephone)
		clientSecret := utils.GenerateCLientSecret(userMod.Mobilephone, userMod.Password)
		instance.authrepository.RemoveClaim(userMod.Mobilephone)
		instance.authrepository.SaveClaim(clientid, clientSecret, "Scope Login", userMod.Mobilephone)
		tk := &model.Token{}
		tk.Access_token = utils.GenerateToken(clientid, clientSecret)
		tk.Grand_type = utils.GenerateRefreshToken(clientid)
		tk.TokenType = "Bearer"
		json.NewEncoder(w).Encode(tk)
		return
	}
	json.NewEncoder(w).Encode(result)
}

//GetUser returns info User
//Params: HashId & Valid Token
//Generate the clientID, clientSecret & Scope and save it in DB
//Returns: User Model, String with the error (if the user doesn't exist in the DB, or the token is not valid)
func (instance *UserI) GetUser(w http.ResponseWriter, r *http.Request) {
	data := json.NewDecoder(r.Body)
	check, claims, err := utils.TokenIsValid(r)
	if err != nil {
		fmt.Println("Error: ", err)
		json.NewEncoder(w).Encode("Error")
	}
	if !check {
		json.NewEncoder(w).Encode("No token valid")
		return
	}
	checked := instance.authrepository.CheckClaim(claims[1], claims[2])
	if !checked {
		json.NewEncoder(w).Encode("No valid")
		return
	}
	var userMod model.User
	err = data.Decode(&userMod)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	resp := instance.userrepository.GetUser(userMod)
	json.NewEncoder(w).Encode(resp)
}

//UpdateField is the function to update some field of the profile ESTÀ PENDENT DE COM S'ACTUALITZA DES DE LA APP
func (instance *UserI) UpdateField(w http.ResponseWriter, r *http.Request) {
	data := json.NewDecoder(r.Body)
	check, claims, err := utils.TokenIsValid(r)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	if !check {
		json.NewEncoder(w).Encode("No token valid")

		return
	}
	checked := instance.authrepository.CheckClaim(claims[1], claims[2])
	if !checked {
		json.NewEncoder(w).Encode("No valid")
		return
	}
	var userMod model.User
	err = data.Decode(&userMod)
	if err != nil && userMod.Hid == "" {
		fmt.Println("Error: ", err)
	}
	//Business Logic
	//field, value, calct := service.FieldToUpdate(userMod)
	//Repository
	//instance.userrepository.UpdateField(field, value, userMod.Hid)
}

//UpdatePhoto is the function to update...
func (instance *UserI) UpdatePhoto(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(5 << 20)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	m := r.MultipartForm
	service.SaveImage(m, r.FormValue("userid"), r.FormValue("imageid"))
	instance.userrepository.InsertPathImage(r.FormValue("userid"), r.FormValue("imageid"))
}
