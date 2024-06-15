package main

import (
	"app/controller"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Program started")

	email := os.Getenv("JQUANTS_EMAIL")
	pass := os.Getenv("JQUANTS_PASS")

	refreshToken, err := jquants.GetRefreshToken(email, pass)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("RefreshToken: ", refreshToken)
}