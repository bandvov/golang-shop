package interfaces

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

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
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	carts, err := h.CartService.GetCarts(ctx)
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
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	cart, err := h.CartService.GetCartByID(ctx, id)
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

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	if err := h.CartService.AddToCart(ctx, &cart); err != nil {
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

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	if err := h.CartService.RemoveFromCart(ctx, id); err != nil {
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

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	if err := h.CartService.UpdateCart(ctx, &cart); err != nil {
		if err == carts.ErrCartNotFound {
			http.Error(w, "cart not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
