package service

import (
	"io"
	"mime/multipart"
	"os"
	"strconv"
	"strings"
	"time"
	"weeroBE/model"
)

func getAge(date string) int {
	currenttime := time.Now()
	birthday := strings.Split(date, "-")
	year, _ := strconv.Atoi(birthday[1])
	month, _ := strconv.Atoi(birthday[2])
	dia, _ := strconv.Atoi(birthday[3])
	if month >= int(currenttime.Month()) && dia > currenttime.Day() {
		return (currenttime.Year() - year - 1)
	}
	return (currenttime.Year() - year)
}

//FieldToUpdate return the field and the value, 1 by 1
//Deute tècnic treure els ifs, posar map and goroutine and treure hardcode
func FieldToUpdate(userMod model.User) ([]string, []string, bool) {
	var fields []string
	var values []string
	lang := false
	nupdates := 0

	if userMod.Name != "" {
		nupdates = nupdates + 1
		fields[nupdates] = "name"
		values[nupdates] = userMod.Name
	}
	if userMod.Surname != "" {
		nupdates = nupdates + 1
		fields[nupdates] = "surname"
		values[nupdates] = userMod.Surname
	}
	if userMod.Birthday != "" {
		nupdates = nupdates + 1
		fields[nupdates] = "birthday"
		values[nupdates] = userMod.Birthday
		nupdates = nupdates + 1
		fields[nupdates] = "age"
		values[nupdates] = strconv.Itoa(getAge(userMod.Birthday))
	}
	if userMod.Email != "" {
		nupdates = nupdates + 1
		fields[nupdates] = "email"
		values[nupdates] = userMod.Email
	}
	if userMod.Interested != 0 {
		nupdates = nupdates + 1
		fields[nupdates] = "interested"
		values[nupdates] = strconv.Itoa(userMod.Interested)
	}
	if userMod.Description != "" {
		nupdates = nupdates + 1
		fields[nupdates] = "description"
		values[nupdates] = userMod.Description
	}
	if userMod.Languages != nil || len(userMod.Languages) > 0 {
		nupdates = nupdates + 1
		fields[nupdates] = "language"
		values = append(values, userMod.Languages...) //Extrany
		lang = true
	}
	return fields, values, lang
}

func PrepareUserInformation(data model.User) {

}

//SaveImage Comprovar que passa si ja està creat???
func SaveImage(multp *multipart.Form, userid string, imageid string) {
	files := multp.File["data"]
	for i := range files {
		//for each fileheader, get a handle to the actual file
		file, err := files[i].Open()
		defer file.Close()
		if err != nil {
			return
		}
		//create destination file making sure the path is writeable.
		dst, err := os.Create("C:/Go/src/weeroBE/images/old/" + userid + "_" + imageid + ".jpg")
		defer dst.Close()
		if err != nil {
			return
		}
		//copy the uploaded file to the destination file
		if _, err := io.Copy(dst, file); err != nil {
			return
		}
	}
}
