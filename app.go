package main

import (
	"fmt"

	"github.com/faruqisan/fnb/user"
)

func main() {
	u := user.Create("John Doe", "john@doe.com")

	fmt.Printf("User : %s created with email : %s", u.Name, u.Email)
}
