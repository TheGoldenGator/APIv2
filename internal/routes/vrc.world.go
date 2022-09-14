package routes

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/thegoldengator/APIv2/internal/apis"
)

func VRCWorld(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	value := chi.URLParam(r, "username")

	worldData, err := apis.VRChat.GetWorld(value)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte(fmt.Sprintf("%v - %v", worldData.Name, worldData.AuthorName)))
}
