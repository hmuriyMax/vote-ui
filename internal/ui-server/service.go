package ui_server

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Server struct {
	srv http.Server
}

func NewServer(port int) *Server {
	router := gin.Default()
	sv := &Server{
		srv: http.Server{
			Addr:    fmt.Sprintf("localhost:%d", port),
			Handler: router,
		},
	}
	sv.initHandlers(router)
	return sv
}

func (s *Server) initHandlers(router *gin.Engine) {
	router.LoadHTMLGlob("web/*.html")
	router.StaticFS("/res/", http.Dir("web/res/"))
	router.GET("/", s.index)
}

// Start запускает сервер
func (s *Server) Start(ctx context.Context) error {
	localCtx, cancelFunc := context.WithCancel(ctx)
	defer cancelFunc()

	errChan := make(chan error, 1)
	go func(ctx context.Context) {
		err := s.srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			errChan <- fmt.Errorf("failed to start httpServer: %w", err)
		}
	}(localCtx)

	log.Printf("started server at %s", s.srv.Addr)
	select {
	case <-ctx.Done():
		err := s.srv.Shutdown(context.Background())
		if err != nil {
			return err
		}
	case err := <-errChan:
		return fmt.Errorf("failed to start: %w", err)
	}
	return nil
}
