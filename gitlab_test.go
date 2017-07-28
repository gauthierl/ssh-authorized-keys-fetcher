package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGitlabFetchPubkeys(t *testing.T) {
	var pubkeys []string
	var err error

	validPubKeys := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, `ssh-rsa dsfj435224jl24j3/34j2l4j23l4j/234ljk243lj24234/sdfs comment1`)
		fmt.Fprintln(w, `ssh-rsa usfj4356754jl24j3/34j2l4j23l4j/234ljk243l4235we/rwe comment2`)
	}))
	defer validPubKeys.Close()

	notFoundPubKeys := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer notFoundPubKeys.Close()

	pubkeys, err = GitlabFetchPubKeys(validPubKeys.URL+"/{{ .UserName }}", "user1")
	if err != nil {
		t.Errorf("Should not return an error (%s).", err.Error())
	}
	if len(pubkeys) != 2 {
		t.Error("Should have two pubkeys.")
	}

	pubkeys, err = GitlabFetchPubKeys(notFoundPubKeys.URL+"/{{ .UserName }}", "user1")
	if err == nil {
		t.Error("Should return an error.")
	}
	if len(pubkeys) != 0 {
		t.Error("Should not return pubkeys on errors.")
	}
}
