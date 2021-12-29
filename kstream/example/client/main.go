package main

import (
	"fmt"
	"time"

	"google.golang.org/grpc"

	"github.com/kl7sn/toolkit/kstream"
	"github.com/kl7sn/toolkit/kstream/pb"
)

const (
	ADDRESS = "127.0.0.1:8080"
)

func main() {
	conn, err := grpc.Dial(ADDRESS, grpc.WithInsecure())
	if err != nil {
		panic(err)
		return
	}
	defer conn.Close()
	c := pb.NewStreamClient(conn)
	p = kstream.InitStream(c, callback, true)
	select {}
}

var p *kstream.ProxyStream

func callback(reply *pb.CellResp) {
	fmt.Println("reply", reply)
	p.PushChan(&pb.CellReq{
		MsgId: 333,
		Msg:   []byte(fmt.Sprintf("%s-%d", "hello", time.Now().Unix())),
	})

}
