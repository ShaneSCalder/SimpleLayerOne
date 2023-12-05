package main

import (
	"SimpleLayerOne/pkg/encryption" // Update this to the correct import path
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
)

// Request structure for encryption
type EncryptRequest struct {
	Data string `json:"data"`
}

// Response structure for encryption
type EncryptResponse struct {
	Message       string `json:"message"`
	EncryptedData string `json:"encryptedData"`
	Key           string `json:"key"`
	Nonce         string `json:"nonce"`
}

// encryptHandler handles requests for data encryption
func encryptHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var req EncryptRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request: "+err.Error(), http.StatusBadRequest)
		return
	}

	data := []byte(req.Data)
	key, err := encryption.GenerateKey()
	if err != nil {
		http.Error(w, "Error generating key: "+err.Error(), http.StatusInternalServerError)
		return
	}

	nonce, err := encryption.GenerateNonce()
	if err != nil {
		http.Error(w, "Error generating nonce: "+err.Error(), http.StatusInternalServerError)
		return
	}

	encryptedData, err := encryption.Encrypt(data, key, nonce)
	if err != nil {
		http.Error(w, "Error encrypting data: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Encode encrypted data, key, and nonce as base64 strings for JSON compatibility
	resp := EncryptResponse{
		Message:       "Data encrypted successfully",
		EncryptedData: base64.StdEncoding.EncodeToString(encryptedData),
		Key:           base64.StdEncoding.EncodeToString(key),
		Nonce:         base64.StdEncoding.EncodeToString(nonce),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func main() {
	http.HandleFunc("/encrypt", encryptHandler)

	fmt.Println("Server is starting on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
