# Grpc-mini-project
This repository is is created in order understand the RPC calls.

## gRPC
RPC is a modern open source high performance Remote Procedure Call (RPC) framework that can run in any environment. It can efficiently connect services in and across data centers with pluggable support for load balancing, tracing, health checking and authentication. It is also applicable in last mile of distributed computing to connect devices, mobile applications and browsers to backend services.

All four kinds of RPC calls(Uanry Api, Client Streaming, Server Streaming, Bi-directional streaming) have been used in here for Demo. Run the Client and Server applications independently.

## command to set the path
```
export PATH="$PATH:$(go env GOPATH)/bin"
```

## command to generate Protobuff file
```
protoc --go_out=. --go-grpc_out=. pb/auth.proto
```
