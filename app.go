package main

import (
	"fmt"
	"log"

	"github.com/faruqisan/fnb/storages/cache"

	"github.com/faruqisan/fnb/user"
)

func main() {
	u := user.Create("John Doe", "john@doe.com")

	fmt.Printf("User : %s created with email : %s \n", u.Name, u.Email)

	cache.Init("")

	c := cache.Get()
	c.Set(u.Name, u.Email)

	fmt.Printf("User %s, stored on cache \n", u.Name)

	val, err := c.Get(u.Name)
	if err != nil {
		log.Println("Err ! ", err)
	}

	fmt.Printf("Retrived value of user %s is : %s", u.Name, val)

}
