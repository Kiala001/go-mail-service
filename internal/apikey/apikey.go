package apikey

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"os"
	"sync"
)

var mu sync.Mutex
var filePath = "data/apikeys.json"

type Store struct {
	Keys []string `json:"keys"`
}

func Generate() (string, error) {
	b := make([]byte, 24)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	key := "sk_" + hex.EncodeToString(b)

	mu.Lock()
	defer mu.Unlock()

	store := Store{}
	file, _ := os.ReadFile(filePath)
	_ = json.Unmarshal(file, &store)

	store.Keys = append(store.Keys, key)

	data, _ := json.MarshalIndent(store, "", "  ")
	_ = os.WriteFile(filePath, data, 0644)

	return key, nil
}

func IsValid(key string) bool {
	mu.Lock()
	defer mu.Unlock()

	store := Store{}
	file, err := os.ReadFile(filePath)
	if err != nil {
		return false
	}

	_ = json.Unmarshal(file, &store)

	for _, k := range store.Keys {
		if k == key {
			return true
		}
	}
	return false
}