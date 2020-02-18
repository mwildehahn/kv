# kv

Work in progress key value store.

## Generate Protobufs

In the current directory, run:

```
$ protoc -I proto proto/kv.proto --go_out=plugins=grpc:proto
```

## Command Line

```
➜ kv help
Work in progress key/value store.

Usage:
  kv [command]

Available Commands:
  delete      Delete a key from the key/value server.
  get         Get a value from the key/value server.
  help        Help about any command
  serve       Start the key/value server.
  set         Set a value in the key/value server.

Flags:
  -h, --help   help for kv

Use "kv [command] --help" for more information about a command.
```
