package movies

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

// GetMovies
func GetMovies(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		movies := Movies()
		dataMovies, err := json.Marshal(movies)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(dataMovies)
		return
	}
	http.Error(w, "ERROR....", http.StatusNotFound)
}

// PostMovie
func PostMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var Mov Movie
	if r.Method == "POST" {
		if r.Header.Get("Content-Type") == "application/json" {
			// parse dari json
			decodeJSON := json.NewDecoder(r.Body)
			if err := decodeJSON.Decode(&Mov); err != nil {
				log.Fatal(err)
			}
		} else {
			// parse dari form
			getID := r.PostFormValue("id")
			id, _ := strconv.Atoi(getID)
			title := r.PostFormValue("title")
			getYear := r.PostFormValue("year")
			year, _ := strconv.Atoi(getYear)
			Mov = Movie{
				ID:    id,
				Title: title,
				Year:  year,
			}
		}

		dataMovie, _ := json.Marshal(Mov) // to byte
		w.Write(dataMovie)                // cetak di browser
		return
	}

	http.Error(w, "NOT FOUND", http.StatusNotFound)
	return
}
