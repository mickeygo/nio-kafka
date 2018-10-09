package goka

import (
	"net/http"
)

// Server Http Server
type Server struct {
}

// Handler http
func (s *Server) Handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// started := time.Now()
	})
}

func accessLog() {

}
