package repository

import (
	"fmt"
)

func (db *Connection) RenewToken() {

}
func (db *Connection) SaveClaim(clientid string, clientsecrect string, scope string, userid int) {
	_, err := db.db.Query(`INSERT INTO claims(clientid, clientpassword, scope, userid) VALUES (?,?,?,?)`, clientid, clientsecrect, scope, userid)
	if err != nil {
		fmt.Println("Error: ", err)
	}
}
func (db *Connection) RemoveClaim(userid int) {
	_, err := db.db.Query(`DELETE FROM claims WHERE userid=?`, userid)
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

//Revisar, no em convenç count ==1 True
func (db *Connection) CheckClaim(clientid string, clientsecrect string) bool {
	var count int
	err := db.db.QueryRow(`SELECT count(*) FROM claims WHERE clientid=? and clientpassword=?`, clientid, clientsecrect).Scan(&count)
	if err != nil && (count > 1 || count == 0) {
		fmt.Println("Error: ", err)
		return false
	}
	return true
}
func (db *Connection) CheckClientId(clientid string) (bool, string) {
	count := 0
	var pss string
	rows, err := db.db.Query(`SELECT clientpassword FROM claims WHERE clientid=?`, clientid)
	for rows.Next() {
		count++
		rows.Scan(&pss)
	}
	if err != nil && (count > 1 || count == 0) {
		fmt.Println("Error: ", err)
		return false, ""
	}
	return true, pss
}
func (db *Connection) RevokeToken() {

}
