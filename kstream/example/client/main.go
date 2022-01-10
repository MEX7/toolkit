package main

import (
	"fmt"
	"sync/atomic"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gotomicro/ego"
	"github.com/gotomicro/ego/core/elog"
	"github.com/gotomicro/ego/server/egin"
	"github.com/gotomicro/ego/server/egovernor"
	"google.golang.org/grpc"

	"github.com/kl7sn/toolkit/kstream"
	"github.com/kl7sn/toolkit/kstream/pb"
	"github.com/kl7sn/toolkit/xgo"
)

const (
	ADDRESS = "127.0.0.1:8080"
)

var (
	do   int64
	succ int64
)

func main() {
	if err := ego.New().
		Invoker(
			initStream,
		).
		Serve(
			egovernor.Load("server.governor").Build(),
			serverHttp(),
		).Run(); err != nil {
		elog.Panic("startup", elog.FieldErr(err))
	}
}

var streamProxy *kstream.ProxyStream

func initStream() error {
	conn, err := grpc.Dial(ADDRESS, grpc.WithInsecure())
	if err != nil {
		return err
	}
	c := pb.NewStreamClient(conn)
	streamProxy = kstream.InitStream(c)
	streamProxy.Writer(1)
	xgo.Go(func() {
		dump()
	})
	return nil
}

func dump() {
	for {
		fmt.Printf("do: %d, succ: %d\n", do, succ)
		time.Sleep(time.Second)
	}
}

func serverHttp() *egin.Component {
	server := egin.Load("server.http").Build()
	server.GET("/hello", func(ctx *gin.Context) {
		for i := 0; i < 1000; i++ {
			xgo.Go(func() {
				cube()
			})
		}
		ctx.JSON(200, "Hello")
		return
	})
	return server
}

func cube() {
	for {
		atomic.AddInt64(&do, 1)
		xgo.Go(func() {
			streamProxy.PushChan(&pb.CellReq{
				MsgId:   111,
				Msg:     []byte("hello,hello,hello,hello,hello,hello,hello,hello,hello,hello,hello,hello,hello,hello,hello,hello,hello,hello,hello,hello,hello,"),
				Headers: nil,
			})
			atomic.AddInt64(&succ, 1)
		})
		time.Sleep(time.Second)
	}
}

var p *kstream.ProxyStream

func callback(reply *pb.CellResp) {
	fmt.Println("reply", reply)
	p.PushChan(&pb.CellReq{
		MsgId: 333,
		Msg:   []byte(fmt.Sprintf("%s-%d", "hello", time.Now().Unix())),
	})

}
