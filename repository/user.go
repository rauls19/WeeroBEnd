package repository

import (
	"fmt"
	"strconv"
	"weeroBE/model"
)

func (db *Connection) CreateUser(data model.User) bool {
	_, err := db.db.Query(`INSERT INTO users(password, mobilephone, hashId) VALUES (?,?,?)`,
		data.Password, data.Mobilephone, data.Hid)
	if err != nil {
		fmt.Println("Error: ", err)
		return false
	}
	return true
}

//LoginUser is not junk but return Model or All info or Token
func (db *Connection) LoginUser(data model.User) string {
	var hashid string
	err := db.db.QueryRow("Select hashid from users where mobilephone= ? and password= ? and hashid = ?", data.Mobilephone, data.Password, data.Hid).Scan(&hashid)
	if err != nil {
		fmt.Println("Error loginUser ", err)
	}
	return hashid
}

//GetUser mirar que retorna
func (db *Connection) GetUser(data model.User) model.User {
	var model model.User
	err := db.db.QueryRow("Select id, name, surname, email, age, interested, location, description, mobilephone from users where mobilephone= ?",
		data.Hid).Scan(&model.ID, &model.Name,
		&model.Surname, &model.Email, &model.Age, &model.Interested,
		&model.Location, &model.Description, &model.Mobilephone)
	if err != nil {
		fmt.Println(err)
	}
	result, err := db.db.Query("Select language from languages where userid=?", model.ID)
	for result.Next() {
		var lang string
		if err := result.Scan(&lang); err != nil {
			fmt.Printf("Error ", err)
		}
		model.Languages = append(model.Languages, lang)
	}
	return model
}

func (db *Connection) RemoveUser() {

}

//ModifyProfile is the function ...
func (db *Connection) ModifyProfile(data model.User) {

}

//UploadPhoto is the function ...
func (db *Connection) UploadPhoto(data model.User) {

}

//RemovePhoto is the function ...
func (db *Connection) RemovePhoto(data model.User) {

}

//Updatefield is the function ...
func (db *Connection) UpdateField(field string, value string, key string) {
	var id int
	var query string
	err := db.db.QueryRow(`SELECT id FROM users WHERE hashid = ?`, key).Scan(&id)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	if _, err := strconv.Atoi(value); err == nil {
		query = "UPDATE users SET " + field + " = " + value + " WHERE id = " + strconv.Itoa(id)
	} else {
		query = "UPDATE users SET " + field + " = \"" + value + "\" WHERE id = " + strconv.Itoa(id)
	}
	_, err = db.db.Query(query)
	if err != nil {
		fmt.Println("Error: ", err)
	}
}
func (db *Connection) InsertPathImage(id string, imageid string) {
	var count int
	key, _ := strconv.Atoi(id)
	imagekey, _ := strconv.Atoi(imageid)
	path := "C:/Go/src/weeroBE/images/old/"
	err := db.db.QueryRow(`SELECT count(*) from pictures where userid=? and imageid = ?`, key, imagekey).Scan(&count)
	if err != nil {
		fmt.Println(err)
	}
	if count == 1 {
		_, err = db.db.Query(`UPTDATE pictures SET userid=?, path=? WHERE userid = ?`, key, path, key)
	}
	_, err = db.db.Query(`INSERT INTO pictures (userid,imageid,path) VALUES (?,?,?)`, key, imagekey, path)
	if err != nil {
		fmt.Println(err)
	}
}
