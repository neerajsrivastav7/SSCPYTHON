package handler
import (
	"fmt"
	"github.com/neerajsrivastav7/SSCPYTHON/SScBackEnd/UserManageMent/handler/validateViaDataBase"
)

type Validate struct {
}

func (validate *Validate) ValidateLogin(userName string, password string, loginType string) (interface{}, error) {
	var responseBody interface{}
	var err error
	var validatePassIns validateViaDataBase.ValidatePassFromDB
	fmt.Println("Request is comming Here-----")
	switch loginType {
	case "password":
		responseBody, err = validatePassIns.ValidateViaPassword(userName, password)
		if err != nil {
			// Handle the error, for example, by logging or returning a specific error message
			return "", fmt.Errorf("validation error: %v", err)
		}
	default:
		// If loginType is not "password", return an error
		return "", fmt.Errorf("unsupported login type: %s", loginType)
	}

	return responseBody, nil
}
