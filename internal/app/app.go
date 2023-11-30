package app

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/artemiyKew/grpc-link-shortener/config"
	"github.com/artemiyKew/grpc-link-shortener/internal/service"
	gatewayapi "github.com/artemiyKew/grpc-link-shortener/pkg/proto/link"
	"github.com/caarlos0/env/v10"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Run() error {
	logrus.Info("starting app")
	cfg := config.Config{}

	if err := env.Parse(&cfg); err != nil {
		return err
	}

	s := grpc.NewServer()
	mux := runtime.NewServeMux()
	ctx, cancel := context.WithCancel(context.Background())

	conn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0"+cfg.GRPCAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		logrus.Fatal(err)
	}

	go runGRPCServer(cfg, s)
	go runHTTPServer(ctx, cfg, mux, conn)

	gracefulShutDown(s, cancel)

	return nil
}

func runGRPCServer(cfg config.Config, s *grpc.Server) {
	gs := service.New()
	gatewayapi.RegisterLinkShortenerServer(s, gs)

	l, err := net.Listen("tcp", cfg.GRPCAddr)
	if err != nil {
		logrus.Fatalf("failed to listen tcp %s, %v", cfg.GRPCAddr, err)
	}

	logrus.Printf("starting listening grpc server at %s", cfg.GRPCAddr)
	if err := s.Serve(l); err != nil {
		logrus.Fatalf("error service grpc server %v", err)
	}

}

func runHTTPServer(ctx context.Context, cfg config.Config, mux *runtime.ServeMux, conn *grpc.ClientConn) {
	if err := gatewayapi.RegisterLinkShortenerHandler(
		ctx, mux, conn,
	); err != nil {
		logrus.Fatal(err)
	}

	logrus.Printf("Starting listen http server at %s", cfg.HTTPAddr)
	if err := http.ListenAndServe(cfg.HTTPAddr, mux); err != nil {
		logrus.Fatal(err)
	}
}

func gracefulShutDown(s *grpc.Server, cancel context.CancelFunc) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(ch)

	sig := <-ch
	errMsg := fmt.Sprintf("Received shutdown signal %v. Shutdown Done", sig)
	logrus.Info(errMsg)
	s.GracefulStop()
	cancel()
}
