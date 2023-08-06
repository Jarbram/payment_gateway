package main

import (
	"net/http"
	"payment_gateway/cart_services/cart"
	"payment_gateway/cart_services/utils"
)

// ...

func main() {
	shoppingCart := cart.NewCart()

	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		// Agregar lógica para decodificar los datos y agregar el producto al carrito.
		shoppingCart.AddItem(cart.NewCartItem("Producto de ejemplo", 10.99))
		utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Producto agregado al carrito."})
	})

	http.HandleFunc("/items", func(w http.ResponseWriter, r *http.Request) {
		// Responder con la lista de productos en formato JSON
		cartItemsJSON, err := shoppingCart.GetCartItemsJSON()
		if err != nil {
			utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Error al obtener los productos del carrito."})
			return
		}

		// Establecer el encabezado Content-Type como application/json
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(cartItemsJSON)
	})

	http.ListenAndServe(":3000", nil)
}
