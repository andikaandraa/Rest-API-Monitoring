package controllers

import (
	"net/http"
	"project/helper"
	"strconv"
)

// HomepageHandler handle / request
func HomepageHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.RequestURI() != "/" {
		helper.ResponseCallback(w, http.StatusNotFound, map[string]string{
			"message": "Page Not Found",
			"code":    strconv.Itoa(http.StatusNotFound),
		})
		return
	}

	if r.Method == "GET" {
		helper.ResponseCallback(w, http.StatusOK, map[string]string{
			"message": "Success",
			"code":    strconv.Itoa(http.StatusOK),
		})
	} else {
		helper.ResponseCallback(w, http.StatusMethodNotAllowed, map[string]string{
			"message": "Invalid request method",
			"code":    strconv.Itoa(http.StatusMethodNotAllowed),
		})
	}
}
