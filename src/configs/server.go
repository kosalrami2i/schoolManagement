package configs

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func StartServer() {
	router := mux.NewRouter()
	port := ":8000"

	fmt.Println("\nListening on port " + port)
	http.ListenAndServe(port, router)
}
