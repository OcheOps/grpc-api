package core

import (
	"context"
)


type NodeServiceGrpcServer struct {
	UnimplementedNodeServiceServer


	CmdChannel chan string
}

func ( n NewNodeServiceGrpcServer)ReportStatus() {
	return &NodeServiceGrpcServer{
		CmdChannel: make(chan string, 1),
	}
}
