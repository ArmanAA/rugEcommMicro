package main

import (
	"context"
	"fmt"
	"net"

	pceadderpb "github.com/ArmanAA/pce/src/proto/pceadder"
	"google.golang.org/grpc"
)

// PORT that pceadder listens to
const PORT = ":4041"

type server struct {
	pceadderpb.UnimplementedPceAdderServer
}

func main() {
	listener, err := net.Listen("tcp", PORT)
	if err != nil {
		panic(err)
	}
	srv := grpc.NewServer()
	fmt.Println("server running on port ", PORT)
	pceadderpb.RegisterPceAdderServer(srv, &server{})
	if err := srv.Serve(listener); err != nil {
		panic(err)
	}

}

func (s *server) Add(ctx context.Context, req *pceadderpb.Request) (*pceadderpb.Response, error) {
	res := req.A + req.B
	return &pceadderpb.Response{
		Result: res,
	}, nil
}
