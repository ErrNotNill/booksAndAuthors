package booksAndAuthors

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) RunServer(port string, router *http.ServeMux) error {
	s.httpServer = &http.Server{Addr: port,
		Handler:      router,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
	}
	return s.httpServer.ListenAndServe()
}

func (s *Server) ShutdownServer(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

/*func (m *Mux) Router(address string) *Mux {
	m.mux.HandleFunc(address, DataHandler) switch InitHandler <=> interface{}
	return
}

type Mux struct {
	mux http.ServeMux
}

func (m *Mux) Route(address string, handler http.HandlerFunc) {
	m.mux.HandleFunc(address, handler)
}

func DataHandler(w http.ResponseWriter, r *http.Request) {
	InitHandler(w, r)
	SecondHandler(w, r)
	ThirdHandler(w, r)
}

//List of Handlers

func InitHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Docker"))
}
func SecondHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Docker"))
}
func ThirdHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Docker"))
}
*/
