package server

import (
	"context"
	"fmt"

	pb "github.com/mwildehahn/kv/proto"
)

var store map[string]string

// KeyValueStoreServer implements the protobuf service
type KeyValueStoreServer struct {
	pb.UnimplementedKeyValueStoreServer

	// TODO: make this a file we serialize
	store map[string]string
}

// Get retrieves a value from the key value store
func (s *KeyValueStoreServer) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	value := s.store[request.Key]
	if value == "" {
		return nil, fmt.Errorf("no value found for key: %s", request.Key)
	}

	return &pb.GetResponse{Key: request.Key, Value: value}, nil
}

// Set sets a value in the key value store
func (s *KeyValueStoreServer) Set(ctx context.Context, request *pb.SetRequest) (*pb.SetResponse, error) {
	s.store[request.Key] = request.Value
	return &pb.SetResponse{Key: request.Key, Value: request.Value}, nil
}

// New creates a new instance of the KeyValueStoreServer
func New() *KeyValueStoreServer {
	s := &KeyValueStoreServer{store: make(map[string]string)}
	return s
}
