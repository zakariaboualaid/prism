package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/zakariaboualaid/prism/internal/backend"
	"github.com/zakariaboualaid/prism/internal/config"
)

var serverPool backend.ServerPool

func main() {

	configMapFlag := flag.String("config-path", "./config-test.yaml", "Path to config file")
	port := flag.Int("port", 8080, "Path to config file")

	flag.Parse()

	fmt.Printf("Parsing config file: %s\n\n", *configMapFlag)

	conf := config.Parse(*configMapFlag)

	for _, service := range conf.Http.Services {
		for _, bknd := range service.Backends {

			proxy := httputil.NewSingleHostReverseProxy(bknd.URL.URL)

			serverPool.AddBackend(&backend.Backend{
				URL:          bknd.URL,
				Alive:        true,
				ReverseProxy: proxy,
			})

		}
	}

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", *port),
		Handler: http.HandlerFunc(lb),
	}

	log.Printf("Load Balancer started at : %d\n", *port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func lb(w http.ResponseWriter, r *http.Request) {
	peer := serverPool.GetNextPeer()
	if peer != nil {
		res := "Hello " + peer.URL.Port()
		w.Write([]byte(res))
		peer.ReverseProxy.ServeHTTP(w, r)
		return
	}
}
