package dataBaseUserManageMent
import (
	"github.com/neerajsrivastav7/SSCPYTHON/SScBackEnd/comman/userModel"
	"time"
	"fmt"
	"github.com/neerajsrivastav7/SSCPYTHON/SScBackEnd/comman/dataBaseUtility"
)

type User struct {
	ed  dataBaseUtility.EncDec
}

func (usr *User)GetUserByUsername(username string) (userModel.UserDetails, error) {
	 dummyUser := userModel.UserDetails{
        ID:        1,
        FirstName: "John",
        LastName:  "Doe",
        Email:     "john.doe@example.com",
        // Use bcrypt to hash a dummy password "password123"
        PasswordHash: usr.ed.HashPassword("password123"),
        CreatedAt:    time.Now(),
        UpdatedAt:    time.Now(),
    }
	fmt.Println(dummyUser)
	return dummyUser, nil
}