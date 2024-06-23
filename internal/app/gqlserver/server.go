package gqlserver

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/kuromii5/posts/internal/graphql/app/resolvers"
	"github.com/kuromii5/posts/internal/graphql/graph"
	"github.com/kuromii5/posts/internal/service"
)

type GQLServer struct {
	log     *slog.Logger
	port    int
	srv     *handler.Server
	httpsrv *http.Server
}

func New(
	log *slog.Logger,
	port int,
	service *service.Service,
) *GQLServer {
	resolver := &resolvers.Resolver{Service: *service}
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))
	srv.AddTransport(&transport.Websocket{}) // add websocket for subscriptions

	return &GQLServer{
		log:  log,
		port: port,
		srv:  srv,
	}
}

func (a *GQLServer) run() error {
	const f = "gqlserver.run"
	mux := http.NewServeMux()
	mux.Handle("/graphql", a.srv)
	mux.Handle("/playground", playground.Handler("GraphQL Playground", "/graphql"))
	mux.Handle("/subscriptions", a.srv)

	addr := fmt.Sprintf(":%d", a.port)
	a.httpsrv = &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	go func() {
		if err := a.httpsrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			a.log.Error("Failed to start GraphQL server", slog.String("func", f), slog.String("error", err.Error()))
			panic(err)
		}
	}()

	a.log.Info("GraphQL server started", slog.Int("port", a.port), slog.String("func", f))

	// Wait for interrupt signal to gracefully shutdown the server with a timeout.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	a.log.Info("Shutting down server...")

	if err := a.httpsrv.Shutdown(context.Background()); err != nil {
		return fmt.Errorf("%s:%w", f, err)
	}

	a.log.Info("Server gracefully stopped")
	return nil
}

func (a *GQLServer) MustRun() {
	if err := a.run(); err != nil {
		panic(err)
	}
}

func (a *GQLServer) Shutdown() {
	const f = "gqlserver.Shutdown"
	a.log.Info("Stopping GraphQL server", slog.String("func", f))
	if err := a.httpsrv.Shutdown(context.Background()); err != nil {
		a.log.Error("Failed to gracefully shutdown the server", slog.String("func", f), slog.String("error", err.Error()))
	}
}
