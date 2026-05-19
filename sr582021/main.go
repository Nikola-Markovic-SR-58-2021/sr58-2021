package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sr582021/handlers"
	"sr582021/middleware"
	"sr582021/model"
	"sr582021/repositories"
	"sr582021/services"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	repo := repositories.NewConfigInMemRepository()
	repo.Add(model.Config{Name: "config", Version: 1, Params: map[string]string{"k1": "v1", "k2": "v2"}})
	service := services.NewCOnfigservice(repo)
	handler := handlers.NewConfigHandler(service)
	groupService := services.NewConfigGroupService(repo)
	groupHandler := handlers.NewConfigGroupHandler(groupService)

	router := mux.NewRouter()
	router.Use(middleware.RateLimiterMiddleware)

	router.HandleFunc("/configs/{name}/{version}", handler.Get).Methods("GET")
	router.HandleFunc("/configs", handler.GetAll).Methods("GET")
	router.HandleFunc("/configs/{name}/{version}", handler.Post).Methods("POST")
	router.HandleFunc("/configs/{name}/{version}", handler.DeleteByVersion).Methods("DELETE")

	router.HandleFunc("/configsGroup/{name}/{version}", groupHandler.GetGroup).Methods("GET")
	router.HandleFunc("/configsGroup", groupHandler.GetAllGroups).Methods("GET")
	router.HandleFunc("/configsGroup/{name}/{version}", groupHandler.PostGroup).Methods("POST")
	router.HandleFunc("/configsGroup/{name}/{version}", groupHandler.DeleteGroupByVersion).Methods("DELETE")
	router.HandleFunc("/configsGroup/{name}/{version}", groupHandler.PutGroup).Methods("PUT")
	router.HandleFunc("/configs/groups/{name}/{version}/configs", groupHandler.GetByLabels).Methods("GET")
	router.HandleFunc("/configs/groups/{name}/{version}/configs", groupHandler.DeleteByLabels).Methods("DELETE")

	server := &http.Server{Addr: "0.0.0.0:8000", Handler: router}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Println("server starting")
		if err := server.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				log.Fatal(err)
			}
		}

	}()
	// Go rutina za periodično čišćenje klijenskih zahteva iz rate limiter-a
	go func() {
		for {
			time.Sleep(5 * time.Minute)
			middleware.CleanupClients()
		}
	}()

	//Graceful shutdown
	<-stop

	log.Println("\nGasenje servera...")

	// Timeout za graceful shutdown
	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Println("Greska pri gasenju:", err)
	}

	fmt.Println("Server uspesno ugasen")
}
