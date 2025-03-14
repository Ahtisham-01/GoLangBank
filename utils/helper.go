// change the name of file utils.go
package utils
import (
    "math/rand"  
)

func NewAccount(firstName, lastName string) *Account {
    return &Account{
        ID:        rand.Intn(10000),
        FirstName: firstName,
        LastName:  lastName,
        Number:    int64(rand.Intn(10000000)),
        Balance:   0,
    }
}