package main

import (
	"io/ioutil"
	"os"
	"testing"
	"time"
)

func TestNewCache(t *testing.T) {
	cache := NewCache()

	if len(cache.Users) != 0 {
		t.Error("Cache should have no users.")
	}
}

func TestCacheLoad(t *testing.T) {
	var err error

	cache := NewCache()

	err = cache.Load("fixtures/cache_not_found.json")
	if err == nil {
		t.Error(`"fixtures/cache_not_found.json" should return error.`)
	}

	err = cache.Load("fixtures/cache_invalid.json")
	if err == nil {
		t.Error(`"fixtures/cache_invalid.json" should return error.`)
	}

	err = cache.Load("fixtures/cache.json")
	if err != nil {
		t.Error(`"fixtures/cache.json" should be loaded without error.`)
	}

	if len(cache.Users) != 1 {
		t.Error("Cache should contains one user.")
	}

	if cache.Users[0].Name != "user1" {
		t.Error("user1 should be in cache")
	}
}

func TestCacheSave(t *testing.T) {
	cache := NewCache()
	cache.AppendUser(NewUser("user1"))
	cache.AppendUser(NewUser("user2"))

	tmp, _ := ioutil.TempFile("", "")
	defer os.Remove(tmp.Name())

	cache.Save(tmp.Name())

	content, _ := ioutil.ReadFile(tmp.Name())
	expected := `{"users":[{"name":"user1","pubkeys":null,"updated_at":"0001-01-01T00:00:00Z"},{"name":"user2","pubkeys":null,"updated_at":"0001-01-01T00:00:00Z"}]}`

	if string(content) != expected {
		t.Errorf(`Content mismatch got "%s" expected "%s")`, string(content), expected)
	}

	err := cache.Save("fixtures/not_exists_directory/cache.json")
	if err == nil {
		t.Errorf(`Should not be able to write to this file ("fixtures/not_exists_directory/cache.json")`)
	}
}

func TestAppendUser(t *testing.T) {
	user1 := NewUser("user1")
	user2 := NewUser("user2")

	cache := NewCache()

	cache.AppendUser(user1)
	if len(cache.Users) != 1 {
		t.Error("Cache should now have one user.")
	}

	if cache.Users[0].Name != "user1" {
		t.Error("Cache should have user1.")
	}

	cache.AppendUser(user2)
	if len(cache.Users) != 2 {
		t.Error("Cache should now have two users.")
	}
}

func TestCacheFindUserByName(t *testing.T) {
	cache := NewCache()
	cache.AppendUser(NewUser("user1"))
	cache.AppendUser(NewUser("user2"))

	if cache.FindUserByName("user1") == nil {
		t.Errorf("Cannot find user1.")
	}

	if cache.FindUserByName("user2") == nil {
		t.Errorf("Cannot find user2.")
	}

	if cache.FindUserByName("user_not_exists") != nil {
		t.Errorf("Found user_not_exists.")
	}
}

func TestCacheCleanExpired(t *testing.T) {
	user1 := NewUser("user1")
	user1.UpdatedAt = time.Now().Add(-1 * time.Minute)

	user2 := NewUser("user2")
	user2.UpdatedAt = time.Now().Add(-3 * time.Minute)

	cache := NewCache()
	cache.AppendUser(user1)
	cache.AppendUser(user2)

	cache.CleanExpired(2 * time.Minute)

	if len(cache.Users) != 1 {
		t.Error("Only one user should be keeped.")
	}

	if cache.Users[0].Name != "user1" {
		t.Error("Only user1 should be keeped.")
	}
}
