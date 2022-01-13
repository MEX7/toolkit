package kstream

import (
	"context"
	"fmt"
	"io"
	"sync"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/kl7sn/toolkit/kstream/pb"
	"github.com/kl7sn/toolkit/xgo"
)

type Callback func(reply *pb.CellResp)

const (
	MsgConnectReq          = 101
	MsgNodeHeartBeatResp   = 1001
	MsgTestStepUpdateResp  = 1002
	MsgWorkerHeartBeatResp = 1003
)

type ModelType int

const (
	ModelTypeOnlyWrite ModelType = iota + 1
	ModelTypeOnlyRead
	ModelTypeBoth
)

type ProxyStream struct {
	sync.Mutex
	stream  pb.Stream_CellClient
	msgChan chan *pb.CellReq
	client  pb.StreamClient
}

func InitStream(client pb.StreamClient, model ModelType, callback Callback) *ProxyStream {
	obj := &ProxyStream{
		msgChan: make(chan *pb.CellReq, 5000),
		client:  client,
	}
	switch model {
	case ModelTypeBoth:
		obj.writer()
		obj.reader(callback)
	case ModelTypeOnlyRead:
		obj.reader(callback)
	case ModelTypeOnlyWrite:
		obj.writer()
	}
	return obj
}

func (p *ProxyStream) PushChan(info *pb.CellReq) {
	p.msgChan <- info
}

func (p *ProxyStream) reader(callback Callback) {
	var (
		reply *pb.CellResp
		err   error
	)
	xgo.Go(func() {
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
	})
}

func (p *ProxyStream) writer() {
	xgo.Go(func() {
		for {
			data := <-p.msgChan
			err := p.GetStream(p.client).Send(data)
			if err != nil {
				if err == io.EOF {
					p.DelStream()
					fmt.Println("Send try again:", string(data.GetMsg()))
				} else {
					fmt.Println("Send err:", err.Error(), string(data.GetMsg()))
				}
				p.msgChan <- data
				continue
			}
		}
	})
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
		} else {
			p.stream = stream
			break
		}
	}
	return p.stream
}
