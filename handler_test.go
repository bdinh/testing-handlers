package testing_handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// Test is pretty redundant, you don't need to refactor it
func TestContext_GetUserHandler(t *testing.T) {
	context := &Context{
		UserStore: MyMockStore{},
	}
	
	validUser := User{
		ID:          1,
		Email:       "Anything you want",
		Password:    "Anything you want",
		FirstName:   "Anything you want",
		LastName:    "Anything you want",
		Description: "Anything you want",
	}
	
	cases := []struct {
		name string
		method string
		idPath string
		expectedStatusCode int
		expectedError bool
		expectedContentType string
		expectedReturn *User
	}{
		{
			"Valid Get Request",
			http.MethodGet,
			"1",
			http.StatusOK,
			false,
			"application/json",
			&validUser,
		},
		{
			"Invalid Get Request",
			http.MethodGet,
			"2",
			http.StatusInternalServerError,
			true,
			"text/plain; charset=utf-8",
			nil,
		},
	}
	
	for _, c := range cases {
		request := httptest.NewRequest(c.method, "/v1/users/" + c.idPath, nil)
		recorder := httptest.NewRecorder()
		context.GetUserHandler(recorder, request)
		response := recorder.Result()
		
		resContentType := response.Header.Get("Content-Type")
		if !c.expectedError && c.expectedContentType != resContentType {
			t.Errorf("case %s: incorrect return type: expected: %s recieved: %s",
				c.name, c.expectedContentType, resContentType)
		}
		
		resStatusCode := response.StatusCode
		if c.expectedStatusCode != resStatusCode {
			t.Errorf("case %s: incorrect status code: expected: %d recieved: %d",
				c.name, c.expectedStatusCode, resStatusCode)
		}
		
		user := &User{}
		err := json.NewDecoder(response.Body).Decode(user)
		if c.expectedError && err == nil {
			t.Errorf("case %s: expected error but revieved none", c.name)
		}
		
		if !c.expectedError && c.expectedReturn.Email!= user.Email && c.expectedReturn.Password != user.Password &&
			c.expectedReturn.FirstName != user.FirstName && c.expectedReturn.LastName != user.LastName {
			t.Errorf("case %s: incorrect return: expected %v but revieved %v",
				c.name, c.expectedReturn, user)
		}
	}
	
	
}

func TestContext_InsertNewUserHandler(t *testing.T) {
	context := &Context{
		UserStore: MyMockStore{},
	}
	
	validUser := User{
		ID:          1,
		Email:       "Anything you want",
		Password:    "Anything you want",
		FirstName:   "Valid User",
		LastName:    "Anything you want",
		Description: "Anything you want",
	}
	
	invalidUser := User{
		ID:          1,
		Email:       "Anything you want",
		Password:    "Anything you want",
		FirstName:   "Error",
		LastName:    "Anything you want",
		Description: "Anything you want",
	}
	
	cases := []struct {
		name string
		requestBody User
		method string
		expectedStatusCode int
		expectedError bool
		expectedContentType string
		expectedReturn *User
	}{
		{
			"Valid Post Request",
			validUser,
			http.MethodPost,
			http.StatusCreated,
			false,
			"application/json",
			&validUser,
		},
		{
			"Invalid Post Request",
			invalidUser,
			http.MethodPost,
			http.StatusInternalServerError,
			true,
			"text/plain; charset=utf-8",
			nil,
		},
	}
	
	
	for _, c := range cases {
		body, _ := json.Marshal(c.requestBody)
		request := httptest.NewRequest(c.method, "/v1/users", strings.NewReader(string(body)))
		recorder := httptest.NewRecorder()
		context.InsertNewUserHandler(recorder, request)
		response := recorder.Result()
		
		resContentType := response.Header.Get("Content-Type")
		if !c.expectedError && c.expectedContentType != resContentType {
			t.Errorf("case %s: incorrect return type: expected: %s recieved: %s",
				c.name, c.expectedContentType, resContentType)
		}
		
		resStatusCode := response.StatusCode
		if c.expectedStatusCode != resStatusCode {
			t.Errorf("case %s: incorrect status code: expected: %d recieved: %d",
				c.name, c.expectedStatusCode, resStatusCode)
		}
		
		user := &User{}
		err := json.NewDecoder(response.Body).Decode(user)
		if c.expectedError && err == nil {
			t.Errorf("case %s: expected error but revieved none", c.name)
		}
		
		if !c.expectedError && c.expectedReturn.Email!= user.Email && c.expectedReturn.Password != user.Password &&
			c.expectedReturn.FirstName != user.FirstName && c.expectedReturn.LastName != user.LastName {
			t.Errorf("case %s: incorrect return: expected %v but revieved %v",
				c.name, c.expectedReturn, user)
		}
	}
}
