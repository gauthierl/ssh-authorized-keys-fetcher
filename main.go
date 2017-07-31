package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	var err error
	var configPath string
	var userName string

	flag.StringVar(&configPath, "config", "/etc/ssh_authorized_keys_gitlab.json", "the config path")
	flag.Parse()

	if flag.NArg() != 1 {
		fmt.Fprintf(os.Stderr, "Error: %s <user_name>\n", os.Args[0])
		return
	}

	userName = flag.Arg(0)

	config := NewConfig()
	err = config.Load(configPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		return
	}

	cache := NewCache()
	cache.Load(config.CachePath)

	cache.CleanExpired(time.Duration(config.CacheTTL))

	user := cache.FindUserByName(userName)
	if user != nil {
		for _, pubKey := range user.PubKeys {
			fmt.Println(pubKey)
		}
		return
	}

	pubKeys, err := FetchPubKeys(config.FetchURL, userName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		return
	}

	if len(pubKeys) != 0 {
		user = NewUser(userName)
		user.SetPubKeys(pubKeys)

		cache.AppendUser(user)
		cache.Save(config.CachePath)
	}

	for _, pubKey := range pubKeys {
		fmt.Println(pubKey)
	}
}
