package validateViaDataBase

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"github.com/neerajsrivastav7/SSCPYTHON/SScBackEnd/userManageMent/client"
	"golang.org/x/crypto/bcrypt"
)

// ValidatePassFromDB is a struct that holds the user database client instance
type ValidatePassFromDB struct {
	clientIns *client.Client
}

// NewValidatePassFromDB initializes ValidatePassFromDB with a client instance
func NewValidatePassFromDB(configFile string) (*ValidatePassFromDB, error) {
	clientIns, err := client.NewClient(configFile)
	if err != nil {
		return nil, err
	}
	return &ValidatePassFromDB{
		clientIns: clientIns,
	}, nil
}

// UserDetails represents the structure of user data expected from the API response
type UserDetails struct {
	PasswordHash string `json:"password_hash"` // Adjust this field based on actual response
}

// ValidateViaPassword validates a username and password against the database
func (validatePass *ValidatePassFromDB) ValidateViaPassword(username, password string) (interface{}, error) {
	// Retrieve the user details from the database
	resp, err := validatePass.clientIns.getUserByUserName(username)
	if err != nil {
		return nil, errors.New("user not found")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("error fetching user details")
	}

	// Read and parse the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("error reading response body")
	}

	var userDetails UserDetails
	if err := json.Unmarshal(body, &userDetails); err != nil {
		return nil, errors.New("error parsing user details")
	}

	// Compare the provided password with the hashed password from the database
	if err := bcrypt.CompareHashAndPassword([]byte(userDetails.PasswordHash), []byte(password)); err != nil {
		return nil, errors.New("invalid password")
	}

	// If the password is correct, return the user details (or any required information)
	return userDetails, nil
}
