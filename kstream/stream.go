package kstream

import (
	"context"
	"fmt"
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

type ProxyStream struct {
	sync.Mutex
	stream  pb.Stream_CellClient
	msgChan chan *pb.CellReq
}

func InitStream(client pb.StreamClient, callback Callback, isRead bool) *ProxyStream {
	obj := &ProxyStream{msgChan: make(chan *pb.CellReq, 1000)}
	obj.PushChan(&pb.CellReq{MsgId: MsgConnectReq})
	obj.sync(client, callback, isRead)
	return obj
}

func (p *ProxyStream) PushChan(info *pb.CellReq) {
	p.msgChan <- info
}
func (p *ProxyStream) sync(client pb.StreamClient, callback Callback, isRead bool) {
	if isRead {
		xgo.Go(func() {
			var (
				reply *pb.CellResp
				err   error
			)
			for {
				reply, err = p.GetStream(client).Recv()
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

	xgo.Go(func() {
		var err error
		for {
			data := <-p.msgChan
			err = p.GetStream(client).Send(data)
			if err != nil {
				fmt.Println("Send err:", err.Error())
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
			continue
		}
		p.stream = stream
		break
	}
	return p.stream
}
