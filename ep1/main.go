package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// The HandleFunc registers a function for the given pattern in the
	// DefaultServeMux. Any request with a pattern not matching any
	// of the handle funcs, will match the default "/" pattern, and
	// the handler function will be executed
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Hello world")
		d, err := ioutil.ReadAll(r.Body) // Since r.Body is an io.ReadCloser,
		// we can use any normal Go functions
		// to read from it. Hence the use of
		// ioutil.ReadAll() here.

		if err != nil {
			http.Error(rw, "Oops", http.StatusBadRequest)
			return
		}

		// The ResponseWriter has a method on it:
		//		Write([]byte) (int, error)
		// This corresponds to ioWriter. So we can use fmt.Fprintf which takes
		// io.Writer as the first argument.
		fmt.Fprintf(rw, "Hello %s", d)
	})

	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("Goodbye world")
	})

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
