package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// We added this struct, so that we can mock the logger during testing.
type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	h.l.Println("Hello world")

	// Since r.Body is an io.ReadCloser,
	// we can use any normal Go functions
	// to read from it. Hence the use of
	// ioutil.ReadAll() here.
	d, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(rw, "Oops", http.StatusBadRequest)
		return
	}

	// The ResponseWriter has a method on it:
	//		Write([]byte) (int, error)
	// This corresponds to ioWriter. So we can use fmt.Fprintf which takes
	// io.Writer as the first argument.
	fmt.Fprintf(rw, "Hello %s", d)
}
