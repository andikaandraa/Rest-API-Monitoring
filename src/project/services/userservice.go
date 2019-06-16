package services

import (
	"encoding/json"
	"errors"
	"net/http"
	"project/helper"
	"project/models"
	"strconv"

	"github.com/jinzhu/gorm"
	// "gopkg.in/go-playground/validator.v9"
)

// User represent user model
type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

// HandleAddUser handle add user request
func HandleAddUser(w http.ResponseWriter, r *http.Request, Db *gorm.DB) {
	var (
		user    models.User
		code    int
		message interface{}
	)

	decoder := json.NewDecoder(r.Body)
	errDec := decoder.Decode(&user)
	if errDec != nil {
		code = http.StatusInternalServerError
		message = map[string]string{
			"message": errDec.Error(),
			"code":    strconv.Itoa(http.StatusInternalServerError),
		}
		helper.ResponseCallback(w, code, message)
		return
	}

	err := Db.Create(&user).Error
	if err != nil {
		code = http.StatusBadRequest
		message = map[string]string{
			"message": err.Error(),
			"code":    strconv.Itoa(http.StatusBadRequest),
		}
	} else {
		code = http.StatusOK
		message = map[string]string{
			"message": "Success",
			"code":    strconv.Itoa(http.StatusOK),
		}
	}

	helper.ResponseCallback(w, code, message)
}

// HandleUpdateUser handle add user request
func HandleUpdateUser(w http.ResponseWriter, r *http.Request, Db *gorm.DB) {

	var (
		user      User
		userModel models.User
		newUser   models.User
		id        string
		code      int
		message   interface{}
	)

	decoder := json.NewDecoder(r.Body)
	errDec := decoder.Decode(&user)
	if errDec != nil {
		code = http.StatusInternalServerError
		message = map[string]string{
			"message": errDec.Error(),
			"code":    strconv.Itoa(http.StatusInternalServerError),
		}
		helper.ResponseCallback(w, code, message)
		return
	}

	id = user.ID
	err := Db.Where("id =?", id).First(&userModel).Error
	if err != nil {
		code = http.StatusNoContent
		message = map[string]string{
			"message": err.Error(),
			"code":    strconv.Itoa(http.StatusNoContent),
		}
		helper.ResponseCallback(w, code, message)
		return
	}

	newUser.Name = user.Name
	newUser.Email = user.Email
	newUser.Password = user.Password

	err = Db.Model(&userModel).Updates(newUser).Error

	if err != nil {
		code = http.StatusBadRequest
		message = map[string]string{
			"message": err.Error(),
			"code":    strconv.Itoa(http.StatusBadRequest),
		}
	} else {
		code = http.StatusOK
		message = map[string]string{
			"message": "Success",
			"code":    strconv.Itoa(http.StatusOK),
		}
	}

	helper.ResponseCallback(w, code, message)
}

// HandleDeleteUser handle add user request
func HandleDeleteUser(w http.ResponseWriter, r *http.Request, Db *gorm.DB) {
	var (
		user      User
		userModel models.User
		id        string
		code      int
		message   interface{}
	)

	decoder := json.NewDecoder(r.Body)
	errDec := decoder.Decode(&user)
	if errDec != nil {
		code = http.StatusInternalServerError
		message = map[string]string{
			"message": errDec.Error(),
			"code":    strconv.Itoa(http.StatusInternalServerError),
		}
		helper.ResponseCallback(w, code, message)
		return
	}

	id = user.ID
	err := Db.First(&userModel, id).Error
	if err != nil {
		helper.ResponseCallback(w, http.StatusNoContent, map[string]string{
			"message": err.Error(),
			"code":    strconv.Itoa(http.StatusNoContent),
		})
		return
	}

	err = Db.Delete(&userModel).Error

	if err != nil {
		code = http.StatusBadRequest
		message = map[string]string{
			"message": "Bad Request",
			"code":    strconv.Itoa(http.StatusBadRequest),
		}
	} else {
		code = http.StatusOK
		message = map[string]string{
			"message": "Success",
			"code":    strconv.Itoa(http.StatusOK),
		}
	}

	helper.ResponseCallback(w, code, message)
}

// HandleGetUsers handle get all users
func HandleGetUsers(w http.ResponseWriter, r *http.Request, Db *gorm.DB) {

	// db := config.DBInit()
	// objDB := &ObjDB{DB: db}

	var (
		users   []models.User
		code    int
		message interface{}
	)

	err := Db.Find(&users).Error
	if err != nil {
		code = http.StatusBadRequest
		message = map[string]string{
			"message": err.Error(),
			"code":    strconv.Itoa(http.StatusBadRequest),
		}
	} else {
		code = http.StatusOK
		message = map[string][]models.User{
			"result": users,
		}
	}

	helper.ResponseCallback(w, code, message)
}

// HandleGetUser handle get all users
func HandleGetUser(w http.ResponseWriter, r *http.Request, id int, Db *gorm.DB) error {

	var (
		user    models.User
		code    int
		message interface{}
	)

	err := Db.Where("id =?", id).First(&user).Error
	if err != nil {
		code = http.StatusBadRequest
		message = map[string]string{
			"message": err.Error(),
			"code":    strconv.Itoa(http.StatusBadRequest),
		}

		helper.ResponseCallback(w, code, message)

		err := errors.New("message: Invalid request method")
		return err
	}

	code = http.StatusOK
	message = map[string]models.User{
		"result": user,
	}
	helper.ResponseCallback(w, code, message)

	return nil
}
