// Package api main file where it all starts
package main

import (
	"crypto/rand"
	"encoding/json"
	"math/big"
	"net/http"
	"sync"
	"time"
)

//	type album struct {
//		ID     string  `json:"id"`
//		Title  string  `json:"title"`
//		Artist string  `json:"artist"`
//		Price  float64 `json:"price"`
//	}
// func getAlbums(c string) (string, error) {
// 	return c, nil
// }

type url struct {
	Short string `json:"short_url"`
	Long  string `json:"long_url"`
}

var (
	urlStore = make(map[string]string)
	mutex    = &sync.Mutex{}
)

func generateShortURL() string {
  const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano())

	b := make([]byte, 6)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
	}
	return string(short)
}

func createShortURLHanlder(w http.ResponseWriter, r *http.Request) {
	var request url 
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request", http.StatusAccepted)
		return
	}

	short := generateShortURL()
	mutex.Lock()
	urlStore[short] = request.Long
	mutex.Unlock()

	respons := url{
		short: "http://localhost:8080/" + short,
		Long: request.Long,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(respons)
}

func getOriginalURL.Handler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	short := vars["short"]

	mutex.Lock()
	longURL, exists := urlStore[short]
	mutex.Unlock()

	if err != nil {
		http.Error(w, "Short URL not found", http.StatusNotFound)
		return

}
	http.Redirect(w, r, longURL, http.StatusFoud)
}
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler).Methods("GET")
	r.
}
