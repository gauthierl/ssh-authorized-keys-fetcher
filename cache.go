package main

import (
	"encoding/json"
	"io/ioutil"
	"time"
)

// Cache is used to cache information to filesystem
type Cache struct {
	Users []User `json:"users"`
}

// NewCache initialize the structure
func NewCache() *Cache {
	return &Cache{}
}

// Load the cache from a JSON file
func (cache *Cache) Load(filePath string) (err error) {
	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		return
	}

	err = json.Unmarshal(fileContent, &cache)
	if err != nil {
		return
	}

	return
}

// Save the cache to a JSON file
func (cache *Cache) Save(filePath string) (err error) {
	data, err := json.Marshal(cache)
	if err != nil {
		return
	}

	err = ioutil.WriteFile(filePath, data, 0600)
	if err != nil {
		return
	}

	return
}

// AppendUser append a user in the cache
func (cache *Cache) AppendUser(user *User) {
	cache.Users = append(cache.Users, *user)
}

// FindUserByName retrieve an user from the cache
func (cache *Cache) FindUserByName(name string) *User {
	for _, user := range cache.Users {
		if user.Name == name {
			return &user
		}
	}
	return nil
}

// CleanExpired remove expired entry
func (cache *Cache) CleanExpired(ttl time.Duration) {
	users := []User{}

	now := time.Now()
	for _, user := range cache.Users {
		since := now.Sub(user.UpdatedAt)

		if since <= ttl {
			users = append(users, user)
		}
	}

	cache.Users = users
}
