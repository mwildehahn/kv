package cmd

import (
	"fmt"
	"log"
	"net"

	pb "github.com/mwildehahn/kv/proto"
	"github.com/mwildehahn/kv/server"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the key/value server.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting service on localhost:3000")
		lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 3000))
		if err != nil {
			log.Fatalf("failed starting server: %v", err)
		}

		grpcServer := grpc.NewServer()
		pb.RegisterKeyValueStoreServer(grpcServer, server.New())
		grpcServer.Serve(lis)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
