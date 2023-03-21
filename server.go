package server

import (
	"context"
	"net/http"
	"time"
)

/* Структура сервера */
type Server struct {
	httpServer *http.Server
}

/* Метод для запуска сервера с текущими настройками */
func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}

/* Метод завершения работы сервера */
func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
