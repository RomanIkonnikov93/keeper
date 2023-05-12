// Package server contains server settings.
package server

import (
	"net"
	"os"
	"os/signal"
	"sync"
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

	listen, err := net.Listen("tcp", cfg.GRPCPort)
	if err != nil {
		logger.Fatal("net.Listen: ", err)
	}

	s := grpc.NewServer()
	pb.RegisterKeeperServer(s, gapi.InitServices(rep, cfg, logger))

	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		<-sigint
		logger.Println("server shutdown gracefully")
		s.GracefulStop()
		wg.Done()
	}()

	logger.Info("gRPC server running")
	if err = s.Serve(listen); err != nil {
		logger.Fatal(err)
	}
	wg.Wait()

	return nil
}
