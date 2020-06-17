package proto

import (
	"bytes"
	"context"
	pb "github.com/hachi-n/grpc_test/pb/auto_gen"
	"io"
	"os"
	"path/filepath"
)

type FileUploader struct{}

func (f *FileUploader) Upload(c context.Context, request *pb.FileRegisterRequest) (*pb.FileRegisterResponse, error) {
	resp := &pb.FileRegisterResponse{
		Identifier: request.Name,
		Status:     "False",
	}

	writer, err := os.Create(filepath.Join("/tmp", request.Name+"_serv.item"))
	if err != nil {
		return resp, err
	}

	buf := bytes.NewBuffer(request.Content)

	io.Copy(writer, buf)
	resp.Status = "True"

	return resp, nil
}
