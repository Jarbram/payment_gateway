package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"payment_gateway/cart_services/utils"
)

type CartItem struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func main() {
	http.HandleFunc("/pay", func(w http.ResponseWriter, r *http.Request) {
		// Realizar una solicitud HTTP al cart_service para obtener los productos en el carrito
		resp, err := http.Get("http://localhost:3000/items")
		if err != nil {
			fmt.Println("Error al conectar con el cart_service:", err)
			utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Error al obtener los productos del carrito."})
			return
		}
		defer resp.Body.Close()

		var cartItems []CartItem
		if err := json.NewDecoder(resp.Body).Decode(&cartItems); err != nil {
			fmt.Println("Error al decodificar la respuesta del cart_service:", err)
			utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Error al procesar la respuesta del cart_service."})
			return
		}

		// Agregar lógica para obtener el método de pago seleccionado desde la solicitud HTTP.

		utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Pago exitoso. Gracias por su compra."})
	})

	http.ListenAndServe(":3001", nil)
}
