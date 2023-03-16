package server

import (
	"context"
	"net/http"
	"time"
)

/* Структура HTTP-сервера */
type Server struct {
	httpServer *http.Server
}

/* Метод запуска сервера */
func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,       // Адрес, по которому работает сервер
		Handler:        handler,          // Обработчики маршрутов
		MaxHeaderBytes: 1 << 20,          // Максимальный размер байт заголовка запроса
		ReadTimeout:    10 * time.Second, // Время чтения ответа на запрос
		WriteTimeout:   10 * time.Second, // Время записи запроса
	}

	// Начало прослушивания по определённому порту и обработка запросов
	return s.httpServer.ListenAndServe()
}

/* Метод завершения работы сервера */
func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
