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

var dbFileName string

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

		kvServer, err := server.New(dbFileName)
		if err != nil {
			log.Fatalf("failed creating the server: %v", err)
		}

		pb.RegisterKeyValueStoreServer(grpcServer, kvServer)
		grpcServer.Serve(lis)
	},
}

func init() {
	serveCmd.Flags().StringVar(&dbFileName, "db-file", "", "Specfiy a db file to read and write to.")
	serveCmd.MarkFlagRequired("db-file")
	rootCmd.AddCommand(serveCmd)
}
