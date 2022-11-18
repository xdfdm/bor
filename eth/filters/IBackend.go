// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ethereum/go-ethereum/eth/filters (interfaces: Backend)

// Package filters is a generated GoMock package.
package filters

import (
	context "context"
	reflect "reflect"

	common "github.com/ethereum/go-ethereum/common"
	core "github.com/ethereum/go-ethereum/core"
	bloombits "github.com/ethereum/go-ethereum/core/bloombits"
	types "github.com/ethereum/go-ethereum/core/types"
	ethdb "github.com/ethereum/go-ethereum/ethdb"
	event "github.com/ethereum/go-ethereum/event"
	rpc "github.com/ethereum/go-ethereum/rpc"
	gomock "github.com/golang/mock/gomock"
)

// MockBackend is a mock of Backend interface.
type MockBackend struct {
	ctrl     *gomock.Controller
	recorder *MockBackendMockRecorder
}

// MockBackendMockRecorder is the mock recorder for MockBackend.
type MockBackendMockRecorder struct {
	mock *MockBackend
}

// NewMockBackend creates a new mock instance.
func NewMockBackend(ctrl *gomock.Controller) *MockBackend {
	mock := &MockBackend{ctrl: ctrl}
	mock.recorder = &MockBackendMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBackend) EXPECT() *MockBackendMockRecorder {
	return m.recorder
}

// BloomStatus mocks base method.
func (m *MockBackend) BloomStatus() (uint64, uint64) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BloomStatus")
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(uint64)
	return ret0, ret1
}

// BloomStatus indicates an expected call of BloomStatus.
func (mr *MockBackendMockRecorder) BloomStatus() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BloomStatus", reflect.TypeOf((*MockBackend)(nil).BloomStatus))
}

// ChainDb mocks base method.
func (m *MockBackend) ChainDb() ethdb.Database {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChainDb")
	ret0, _ := ret[0].(ethdb.Database)
	return ret0
}

// ChainDb indicates an expected call of ChainDb.
func (mr *MockBackendMockRecorder) ChainDb() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChainDb", reflect.TypeOf((*MockBackend)(nil).ChainDb))
}

