package handlers

import (
	"fmt"
	"github.com/neeraj777/sscTutorial/userManageMent/handler/validateViaDataBase"
)
type Validate struct {
	validateViaPassword  *validateViaDataBase.ValidatePassFromDB
}

func (validate *Validate)ValidateLogin(userName string, password string, loginType string) {
	if loginType == "password" {
		responseBody , err := validate.validateViaPassword.ValidateViaPassword()
		if err != nil {

		} else {
			fmt.Println(responseBody)
		}
	}
}