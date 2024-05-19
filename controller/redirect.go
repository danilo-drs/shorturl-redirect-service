package controller

import (
	"net/http"

	"meli-redirect-service/model"

	"github.com/gorilla/mux"
)

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	// get the key from the URL
	key := mux.Vars(r)["key"]
	if key == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"success": false, "message": "Key is required"}`))
		return
	}

	// get the short URL
	shortUrl := model.ShortUrl{Key: key}
	found, err := shortUrl.FillFromKey()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"success": false, "message": "Error getting short URL: ` + err.Error() + `"}`))
		return
	}
	if !found {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"success": false, "message": "Short URL not found"}`))
		return
	}

	// redirect to the original URL
	http.Redirect(w, r, shortUrl.OriginalURL, http.StatusTemporaryRedirect)

}
