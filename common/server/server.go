package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/pkg/errors"

	"github.com/veljkomatic/user-account/common/log"
)

const (
	DefaultIntServerPort         = 6060
	defaultServerPort            = 8000
	defaultServerReadTimeout     = 300
	defaultServerWriteTimeout    = 300
	defaultServerIdleTimeout     = 1200
	defaultServerShutdownTimeout = 30
)

type Config struct {
	Port int `json:"port" yaml:"port"`

	ReadTimeout  int `json:"read_timeout"  yaml:"read_timeout"`
	WriteTimeout int `json:"write_timeout" yaml:"write_timeout"`
	IdleTimeout  int `json:"idle_timeout"  yaml:"idle_timeout"`

	ShutdownTimeout int `json:"shutdown_timeout" yaml:"shutdown_timeout"`
}

func NewDefaultIntServerConfig() *Config {
	return &Config{
		Port:            DefaultIntServerPort,
		ReadTimeout:     defaultServerReadTimeout,
		WriteTimeout:    defaultServerWriteTimeout,
		IdleTimeout:     defaultServerIdleTimeout,
		ShutdownTimeout: defaultServerShutdownTimeout,
	}
}

func (cfg *Config) Validate() error {
	if cfg.Port == 0 {
		cfg.Port = defaultServerPort
	}

	if cfg.ReadTimeout == 0 {
		cfg.ReadTimeout = defaultServerReadTimeout
	}
	if cfg.WriteTimeout == 0 {
		cfg.WriteTimeout = defaultServerWriteTimeout
	}
	if cfg.IdleTimeout == 0 {
		cfg.IdleTimeout = defaultServerIdleTimeout
	}

	if cfg.ShutdownTimeout == 0 {
		cfg.ShutdownTimeout = defaultServerShutdownTimeout
	}

	return nil
}

type Server struct {
	srv *http.Server

	shutdownTimeout time.Duration
}

func ConfigureServer(cfg *Config) *Server {
	return &Server{
		srv: &http.Server{
			Addr:         fmt.Sprintf(":%d", cfg.Port),
			ReadTimeout:  time.Duration(cfg.ReadTimeout) * time.Second,
			WriteTimeout: time.Duration(cfg.WriteTimeout) * time.Second,
			IdleTimeout:  time.Duration(cfg.IdleTimeout) * time.Second,
		},

		shutdownTimeout: time.Duration(cfg.ShutdownTimeout) * time.Second,
	}
}

// ListenAndServe listens on the TCP network address srv.Addr and then
// calls Serve to handle requests on incoming connections.
// Accepted connections are configured to enable TCP keep-alives.
//
// ListenAndServe always returns a non-nil error unless Shutdown is called.
func (s *Server) ListenAndServe(ctx context.Context, handler http.Handler) error {
	s.srv.Handler = handler
	log.Info(ctx, "http server started listening for requests", s)

	shutdownCtx, shutdownCancel := context.WithTimeout(ctx, s.shutdownTimeout)
	shutdownCh := make(chan error, 1)
	go func() {
		defer close(shutdownCh)
		defer shutdownCancel()
		<-ctx.Done()
		shutdownCh <- s.Shutdown(shutdownCtx)
	}()

	err := s.srv.ListenAndServe()
	// ListenAndServe always returns a non-nil error. After Shutdown or Close,
	// the returned error is ErrServerClosed.
	if errors.Is(err, http.ErrServerClosed) {
		defer log.Info(ctx, "http server stopped listening for requests", s)
		err := <-shutdownCh
		if err != nil {
			defer func() {
				if err := s.srv.Close(); err != nil {
					log.Error(ctx, err, "http server failed closing", s)
				}
			}()
			return err
		}
		return nil
	}

	log.Error(ctx, err, "http server failed listening for requests", s)
	return err
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}

func (s *Server) LogFields() log.Fields {
	return log.Fields{"addr": s.srv.Addr, "shutdown_timeout": s.shutdownTimeout}
}
