package handler

import (
	"encoding/json"
	"goapp/internal/service"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterUserRoutes(r *mux.Router, userService *service.UserService) {
	r.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		users, err := userService.GetAllUsers()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
	}).Methods("GET")
}
