package controller

import (
	"fmt"
	"io"
	"net/http"

	"github.com/BogosPontifice/distopia_RSA/manager"
)

func DecryptController() {
	http.HandleFunc("/decrypt", func(w http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodPost {
			http.NotFound(w, req)
			return
		}

		body, err := io.ReadAll(req.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		msgDecrypted, err := manager.Decrypt(string(body))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		_, err = fmt.Fprintln(w, msgDecrypted)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Println("Decrypted response body: ", msgDecrypted)
	})
}

func EncryptController() {
	http.HandleFunc("/encrypt", func(w http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodPost {
			http.NotFound(w, req)
			return
		}

		body, err := io.ReadAll(req.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		msgEncryptedInBase64, err := manager.Encrypt(string(body))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		_, err = fmt.Fprintln(w, msgEncryptedInBase64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Println("Encrypted response body: ", msgEncryptedInBase64)
	})
}
