//go:build !wasm

package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func initEnv() (email, password string, err error) {
	err = godotenv.Load()
	if err != nil {
		err = fmt.Errorf("Error loading .env file: %v", err)
		return
	}
	// check if envs are set: MODEUS_EMAIL and MODEUS_PASSWORD
	email, password = os.Getenv("MODEUS_EMAIL"), os.Getenv("MODEUS_PASSWORD")
	if email == "" || password == "" {
		err = fmt.Errorf("Email or password not set in .env file")
		return
	}
	return
}
