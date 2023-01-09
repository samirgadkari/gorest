package main

import (
	"log"
	"net/http"
	"os"
)

func main() {

	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)

	// Constructs an HTTP server and registers a default handler to it.
	// Since we have not specified a default handler, it will use the
	// default ServeMux's handler to handle requests to "/".
	// A ServeMux is an HTTP request multiplexer.
	// If you don't set up a handler, the DefaultServeMux will be used.
	// ServeMux has functions like
	//		Handle: Registers an HTTP handler for a path.
	//				A handler is just an interface taking in the
	//				http.ResponseWriter, and *http.Request as parameters.
	//		HandleFunc: Convenience function to add the function to the
	//					DefaultServeMux for a route.
	http.ListenAndServe(":9090", // "ip:port, here we're saying all IPs trying to access using this port"
		nil) // default handler aka default serve mux
}
