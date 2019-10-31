package protobuff

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type server struct {
}

func GrpcServer() {
	listner, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}
	srv := grpc.NewServer()
	RegisterAddServiceServer(srv, &server{})
	reflection.Register(srv)

	if e := srv.Serve(listner); e != nil {
		panic(e)
	}
}

func (s *server) Add(ctx context.Context, request *Request) (*Response, error) {
	a, b := request.GetA(), request.GetB()
	result := a + b
	return &Response{
		Result: result,
	}, nil
}

func (s *server) Multiply(ctx context.Context, request *Request) (*Response, error) {
	a, b := request.GetA(), request.GetB()
	result := a * b
	return &Response{
		Result: result,
	}, nil
}
