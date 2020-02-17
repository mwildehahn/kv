# kv

Work in progress key value store.

## Generate Protobufs

In the current directory, run:

```
$ protoc -I proto proto/kv.proto --go_out=plugins=grpc:proto
```
