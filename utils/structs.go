package utils

type Account struct {
    ID        int `json:"id"`
    FirstName string `json:"firstName"`
    LastName  string`json:"lastName"`
    Number    int64 `json:"number"`
    Balance   int64 `json:"balance"`
}

