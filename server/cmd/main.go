package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"buf.build/gen/go/wcygan/hello/connectrpc/go/hello/v1/hellov1connect"
	v1 "buf.build/gen/go/wcygan/hello/protocolbuffers/go/hello/v1"
	"connectrpc.com/connect"
	"connectrpc.com/grpcreflect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type GreeterServer struct{}

func (s *GreeterServer) SayHello(
	ctx context.Context,
	req *connect.Request[v1.SayHelloRequest],
) (*connect.Response[v1.SayHelloResponse], error) {
	log.Println("Request headers: ", req.Header())
	res := connect.NewResponse(&v1.SayHelloResponse{
		Message: fmt.Sprintf("Hello %s!", req.Msg.Name),
	})
	res.Header().Set("Greet-Version", "v1")
	return res, nil
}

func main() {
	greeter := &GreeterServer{}
	mux := http.NewServeMux()
	path, handler := hellov1connect.NewGreeterServiceHandler(greeter)
	mux.Handle(path, handler)

	reflector := grpcreflect.NewStaticReflector(
		hellov1connect.GreeterServiceName, // Add the service name here
	)
	mux.Handle(grpcreflect.NewHandlerV1(reflector))
	mux.Handle(grpcreflect.NewHandlerV1Alpha(reflector))

	fmt.Println("Server listening on :8080 with reflection enabled")
	err := http.ListenAndServe(
		"localhost:8080",
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
		// mux, // Simple HTTP/1.1 for now
	)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
