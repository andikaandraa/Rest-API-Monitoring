package controllers

import (
	"errors"
	"net/http"
	"project/helper"
	"project/services"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// HandleUserRequest handle user requests
func HandleUserRequest(w http.ResponseWriter, r *http.Request, Db *gorm.DB) error {
	if r.URL.RequestURI() != "/user" {
		helper.ResponseCallback(w, http.StatusNotFound, map[string]string{
			"message": "Page Not Found",
			"code":    strconv.Itoa(http.StatusNotFound),
		})
		err := errors.New("message: Page Not Found")
		return err
	}
	// log.Println("Access /user endpoint")
	switch r.Method {
	case "POST":
		services.HandleAddUser(w, r, Db)
	case "PUT":
		services.HandleUpdateUser(w, r, Db)
	case "DELETE":
		services.HandleDeleteUser(w, r, Db)
	default:
		helper.ResponseCallback(w, http.StatusMethodNotAllowed, map[string]string{
			"message": "Invalid request method",
			"code":    strconv.Itoa(http.StatusMethodNotAllowed),
		})
		err := errors.New("message: Invalid request method")
		return err
	}

	return nil
}

// HandleUsersRequest handle get all users request
func HandleUsersRequest(w http.ResponseWriter, r *http.Request, Db *gorm.DB) error {
	if r.URL.RequestURI() != "/users" {
		helper.ResponseCallback(w, http.StatusNotFound, map[string]string{
			"message": "Page Not Found",
			"code":    strconv.Itoa(http.StatusNotFound),
		})
		err := errors.New("message: Page Not Found")
		return err
	}
	// log.Println("Access /users endpoint")
	switch r.Method {
	case "GET":
		services.HandleGetUsers(w, r, Db)
	default:
		helper.ResponseCallback(w, http.StatusMethodNotAllowed, map[string]string{
			"message": "Invalid request method",
			"code":    strconv.Itoa(http.StatusMethodNotAllowed),
		})
		err := errors.New("message: Invalid request method")
		return err
	}

	return nil
}

// HandleGetUserRequest handle get all users request
func HandleGetUserRequest(w http.ResponseWriter, r *http.Request, Db *gorm.DB) error {
	var id int

	vars := mux.Vars(r)
	ids, err := strconv.Atoi(vars["id"])
	if err != nil {
		helper.ResponseCallback(w, http.StatusNotFound, map[string]string{
			"message": "Bad request",
			"code":    strconv.Itoa(http.StatusBadRequest),
		})
		return err
	}
	id = ids

	// log.Printf("Access /user/%v endpoint\n", id)
	switch r.Method {
	case "GET":
		err := services.HandleGetUser(w, r, id, Db)
		return err
	default:
		helper.ResponseCallback(w, http.StatusMethodNotAllowed, map[string]string{
			"message": "Invalid request method",
			"code":    strconv.Itoa(http.StatusMethodNotAllowed),
		})
		err := errors.New("message: Invalid request method")
		return err
	}
}
