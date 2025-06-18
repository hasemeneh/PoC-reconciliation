package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/hasemeneh/PoC-OnlineStore/helper/webservice"
	"github.com/hasemeneh/PoC-reconciliation/svc/reconcileapp/src/config"
	http_public "github.com/hasemeneh/PoC-reconciliation/svc/reconcileapp/src/handler/http/public"
	"github.com/hasemeneh/PoC-reconciliation/svc/reconcileapp/src/service"
)

func main() {
	cfg := config.New().Read()
	serviceObj := service.New(cfg)
	publicHttpHandler := http_public.NewHandler(serviceObj)
	ws := webservice.NewWebService(
		cfg.RunningPort,
		publicHttpHandler,
	)

	go ws.Run()
	select {
	case <-terminateSignal():
		log.Println("Exiting gracefully...")
	}
}

func terminateSignal() chan os.Signal {
	term := make(chan os.Signal)
	signal.Notify(term, os.Interrupt, syscall.SIGTERM)
	return term
}
