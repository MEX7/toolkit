package kstream

import (
	"sync"

	"github.com/gotomicro/ego/core/elog"

	"github.com/kl7sn/toolkit/kstream/pb"
)

var StreamStore *stores

type stores struct {
	sync.RWMutex
	data map[string]pb.Stream_CellServer
}

func InitStreamStores() {
	StreamStore = &stores{
		data: make(map[string]pb.Stream_CellServer),
	}
}

func (s *stores) AddStream(ip string, stream pb.Stream_CellServer) {
	s.Lock()
	defer s.Unlock()
	s.data[ip] = stream
	elog.Debug("stores", elog.String("ip", ip), elog.String("step", "AddStream"))
}

func (s *stores) GetStream(ip string) pb.Stream_CellServer {
	s.RLock()
	defer s.RUnlock()
	if _, ok := s.data[ip]; ok {
		return s.data[ip]
	}
	return nil
}

func (s *stores) Broadcast(data *pb.CellResp) error {
	s.RLock()
	defer s.RUnlock()
	for _, stream := range s.data {
		if err := stream.Send(data); err != nil {
			return err
		}
	}
	return nil
}

func (s *stores) DeleteStream(ip string) {
	s.Lock()
	defer s.Unlock()
	elog.Debug("stores", elog.String("ip", ip), elog.String("step", "DeleteStream"))
	delete(s.data, ip)
}

func (s *stores) IsStreamExist(ip string) bool {
	s.RLock()
	defer s.RUnlock()
	if _, ok := s.data[ip]; ok {
		return true
	}
	return false
}
