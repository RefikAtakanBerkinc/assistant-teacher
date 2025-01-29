package main

import (
	"net/http"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	var payload = struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Version string `json:"version"`
	}{
		Status:  "active",
		Message: "Assistant teacher up and running",
		Version: "1.0.0",
	}

	_ = app.wrtieJSON(w, http.StatusOK, payload)
}

func (app *application) allStutendts(w http.ResponseWriter, r *http.Request) {
	students, err := app.DB.AllStudents()
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	_ = app.wrtieJSON(w, http.StatusOK, students)
}
