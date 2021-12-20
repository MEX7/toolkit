package main

import (
	"fmt"
	"time"

	"github.com/gotomicro/ego"
	"github.com/gotomicro/ego/core/elog"
	"github.com/gotomicro/ego/server"
	"github.com/gotomicro/ego/server/egrpc"

	"github.com/kl7sn/toolkit/knet"
	"github.com/kl7sn/toolkit/kstream"
	"github.com/kl7sn/toolkit/kstream/pb"
)

type Greeter struct {
	server *egrpc.Component
}

type grpcServer struct {
}

var GrpcServer pb.StreamServer = &grpcServer{}

func main() {
	// server
	if err := ego.New().Invoker(process).Serve(func() server.Server {
		ser := egrpc.Load("server.grpc").Build()
		pb.RegisterStreamServer(ser.Server, GrpcServer)
		return ser
	}()).Run(); err != nil {
		elog.Panic("startup", elog.Any("err", err))
	}
}

func process() error {
	kstream.InitStreamStores()
	go func() {
		for {
			err := kstream.StreamStore.Broadcast(&pb.CellResp{
				Msg: []byte(fmt.Sprintf("%d", time.Now().Unix())),
			})
			if err != nil {
				fmt.Println("broadcast error: ", err.Error())
			}
			time.Sleep(time.Second)
		}
	}()
	return nil
}

func (g grpcServer) Cell(stream pb.Stream_CellServer) error {
	ip := knet.GetPeerIP(stream.Context())
	go func() {
		// todo
		<-stream.Context().Done()
		// 链接结束做什么
		kstream.StreamStore.DeleteStream(ip)
	}()
	kstream.StreamStore.AddStream(ip, stream)

	for {
		req, err := stream.Recv()
		if err != nil {
			break
		}
		elog.Info("cell", elog.Any("req", req))
	}
	return nil
}
