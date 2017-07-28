package main

import "time"

// User contain all user related data
type User struct {
	Name      string    `json:"name"`
	PubKeys   []string  `json:"pubkeys"`
	UpdatedAt time.Time `json:"updated_at"`
}

// NewUser create a new user and set his name
func NewUser(name string) *User {
	return &User{
		Name: name,
	}
}

// SetPubKeys define pubkeys and update UpdatedAt
func (user *User) SetPubKeys(PubKeys []string) {
	user.PubKeys = PubKeys
	user.UpdatedAt = time.Now()
}
