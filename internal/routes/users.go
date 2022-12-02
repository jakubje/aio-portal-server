package routes

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/ifreddyrondon/bastion/render"
	"net/http"
	"server/internal/database"
)

type UsersResource struct{}

func (rs UsersResource) LoginUser(w http.ResponseWriter, r *http.Request) {
	user, err := database.GetUserByCredentials(r.Body)

	if err != nil {
		render.JSON.BadRequest(w, err)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (rs UsersResource) SignUpUser(w http.ResponseWriter, r *http.Request) {
	user, err := database.CreateUser(r.Body)

	if err != nil {
		render.JSON.BadRequest(w, err)
		return
	}

	json.NewEncoder(w).Encode(user)
}
func SettingsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Context-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		next.ServeHTTP(w, r)
	})
}

func (rs UsersResource) Routes() chi.Router {
	router := chi.NewRouter()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})

	router.Use(middleware.Logger)
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(SettingsMiddleware)
	router.Use(c.Handler)

	router.Group(func(r chi.Router) {
		r.Post("/login", rs.LoginUser)
		r.Post("/signup", rs.SignUpUser)
	})

	return router
}
