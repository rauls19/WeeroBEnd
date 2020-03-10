package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"weeroBE/api/config"
	"weeroBE/api/handler"
	"weeroBE/api/repository"
)

//HandlerDependency is
type HandlerDependency struct {
	userDIH  *handler.UserI
	tokenDIH *handler.AuthI
}

func HomeEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world :)")
}

func router(in HandlerDependency) {
	http.HandleFunc("/", HomeEndpoint)
	http.HandleFunc("/getNToken", in.tokenDIH.GetToken)
	http.HandleFunc("/SignUp", in.userDIH.SignUp)
	http.HandleFunc("/getUser", in.userDIH.GetUser)
	http.HandleFunc("/loginUser", in.userDIH.LoginUser)
	http.HandleFunc("/updateInformation", in.userDIH.UpdateField)
	http.HandleFunc("/UpdatePhoto", in.userDIH.UpdatePhoto)
}

//InjectionDependency is
func InjectionDependency(db *sql.DB) HandlerDependency {
	InstanceDb := repository.InstanceDB(db)
	handlerdependency := HandlerDependency{
		userDIH:  handler.InstanceUser(InstanceDb, InstanceDb),
		tokenDIH: handler.InstanceAuth(InstanceDb),
	}
	return handlerdependency
}

func main() {
	db, err := config.ConnectionDb()
	if err != nil {
		fmt.Println("Error Connexion")
	}
	dependencies := InjectionDependency(db)
	router(dependencies)
	fmt.Println("Running")
	http.ListenAndServe(":3000", nil)
}
