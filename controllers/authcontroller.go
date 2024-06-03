package controllers

import (
	"encoding/json"
	"first-jwt/configs"
	"first-jwt/dto"
	"first-jwt/helpers"
	"first-jwt/models"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	// var register contains Register model
	var register dto.Register
	// NewDecoder for collect the body, and Decode function for decode the body into register shape / mapping into (Register Model)
	if err := json.NewDecoder(r.Body).Decode(&register); err != nil{ // Decode function returns error (err)
		helpers.Response(w, 500, err.Error(), nil)
		return
	}
	defer r.Body.Close() // will execute at the end of the function

	if register.Password != register.PasswordConfirm {
		helpers.Response(w, 500, "Confirm Password not Match", nil)
		return
	}
	// Hashing the password (bcrypting)
	PasswordHash, err := helpers.HashPassword(register.Password)
	if err != nil {
		helpers.Response(w, 400, err.Error(), nil)
		return
	}

	// Mapping the register data to real User Model Struct
	user := models.User {
		Name: register.Name,
		Email: register.Email,
		Password: PasswordHash,
	}

	// Create function do create data to database, and Error to catch the error
	if err := configs.DB.Create(&user).Error; err != nil{
		helpers.Response(w, 500, err.Error(), nil)
		return
	}
	helpers.Response(w, 201, "Register Successfully", nil)
}



func Login(w http.ResponseWriter, r *http.Request) {
	// var register contains Register model
	var login dto.Login

	// NewDecoder for collect the body, and Decode function for decode the body into register shape / mapping into (Register Model)
	if err := json.NewDecoder(r.Body).Decode(&login); err != nil { // Decode function returns error (err)
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	var user models.User
	// find the data based of the query "email = ?"
	if err := configs.DB.First(&user, "email = ?", login.Email).Error; err != nil {
		helpers.Response(w, 404, "Wrong Email or Password", nil)
		return
	}

	// verifying password
	if err := helpers.VerifyPassword(user.Password, login.Password); err != nil {
		helpers.Response(w, 404, "Wrong Email or Password", nil)
		return
	}

	// creating token 
	token, err := helpers.CreateToken(&user)
	if err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}
	// end response
	helpers.Response(w, 200, "SuccessFuly Login", token)
}