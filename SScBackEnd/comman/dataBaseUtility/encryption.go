package dataBaseUtility

import (
	"golang.org/x/crypto/bcrypt"
)

type EncDec struct {

}
func (ed *EncDec)HashPassword(password string) []byte {
    hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        panic(err) 
    }
    return hash
}
