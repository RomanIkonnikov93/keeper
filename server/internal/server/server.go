// Package server contains server settings.
package server

import (
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/RomanIkonnikov93/keeper/server/internal/config"
	"github.com/RomanIkonnikov93/keeper/server/internal/gapi"
	pb "github.com/RomanIkonnikov93/keeper/server/internal/proto"
	"github.com/RomanIkonnikov93/keeper/server/internal/repository"
	"github.com/RomanIkonnikov93/keeper/server/pkg/logging"

	"google.golang.org/grpc"
)

// StartServer starts the server.
func StartServer(rep repository.Reps, cfg config.Config, logger *logging.Logger) error {

	idleConnsClosed := make(chan struct{})
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	go func() {
		listen, err := net.Listen("tcp", cfg.GRPCPort)
		if err != nil {
			logger.Fatal("net.Listen: ", err)
		}
		s := grpc.NewServer(
			grpc.UnaryInterceptor(gapi.UnaryUserValidationInterceptor))
		pb.RegisterKeeperServer(s, gapi.InitServices(rep, cfg, logger))

		go func() {
			<-sigint
			s.GracefulStop()
		}()

		logger.Info("gRPC server running")
		if err = s.Serve(listen); err != nil {
			logger.Fatal(err)
		}
	}()

	<-idleConnsClosed
	logger.Println("Servers Shutdown gracefully")

	return nil
}
