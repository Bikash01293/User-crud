package controllers

import (
	"encoding/json"
	"net/http"
	"crud/services"
)

func Create(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		service.CreateUser(w, r)
	default:
		json.NewEncoder(w).Encode("Bad request !")
	}
}

func Get(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		service.GetUser(w, r)
	default:
		json.NewEncoder(w).Encode("Bad Request !")
	}
}

func Update(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "PUT":
		service.UpdateUser(w, r)
	default:
		json.NewEncoder(w).Encode("Bad Request !")
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "DELETE":
		service.DeleteUser(w, r)
	default:
		json.NewEncoder(w).Encode("Bad Request !")
	}
}