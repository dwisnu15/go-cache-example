package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"go_cache_example/db"
	"go_cache_example/src/handler"
	Crepo "go_cache_example/src/repo"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	db.InitDBConnection()
	repo := Crepo.CryptoRepo{}
	serv := handler.CreateCryptoService(repo)

	router := gin.New()

	public := router.Group("/api/v1")
	{
		public.GET("/crypto", serv.GetByID)
		public.GET("/crypto/list", serv.GetAll)
	}

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalf("Failed to initialize server: %v\n", err)
		}
	}()
	log.Printf("Listening on port %v\n", srv.Addr)

	//create kill signal of channel
	stopServer := make(chan os.Signal, 1)

	signal.Notify(stopServer, syscall.SIGINT, syscall.SIGTERM)

	//blocks until a signal is sent to stop server channel
	<-stopServer

	//inform context that it has 7 seconds to finish
	//handling received request
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	//shut down database source
	if err := db.CloseDBConnection(); err != nil {
		log.Fatalf("Shutting down db on problem: %v\n", err)
	}

	//shutdown api (server)
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v\n", err)
	}
}
