package data

import (
	"sync"

	"github.com/ygongq/grpc-demo/grpc-kindle/02_orderManagement/server/proto"
)

type DataMap struct {
	m  map[string]*proto.Order
	rw sync.RWMutex
}

func NewDataMap() DataMap {
	return DataMap{m: make(map[string]*proto.Order), rw: sync.RWMutex{}}
}

func (m *DataMap) Set(order *proto.Order) {
	m.rw.Lock()
	defer m.rw.Unlock()
	m.m[order.Id] = order
}

func (m *DataMap) Get(orderID string) (*proto.Order, bool) {
	m.rw.RLock()
	defer m.rw.RUnlock()
	v, ok := m.m[orderID]
	return v, ok
}

func (m *DataMap) Del(orderID string) {
	m.rw.Lock()
	defer m.rw.Unlock()
	delete(m.m, orderID)
}

func (m *DataMap) List() map[string]*proto.Order {
	m.rw.Lock()
	defer m.rw.Unlock()
	return m.m
}
