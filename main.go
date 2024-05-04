package main

import (
	"fmt"
	"net/http"

	"github.com/BogosPontifice/distopia_RSA/controller"
	"github.com/BogosPontifice/distopia_RSA/manager"
)

func main() {
	controller.EncryptController()
	controller.DecryptController()
	manager.GenerateRSAKeyPair()
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