// GetBorBlockLogs mocks base method.
func (m *MockBackend) GetBorBlockLogs(arg0 context.Context, arg1 common.Hash) ([]*types.Log, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBorBlockLogs", arg0, arg1)
	ret0, _ := ret[0].([]*types.Log)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBorBlockLogs indicates an expected call of GetBorBlockLogs.
func (mr *MockBackendMockRecorder) GetBorBlockLogs(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBorBlockLogs", reflect.TypeOf((*MockBackend)(nil).GetBorBlockLogs), arg0, arg1)
}

// GetBorBlockReceipt mocks base method.
func (m *MockBackend) GetBorBlockReceipt(arg0 context.Context, arg1 common.Hash) (*types.Receipt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBorBlockReceipt", arg0, arg1)
	ret0, _ := ret[0].(*types.Receipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBorBlockReceipt indicates an expected call of GetBorBlockReceipt.
func (mr *MockBackendMockRecorder) GetBorBlockReceipt(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBorBlockReceipt", reflect.TypeOf((*MockBackend)(nil).GetBorBlockReceipt), arg0, arg1)
}

// GetLogs mocks base method.
func (m *MockBackend) GetLogs(arg0 context.Context, arg1 common.Hash) ([][]*types.Log, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLogs", arg0, arg1)
	ret0, _ := ret[0].([][]*types.Log)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLogs indicates an expected call of GetLogs.
func (mr *MockBackendMockRecorder) GetLogs(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLogs", reflect.TypeOf((*MockBackend)(nil).GetLogs), arg0, arg1)
}

// GetReceipts mocks base method.
func (m *MockBackend) GetReceipts(arg0 context.Context, arg1 common.Hash) (types.Receipts, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReceipts", arg0, arg1)
	ret0, _ := ret[0].(types.Receipts)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReceipts indicates an expected call of GetReceipts.
func (mr *MockBackendMockRecorder) GetReceipts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReceipts", reflect.TypeOf((*MockBackend)(nil).GetReceipts), arg0, arg1)
}

// HeaderByHash mocks base method.
func (m *MockBackend) HeaderByHash(arg0 context.Context, arg1 common.Hash) (*types.Header, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HeaderByHash", arg0, arg1)
	ret0, _ := ret[0].(*types.Header)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HeaderByHash indicates an expected call of HeaderByHash.
func (mr *MockBackendMockRecorder) HeaderByHash(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HeaderByHash", reflect.TypeOf((*MockBackend)(nil).HeaderByHash), arg0, arg1)
}

// HeaderByNumber mocks base method.
func (m *MockBackend) HeaderByNumber(arg0 context.Context, arg1 rpc.BlockNumber) (*types.Header, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HeaderByNumber", arg0, arg1)
	ret0, _ := ret[0].(*types.Header)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HeaderByNumber indicates an expected call of HeaderByNumber.
func (mr *MockBackendMockRecorder) HeaderByNumber(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HeaderByNumber", reflect.TypeOf((*MockBackend)(nil).HeaderByNumber), arg0, arg1)
}

// ServiceFilter mocks base method.
func (m *MockBackend) ServiceFilter(arg0 context.Context, arg1 *bloombits.MatcherSession) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ServiceFilter", arg0, arg1)
}

// ServiceFilter indicates an expected call of ServiceFilter.
func (mr *MockBackendMockRecorder) ServiceFilter(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServiceFilter", reflect.TypeOf((*MockBackend)(nil).ServiceFilter), arg0, arg1)
}

// SubscribeChainEvent mocks base method.
func (m *MockBackend) SubscribeChainEvent(arg0 chan<- core.ChainEvent) event.Subscription {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscribeChainEvent", arg0)
	ret0, _ := ret[0].(event.Subscription)
	return ret0
}

// SubscribeChainEvent indicates an expected call of SubscribeChainEvent.
func (mr *MockBackendMockRecorder) SubscribeChainEvent(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeChainEvent", reflect.TypeOf((*MockBackend)(nil).SubscribeChainEvent), arg0)
}

// SubscribeLogsEvent mocks base method.
func (m *MockBackend) SubscribeLogsEvent(arg0 chan<- []*types.Log) event.Subscription {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscribeLogsEvent", arg0)
	ret0, _ := ret[0].(event.Subscription)
	return ret0
}

// SubscribeLogsEvent indicates an expected call of SubscribeLogsEvent.
func (mr *MockBackendMockRecorder) SubscribeLogsEvent(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeLogsEvent", reflect.TypeOf((*MockBackend)(nil).SubscribeLogsEvent), arg0)
}

// SubscribeNewTxsEvent mocks base method.
func (m *MockBackend) SubscribeNewTxsEvent(arg0 chan<- core.NewTxsEvent) event.Subscription {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscribeNewTxsEvent", arg0)
	ret0, _ := ret[0].(event.Subscription)
	return ret0
}

// SubscribeNewTxsEvent indicates an expected call of SubscribeNewTxsEvent.
func (mr *MockBackendMockRecorder) SubscribeNewTxsEvent(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeNewTxsEvent", reflect.TypeOf((*MockBackend)(nil).SubscribeNewTxsEvent), arg0)
}

// SubscribePendingLogsEvent mocks base method.
func (m *MockBackend) SubscribePendingLogsEvent(arg0 chan<- []*types.Log) event.Subscription {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscribePendingLogsEvent", arg0)
	ret0, _ := ret[0].(event.Subscription)
	return ret0
}

// SubscribePendingLogsEvent indicates an expected call of SubscribePendingLogsEvent.
func (mr *MockBackendMockRecorder) SubscribePendingLogsEvent(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribePendingLogsEvent", reflect.TypeOf((*MockBackend)(nil).SubscribePendingLogsEvent), arg0)
}

// SubscribeRemovedLogsEvent mocks base method.
func (m *MockBackend) SubscribeRemovedLogsEvent(arg0 chan<- core.RemovedLogsEvent) event.Subscription {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscribeRemovedLogsEvent", arg0)
	ret0, _ := ret[0].(event.Subscription)
	return ret0
}

// SubscribeRemovedLogsEvent indicates an expected call of SubscribeRemovedLogsEvent.
func (mr *MockBackendMockRecorder) SubscribeRemovedLogsEvent(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeRemovedLogsEvent", reflect.TypeOf((*MockBackend)(nil).SubscribeRemovedLogsEvent), arg0)
}

// SubscribeStateSyncEvent mocks base method.
func (m *MockBackend) SubscribeStateSyncEvent(arg0 chan<- core.StateSyncEvent) event.Subscription {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscribeStateSyncEvent", arg0)
	ret0, _ := ret[0].(event.Subscription)
	return ret0
}

// SubscribeStateSyncEvent indicates an expected call of SubscribeStateSyncEvent.
func (mr *MockBackendMockRecorder) SubscribeStateSyncEvent(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeStateSyncEvent", reflect.TypeOf((*MockBackend)(nil).SubscribeStateSyncEvent), arg0)
}