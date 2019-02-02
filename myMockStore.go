package testing_handlers

import "fmt"

type MyMockStore struct {

}

// The concept here is that we can mock the behavior of our data store.
// Since we are primarily testing our handlers, we can assume any interaction
// with the data store behaves normally. The only way to dictate that is to
// mock the behavior to our liking. See below
func (ms *MyMockStore) GetByID(id int64) (*User, error) {
	// We can trigger an error by passing in an id of 2
	if id == 2 {
		return nil, fmt.Errorf("Error getting user with id: %d", id)
	}
	user := &User{
		ID:          1,
		Email:       "Anything you want",
		Password:    "Anything you want",
		FirstName:   "Anything you want",
		LastName:    "Anything you want",
		Description: "Anything you want",
	}
	
	return user, nil
}


func (ms *MyMockStore) InsertNewUser(user *User) (*User, error){
	if user.FirstName == "Error" {
		return nil, fmt.Errorf("Error Inserting New User")
	}
	
	// Assumes that if user without the FirstName field being equal to "Error" will
	// always result in a successful insert
	return user, nil
}


