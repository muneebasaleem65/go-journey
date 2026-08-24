package main

import (
	"fmt"
	"github.com/muneebasaleem65/podcast/auth"
	"github.com/muneebasaleem65/podcast/user"
)

func main() {
	auth.LoginWithCredentials("muneeba", "secret")
	session := auth.GetSession()
	fmt.Println("Session:", session)

	user := user.User{
		Email: "user@email.com",
		Name:  "John Doe",
	}

	fmt.Println(user.Email, user.Name)
}
