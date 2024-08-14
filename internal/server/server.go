package server

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"vpn-handler/internal/database"
)

type Server struct {
	mux    *http.ServeMux
	server *http.Server
	db     *database.DB
}

func NewServer() (*Server, error) {
	db, err := database.NewDB()
	if err != nil {
		return nil, err
	}

	server := &Server{
		mux: http.NewServeMux(),
		db:  db,
	}

	return server, nil
}

func (s *Server) Run() error {
	s.beginRouting()

	err := s.server.ListenAndServe()
	if err != nil {
		return err
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		<-sig
		s.Stop()
		os.Exit(0)

	}()

	return nil
}

func (s *Server) Stop() {
	s.db.Close()
}

func (s *Server) beginRouting() {
	s.server = &http.Server{
		Addr:    "127.0.0.1:8800",
		Handler: s.mux,
	}

	s.mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ZALUPA1488"))
	})

	s.mux.HandleFunc("/create_user", createUser)
	s.mux.HandleFunc("/read_user", readUser)
	s.mux.HandleFunc("/update_user", updateUser)
	s.mux.HandleFunc("/delete_user", deleteUser)
}
