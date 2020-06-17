package main

import (
	"context"
	"fmt"
	pb "github.com/hachi-n/grpc_test/pb/auto_gen"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:19003", grpc.WithInsecure())
	if err != nil {
		log.Fatal("client connection error:", err)
	}
	defer conn.Close()

	filePath := "pb/src/file.proto"

	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("err", err)
		os.Exit(1)
	}

	client := pb.NewFileUploaderClient(conn)
	message := &pb.FileRegisterRequest{
		Name: filepath.Base(filePath),
		Content: b,
	}

	resp, err := client.Upload(context.TODO(), message)

	fmt.Printf("result:%#v \n", resp)
	fmt.Printf("error::%#v \n", err)
}
