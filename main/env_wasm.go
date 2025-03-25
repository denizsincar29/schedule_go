//go:build wasm

package main

import (
	"fmt"
	"syscall/js"
)

func initEnv() (email, password string, err error) {
	// wasm prompt for email and password
	// check if envs are set: MODEUS_EMAIL and MODEUS_PASSWORD
	email, password = js.Global().Get("prompt").Invoke("Enter email:").String(), js.Global().Get("prompt").Invoke("Enter password:").String()
	if email == "" || password == "" {
		err = fmt.Errorf("email or password not valid")
		return
	}
	return
}
