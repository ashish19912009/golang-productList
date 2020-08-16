package main

import (
	"os/signal"
	"context"
	"time"
	"github.com/ashish19912009/anotherProject/handlers"
	"os"
	"log"
	"net/http"
)


func main() {
	l := log.New(os.Stdout,"product-api", log.LstdFlags)

	ph := handlers.NewProducts(l)	
	sm := http.NewServeMux()
	sm.Handle("/", ph)
	// sm.Handle("/goodbye", gh)

	s := &http.Server{
		Addr: ":9090",
		Handler: sm,
		IdleTimeout: 120*time.Second,
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	go func() {
	err := s.ListenAndServe()

		if err != nil {
			l.Fatal(err)
		}
	}()

	sinChan := make(chan os.Signal)
	signal.Notify(sinChan, os.Interrupt)
	signal.Notify(sinChan, os.Kill)
	sig := <- sinChan
	l.Println("Recieved terminate, graceful shutdown", sig)
	tc,_ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}