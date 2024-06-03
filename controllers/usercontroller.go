package controllers

import (
	"first-jwt/dto"
	"first-jwt/helpers"
	"net/http"
)

func Me(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("userinfo").(*helpers.CustomClaims) // convert into a struct
	userResponse := &dto.Profile{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
	}

helpers.Response(w, 200, "My Profile", userResponse)
}