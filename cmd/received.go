package cmd

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"path/filepath"

	porter "github.com/fukurose/sam/grpc"
	"github.com/spf13/cobra"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type porterServer struct {
	porter.UnimplementedPorterServiceServer
}

func (s *porterServer) ListSegmentStream(req *porter.LSRequest, stream porter.PorterService_ListSegmentStreamServer) error {
	err := filepath.Walk(req.GetPath(), func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			f, err := os.Open(path)
			if err != nil {
				return err
			}
			defer f.Close()

			if err := stream.Send(&porter.LSResponse{
				Path: path,
				Size: info.Size(),
			}); err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (s *porterServer) OrderStream(req *porter.OrderRequest, stream porter.PorterService_OrderStreamServer) error {
	f, err := os.Open(req.GetFilePath())
	if err != nil {
		return err
	}
	defer f.Close()
	buf := make([]byte, 4194304) // Limit 4MB
	for {
		n, err := f.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		if err := stream.Send(&porter.OrderResponse{
			Data: buf[:n],
		}); err != nil {
			return err
		}
	}

	return nil
}

var receivedCmd = &cobra.Command{
	Use:   "received",
	Short: "Start gRPC server and waiting for orders",

	RunE: func(cmd *cobra.Command, args []string) error {
		address, err := cmd.Flags().GetString("address")
		if err != nil {
			return err
		}

		listener, err := net.Listen("tcp", address)
		if err != nil {
			panic(err)
		}

		s := grpc.NewServer()
		porter.RegisterPorterServiceServer(s, &porterServer{})
		reflection.Register(s)

		go func() {
			log.Printf("start gRPC server: %s", address)
			s.Serve(listener)
		}()

		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt)
		<-quit
		log.Println("stopping gRPC server...")
		s.GracefulStop()
		return nil
	},
}

func init() {
	rootCmd.AddCommand(receivedCmd)
}
