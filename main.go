package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type URL struct{
	ID string `json:"id"`
	OriginalURL string `json:"originalurl"`
	ShortURL string `json:"short_url"`
	Creation time.Time `json:"creation"`
}
/*



*/
var URLDB = make(map[string]URL)

func generateShortURL(OriginalURL string) string{
	hasher := md5.New() // It creates an empty MD5 hash calculator that’s ready to take data and compute a hash.
	hasher.Write([]byte(OriginalURL)) // It converts OriginalURL string to a byte slice
	fmt.Println("hasher : ",hasher) // It will print the internal state of the hasher not the hash value.
	data := hasher.Sum(nil) // Finishes the hashing and gives you the final hash value in raw bytes.
	fmt.Println("hasher data : ",data) // Prints the hash value.
	hash := hex.EncodeToString(data) // It merges all bytes into one continuous hexadecimal string.
	fmt.Println("Encode to String: ",hash) // Prints the readable hash string.
	fmt.Println("Final string: ",hash[:8]) // Prints the short hash string (upto 8 characters).

	return hash[:8]
}

func createURL(OriginalURL string) string {
	shortURL := generateShortURL(OriginalURL)
	id := shortURL // Use the short url id for simplicity.
	URLDB[id] = URL{
		ID : id,
		OriginalURL : OriginalURL,
		ShortURL : shortURL,
		Creation : time.Now(),
	}

	return shortURL
}

func getURL(id string) (URL, error){
	url, ok := URLDB[id]
	if !ok {
		return URL{}, errors.New("URL Not Found")
	}
	return url, nil
}

//when we paste root link on browser,it will print GET Method in response.
func rootPageURL(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "GET Method")
	
}

// shortURLHandler is the server side function which processes the request and gives short url in response.
func shortURLHandler(w http.ResponseWriter, r *http.Request){
	var data struct{
		URL string `json:"url"`
	}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil{
		http.Error(w, "Invalid request Body", http.StatusBadRequest)
		return
	}

	short_URL := createURL(data.URL)
	// fmt.Fprintf(w,short_URL)
	response := struct{
		ShortURL string `json:"short_url"`
	}{ShortURL : short_URL}

// Server tells the client that we are sending response in json format.
w.Header().Set("Content-type","application/json")

// Writes the json into the http response body and sends it to client.
	err = json.NewEncoder(w).Encode(&response)
	if err != nil{
		http.Error(w,"Failed to Encode Response", http.StatusInternalServerError)
		return
	}

}

func redirectURLHandler(w http.ResponseWriter, r *http.Request){
	id := r.URL.Path[len("/redirect/"):]
	url, err := getURL(id)
	if err != nil{
		http.Error(w, "Invalid Request", http.StatusNotFound)
	}
	http.Redirect(w, r, url.OriginalURL, http.StatusFound)


}


func main() {
	// fmt.Println("Starting URL-shortener...")
	// OriginalURL := "https://github.com/Krish-Chaudhary167"
	// generateShortURL(OriginalURL)

	// Register the handler function to handle the requests to the root URL ("/")
	http.HandleFunc("/", rootPageURL)
	http.HandleFunc("/shorten", shortURLHandler)
	http.HandleFunc("/redirect/", redirectURLHandler)


	// Start the HTTP Server on port 3000
	fmt.Println("Starting Server on port 3000: ")
	err := http.ListenAndServe(":3000", nil)
	if err != nil{
		fmt.Println("Error on starting server :",err)
		return
	}


}