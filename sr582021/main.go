package sr582021

import (
	"log"
	"sr582021/handlers"
	"sr582021/model"
	"sr582021/repositories"
	"sr582021/services"

	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	repo := repositories.NewConfigInMemRepository()
	repo.Add(model.Config{Name: "config", Version: 1, Params: map[string]string{"k1": "v1", "k2": "v2"}})
	service := services.NewCOnfigservice(repo)
	handler := handlers.NewConfigHandler(service)

	router := mux.NewRouter()
	router.HandleFunc("/configs/{name}/{version}", handler.Get).Methods("GET")

	server := &http.Server{Addr: "0.0.0.0:8000", Handler: router}
	log.Println("server starting")
	if err := server.ListenAndServe(); err != nil {
		if err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}

	// Kanal za OS signale
	stop := make(chan os.Signal, 1)

	// Slusanje SIGINT i SIGTERM signala
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Cekanje signala
	<-stop

	fmt.Println("\nGasenje servera...")

	// Timeout za graceful shutdown
	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)
	defer cancel()

	// Graceful shutdown
	if err := server.Shutdown(ctx); err != nil {
		fmt.Println("Greska pri gasenju:", err)
	}

	fmt.Println("Server uspesno ugasen")
}
