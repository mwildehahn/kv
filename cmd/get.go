package cmd

import (
	"context"
	"errors"
	"log"
	"time"

	pb "github.com/mwildehahn/kv/proto"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a value from the key/value server.",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("specify a key to retrieve")
		} else if len(args) > 1 {
			return errors.New("can only specify one key at a time")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		key := args[0]

		conn, err := grpc.Dial("localhost:3000", grpc.WithInsecure())
		if err != nil {
			log.Fatalf("failed connecting to kv server: %v", err)
		}
		defer conn.Close()

		client := pb.NewKeyValueStoreClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		res, err := client.Get(ctx, &pb.GetRequest{Key: key})
		if err != nil {
			log.Fatalf("%v.Get(_) = _, %v: ", client, err)
		}

		log.Println(res)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
