package cmd

import (
	"context"
	"errors"
	"log"
	"regexp"
	"time"

	pb "github.com/mwildehahn/kv/proto"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set a value in the key/value server.",
	Args: func(cmd *cobra.Command, args []string) error {
		re := regexp.MustCompile(`\w+=\S`)
		if len(args) != 1 || !re.MatchString(args[0]) {
			return errors.New("provide a key/value to set in the form: key=value")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		re := regexp.MustCompile(`=`)
		s := re.Split(args[0], 2)
		key := s[0]
		value := s[1]

		conn, err := grpc.Dial("localhost:3000", grpc.WithInsecure())
		if err != nil {
			log.Fatalf("failed connecting to kv server: %v", err)
		}
		defer conn.Close()

		client := pb.NewKeyValueStoreClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		res, err := client.Set(ctx, &pb.SetRequest{Key: key, Value: value})
		if err != nil {
			log.Fatalf("%v.Get(_) = _, %v: ", client, err)
		}

		log.Println(res)
	},
}

func init() {
	rootCmd.AddCommand(setCmd)
}
