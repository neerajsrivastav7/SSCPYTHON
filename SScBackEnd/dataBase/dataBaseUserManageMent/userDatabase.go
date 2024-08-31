package dataBaseUserManageMent
import (
	"github.com/neeraj777/sscTutorial/comman/userMdel"
)

type User struct {

}

func (usr *User)GetUserByUsername(username string) (userMdel.UserDetails, error) {
	return userMdel.UserDetails{}, nil
}