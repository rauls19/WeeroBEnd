package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"weeroBE/api/model"
	"weeroBE/api/repository"
	"weeroBE/api/utils"
)

type AuthI struct {
	token repository.IAuthRepository
}

//InstanceAuth is
func InstanceAuth(instance repository.IAuthRepository) *AuthI {
	return &AuthI{instance}
}

//GetToken is get token
func (instance *AuthI) GetToken(w http.ResponseWriter, r *http.Request) {
	data := json.NewDecoder(r.Body)
	var autht model.Token
	err := data.Decode(&autht)
	if err != nil && autht.Grand_type != "" {
		panic(err)
	}
	tk := &model.Token{}
	check, claims, err := utils.TokenIsValid(r)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	if !check {
		json.NewEncoder(w).Encode("No token valid")
		return
	}
	checked, pss := instance.token.CheckClientId(claims[4])
	if !checked {
		json.NewEncoder(w).Encode("No valid")
		return
	}
	tk.Access_token = utils.GenerateToken(claims[4], pss)
	//tk.Grand_type = utils.GenerateRefreshToken(strconv.Itoa(autht.CliendId))
	tk.TokenType = "Bearer"
	json.NewEncoder(w).Encode(tk)
}
