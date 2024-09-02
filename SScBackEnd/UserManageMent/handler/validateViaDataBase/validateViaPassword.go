// validateViaDataBase.go
package validateViaDataBase

import (
	"errors"
	"github.com/neerajsrivastav7/SSCPYTHON/SScBackEnd/DataBase/dataBaseUserManageMent"
	"golang.org/x/crypto/bcrypt"
)

// ValidatePassFromDB is a struct that holds the user database
type ValidatePassFromDB struct {
	userDatabase dataBaseUserManageMent.User
}

// NewValidatePassFromDB initializes a new instance of ValidatePassFromDB
func NewValidatePassFromDB() *ValidatePassFromDB {
	return &ValidatePassFromDB{
		userDatabase:  dataBaseUserManageMent.User{},
	}
}

// ValidateViaPassword validates a username and password against the database
func (validatePass *ValidatePassFromDB) ValidateViaPassword(username, password string) (interface{}, error) {
	// Retrieve the user details from the database
	userDetails, err := validatePass.userDatabase.GetUserByUsername(username)
	if err != nil {
		return nil, errors.New("user not found")
	}

	// Assuming the userDetails has a field 'PasswordHash' which stores the hashed password
	if err := bcrypt.CompareHashAndPassword([]byte(userDetails.PasswordHash), []byte(password)); err != nil {
		return nil, errors.New("invalid password")
	}

	// If the password is correct, return the user details (or any required information)
	return userDetails, nil
}
