package main

import (
	"github.com/hachi-n/grpc_test/lib/proto"
	pb "github.com/hachi-n/grpc_test/pb/auto_gen"
	"log"
	"net"
	"google.golang.org/grpc"

)

func main() {
	listenPort, err := net.Listen("tcp", ":19003")
	if err != nil {
		log.Fatalln(err)
	}
	server := grpc.NewServer()
	fileService := &proto.FileUploader{}
	pb.RegisterFileUploaderServer(server, fileService)

	server.Serve(listenPort)
}
