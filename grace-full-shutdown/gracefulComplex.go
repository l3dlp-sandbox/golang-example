package grace_full_shutdown

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"context"
	"time"
)

type Server struct {
	logger *log.Logger
	mux *http.ServeMux
}

func NewServer(options ...func(server *Server) )*Server{
	s:= &Server{
		logger: log.New(os.Stdout,"",0),
		mux: http.NewServeMux(),
	}
	for _,f :=range options{
		f(s)
	}

	s.mux.HandleFunc("/",s.index)
	return s
}

func (s * Server) ServeHTTP(w http.ResponseWriter,r *http.Request){
	s.mux.ServeHTTP(w,r)
}

func (s *Server)index(w http.ResponseWriter, r *http.Request)  {
	s.logger.Println("GET /")
	w.Write([]byte("Hello World"))
}

func ComplexShutDownMain(){
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	s:=NewServer(func(s *Server) {

	})

	h := &http.Server{
		Addr: ":8080",
		Handler: s,
	}
	go func() {
		if err := h.ListenAndServe(); err != nil {

		}
	}()

	ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
	h.Shutdown(ctx)
}
