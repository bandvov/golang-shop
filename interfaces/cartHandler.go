package interfaces

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/bandvov/golang-shop/application"
	"github.com/bandvov/golang-shop/domain/carts"
)

type CartHandler struct {
	CartService *application.CartService
}

func NewCartHandler(service *application.CartService) *CartHandler {
	return &CartHandler{CartService: service}
}

func (h *CartHandler) GetCarts(w http.ResponseWriter, r *http.Request) {
	carts, err := h.CartService.GetCarts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(carts)
}

func (h *CartHandler) GetCart(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid cart ID", http.StatusBadRequest)
		return
	}
	cart, err := h.CartService.GetCartByID(id)
	if err != nil {
		if err == carts.ErrCartNotFound {
			http.Error(w, "cart not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(cart)
}

func (h *CartHandler) AddToCart(w http.ResponseWriter, r *http.Request) {
	var cart carts.Cart
	if err := json.NewDecoder(r.Body).Decode(&cart); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	if err := h.CartService.AddToCart(&cart); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *CartHandler) RemoveFromCart(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid cart ID", http.StatusBadRequest)
		return
	}

	if err := h.CartService.RemoveFromCart(id); err != nil {
		if err == carts.ErrCartNotFound {
			http.Error(w, "cart not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *CartHandler) UpdateCart(w http.ResponseWriter, r *http.Request) {
	var cart carts.Cart
	if err := json.NewDecoder(r.Body).Decode(&cart); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	if err := h.CartService.UpdateCart(&cart); err != nil {
		if err == carts.ErrCartNotFound {
			http.Error(w, "cart not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
