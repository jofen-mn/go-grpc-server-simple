package main

import (
	"runtime"
	"net"
	"log"
	"google.golang.org/grpc"
	"context"
	"strconv"
	"go-grpc-server-simple/inf"
)

type Data struct {}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	lis, err := net.Listen("tcp", ":" + "41005")
	if err != nil {
		log.Fatal("failed to listen: %v", err)
	}

	log.Println("grpc server in: %s", "41005")

	s := grpc.NewServer()
	inf.RegisterDataServer(s, &Data{})
	s.Serve(lis)
	log.Println("grpc server in: %s", "41005")
}

func (d *Data) GetUser(ctx context.Context, req *inf.UserRq) (response *inf.UserRp, err error) {
	response = &inf.UserRp{
		Name:strconv.Itoa(int(req.Id)) + ":test",
	}
	return response, err
}
