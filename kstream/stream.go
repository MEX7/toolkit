package kstream

import (
	"context"
	"fmt"
	"sync"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/kl7sn/toolkit/kstream/pb"
)

type Callback func(reply *pb.CellResp)

const (
	MsgConnectReq          = 101
	MsgNodeHeartBeatResp   = 1001
	MsgTestStepUpdateResp  = 1002
	MsgWorkerHeartBeatResp = 1003
)

const ()

type ProxyStream struct {
	sync.Mutex
	stream  pb.Stream_CellClient
	msgChan chan *pb.CellReq
	client  pb.StreamClient
}

func InitStream(client pb.StreamClient) *ProxyStream {
	obj := &ProxyStream{
		msgChan: make(chan *pb.CellReq, 1000),
		client:  client,
	}
	return obj
}

func (p *ProxyStream) Write(data *pb.CellReq) {
	_ = p.GetStream(p.client).Send(data)
}

func (p *ProxyStream) PushChan(info *pb.CellReq) {
	p.msgChan <- info
}

func (p *ProxyStream) SingleThreadRead(callback Callback) {
	var (
		reply *pb.CellResp
		err   error
	)
	for {
		reply, err = p.GetStream(p.client).Recv()
		if err != nil {
			replyStatus, _ := status.FromError(err)
			if replyStatus.Code() == codes.Unavailable {
				fmt.Println("与服务器的连接被断开, 进行重试")
				fmt.Println("Receive reply error:", err.Error())
				p.DelStream()
				continue
			}
			continue
		}
		if reply.Code != 0 {
			fmt.Println("Receive reply code is not 0:", reply.Code)
			continue
		}
		callback(reply)
	}
}

func (p *ProxyStream) SingleThreadWrite() {
	var err error
	for {
		data := <-p.msgChan
		err = p.GetStream(p.client).Send(data)
		if err != nil {
			fmt.Println("Send err:", err.Error())
			continue
		}
	}
}

func (p *ProxyStream) DelStream() {
	p.Lock()
	defer p.Unlock()
	p.stream = nil
}

func (p *ProxyStream) GetStream(client pb.StreamClient) pb.Stream_CellClient {
	p.Lock()
	defer p.Unlock()
	if p.stream != nil {
		return p.stream
	}
	ctx := context.Background()
	for {
		stream, err := client.Cell(ctx)
		if err != nil {
			fmt.Println("stream error: ", err.Error())
			time.Sleep(time.Second)
			continue
		}
		p.stream = stream
		break
	}
	return p.stream
}
