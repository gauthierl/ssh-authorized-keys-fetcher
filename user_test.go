package main

import "testing"

func TestNewUser(t *testing.T) {
	user := NewUser("my_test_name")
	if user.Name != "my_test_name" {
		t.Fail()
	}
	if len(user.PubKeys) != 0 {
		t.Fail()
	}
}

func TestSetPubKeys(t *testing.T) {
	user := NewUser("my_test_name")
	beforeUpdate := user.UpdatedAt
	user.SetPubKeys([]string{
		"ssh-rsa dfskjwklejlwensdfslfjkwklefjskljdlskfjdsld user@host",
	})
	if len(user.PubKeys) != 1 {
		t.Fail()
	}
	if beforeUpdate == user.UpdatedAt {
		t.Fail()
	}
}
