package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/mohitdhaundiyal/go-auth/internal/config"
	"github.com/mohitdhaundiyal/go-auth/internal/models"
	"github.com/mohitdhaundiyal/go-auth/internal/utils"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello JWT")
}

func Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.User

	// decode data from body
	json.NewDecoder(r.Body).Decode(&user)

	// check if user already exists or not
	res := config.DB.Find(&user, "email = ?", user.Email)

	if res.RowsAffected == 1 {
		json.NewEncoder(w).Encode("user already exists with this email id.")
		return
	}

	password, err := utils.GeneratehashPassword(user.Password)

	if err != nil {
		log.Fatal(err.Error())
	} else {
		user.Password = password
	}

	// create new record
	config.DB.Save(&user)
	json.NewEncoder(w).Encode(&user) // new user is registered.
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user models.User
	var auth models.Authentication

	// decode data from body
	json.NewDecoder(r.Body).Decode(&auth)

	// check email exists or not
	res := config.DB.Find(&user, "email = ?", auth.Email)

	if res.RowsAffected != 1 {
		json.NewEncoder(w).Encode("Username or Password is Incorrect.")
		return
	}

	// if email exits -> check password matches or not.
	passwordMatch := utils.CheckPasswordHash(auth.Password, user.Password)

	if !passwordMatch {
		json.NewEncoder(w).Encode("Username or Password is Incorrect.")
		return
	}

	// if password matches :
	validToken, err := utils.GenerateJWT(user.Email, strconv.Itoa(int(user.ID)))

	if err != nil {
		log.Fatal(err.Error())
	}

	var token models.Token
	token.Email = auth.Email
	token.TokenString = validToken
	token.ID = strconv.Itoa(int(user.ID))

	json.NewEncoder(w).Encode(token.TokenString)
}
