package routes

import (
	"waysbucks/handlers"
	"waysbucks/pkg/middleware"
	"waysbucks/pkg/mysql"
	"waysbucks/repositories"

	"github.com/gorilla/mux"
)

func CartRoutes(r *mux.Router) {
	cartRepository := repositories.RepositoryCart(mysql.DB)
	h := handlers.HandlerCart(cartRepository)

	r.HandleFunc("/cart", middleware.Auth(h.CreateCart)).Methods("POST")
	r.HandleFunc("/carts-id", middleware.Auth(h.FindCartId)).Methods("GET")
	r.HandleFunc("/cart/{id}", h.UpdateCartTrans).Methods("PATCH")
}
