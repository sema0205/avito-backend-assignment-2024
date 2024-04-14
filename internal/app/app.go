package app

import (
	"context"
	"errors"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sema0205/avito-backend-assignment-2024/config"
	delivery "github.com/sema0205/avito-backend-assignment-2024/internal/delivery/http"
	"github.com/sema0205/avito-backend-assignment-2024/internal/repository"
	"github.com/sema0205/avito-backend-assignment-2024/internal/service"
	"github.com/sema0205/avito-backend-assignment-2024/pkg/auth/jwt"
	"github.com/sema0205/avito-backend-assignment-2024/pkg/cache/client"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// @title           Banner Management Service
// @version         1.0
// @description     This service provides an interface for managing and retrieving banners.

// @securityDefinitions.apikey  JWT
// @in                          header
// @name                        Authorization
// @description					JWT token

func Run(configPath string) {

	cfg, err := config.NewConfig(configPath)
	if err != nil {
		log.Fatal(err)
		return
	}

	SetLogrus(cfg.Log.Level)

	ctx := context.Background()
	db, err := pgxpool.New(ctx, cfg.PG.URL)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer db.Close()

	tokenManager, err := jwt.NewTokenManager(cfg.Auth.SignKey, cfg.Auth.TokenTTL)
	if err != nil {
		log.Fatal(err)
		return
	}

	cache := client.NewCacheClient(cfg.Cache.ExpiredTTl, cfg.Cache.CleanupTTL)

	repos := repository.NewRepositories(db, squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar))
	services := service.NewServices(service.Deps{
		Repos:         repos,
		TokenManager:  tokenManager,
		CacheProvider: cache,
	})

	handlers := delivery.NewHandler(services, tokenManager)

	srv := NewServer(cfg, handlers.Init())

	go func() {
		if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
			log.Errorf("error occurred while running http server: %s\n", err.Error())
		}
	}()

	log.Info("Server started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	if err := srv.Stop(ctx); err != nil {
		log.Errorf("failed to stop server: %v", err)
	}
}

type Server struct {
	httpServer *http.Server
}

func NewServer(cfg *config.Config, handler http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:    ":" + cfg.HTTP.Port,
			Handler: handler,
		},
	}
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
