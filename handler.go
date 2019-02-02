package testing_handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"strconv"
)

type Context struct {
	UserStore MyMockStore
}

/*
	WARNING: The code written here should be used to understand the concept of testing
	http handlers and mocking data store behaviors. DO NOT expect the code here to be consider
	"good" code. :)
 */

// Simple handler that decodes the request body for a user and adds it to
// some data store.
func (ctx *Context) InsertNewUserHandler(w http.ResponseWriter, r *http.Request) {
	user := &User{}
	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		http.Error(w, fmt.Sprintf("error decoding JSON: %v", err),
			http.StatusBadRequest)
		return
	}
	
	user, err := ctx.UserStore.InsertNewUser(user)
	if err != nil {
		// This example isn't checking what type of error our UserStore returns
		// You may need to do this to differentiate from a duplicate entry vs.
		// an error from the database
		http.Error(w, fmt.Sprintf("error inserting new user: %v", err),
			http.StatusInternalServerError)
		return
	}
	
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, fmt.Sprintf("Error encoding JSON: %v", err),
			http.StatusInternalServerError)
		return
	}
}


func (ctx *Context) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	id := path.Base(r.URL.Path)
	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Enable to convert string to int"),
			http.StatusBadRequest)
		return
	}
	
	user, err := ctx.UserStore.GetByID(int64(idInt))
	if err != nil {
		http.Error(w, fmt.Sprintf("Error getting user with id: %v", err),
			http.StatusInternalServerError)
		return
	}
	
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, fmt.Sprintf("Error encoding JSON: %v", err),
			http.StatusInternalServerError)
		return
	}
}