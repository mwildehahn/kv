package server

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/golang/protobuf/proto"
	pb "github.com/mwildehahn/kv/proto"
)

// KeyValueStoreServer implements the protobuf service
type KeyValueStoreServer struct {
	pb.UnimplementedKeyValueStoreServer

	dbFileName string
	store      *pb.DataStore
}

// Get retrieves a value from the key value store
func (s *KeyValueStoreServer) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	value := s.store.Data[request.Key]
	if value == "" {
		return nil, fmt.Errorf("no value found for key: %s", request.Key)
	}

	return &pb.GetResponse{Key: request.Key, Value: value}, nil
}

// Set sets a value in the key value store
func (s *KeyValueStoreServer) Set(ctx context.Context, request *pb.SetRequest) (*pb.SetResponse, error) {
	s.store.Data[request.Key] = request.Value
	return &pb.SetResponse{Key: request.Key, Value: request.Value}, nil
}

// Delete deletes a key from the key value store
func (s *KeyValueStoreServer) Delete(ctx context.Context, request *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	delete(s.store.Data, request.Key)
	return &pb.DeleteResponse{Key: request.Key}, nil
}

// Shutdown safely shuts down the key value store server
func (s *KeyValueStoreServer) closeHandler() {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGINT)
	go func() {
		<-c

		out, err := proto.Marshal(s.store)
		if err != nil {
			log.Fatalf("failed serializing data: %v", err)
		}

		if err := ioutil.WriteFile(s.dbFileName, out, 0644); err != nil {
			log.Fatalf("failed writing data: %v", err)
		}

		os.Exit(0)
	}()
}

// New creates a new instance of the KeyValueStoreServer
func New(dbFileName string) (*KeyValueStoreServer, error) {
	in, err := ioutil.ReadFile(dbFileName)
	if err != nil {
		// TODO wrap error
		return nil, err
	}

	store := &pb.DataStore{}
	if err := proto.Unmarshal(in, store); err != nil {
		// TODO wrap error
		return nil, err
	}

	if store.Data == nil {
		store.Data = make(map[string]string)
	}

	s := &KeyValueStoreServer{store: store, dbFileName: dbFileName}
	s.closeHandler()

	return s, nil
}
