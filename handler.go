package main

import "net/http"

var jwtKey = []byte("secret_key")

var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

func Login(w http.ResponseWriter, r *http.Request) {

}

func Signup(w http.ResponseWriter, r *http.Request) {

}

func Home(w http.ResponseWriter, r *http.Request) {

}

func Refresh(w http.ResponseWriter, r *http.Request) {

}
