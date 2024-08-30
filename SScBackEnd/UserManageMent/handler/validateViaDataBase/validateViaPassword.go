package validateViaDataBase

import (
	"github.com/neeraj777/sscTutorial/DataBase/dataBaseUserManageMent"
)

type ValidatePassFromDB struct {
	userdatabase  dataBaseUserManageMent.User
}
func (validatePass *ValidatePassFromDB)ValidateViaPassword() (interface{}, error){
	//Database Part is not implemented so implementing the dummyuserId and password
	
	return "", nil
}