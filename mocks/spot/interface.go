// Code generated by MockGen. DO NOT EDIT.
// Source: degate/spot/interface.go

// Package mock_spot is a generated GoMock package.
package mock_spot

import (
	reflect "reflect"

	binance "github.com/degatedev/degate-sdk-golang/degate/binance"
	model "github.com/degatedev/degate-sdk-golang/degate/model"
	gomock "github.com/golang/mock/gomock"
)

// MockSpotClient is a mock of SpotClient interface.
type MockSpotClient struct {
	ctrl     *gomock.Controller
	recorder *MockSpotClientMockRecorder
}

// MockSpotClientMockRecorder is the mock recorder for MockSpotClient.
type MockSpotClientMockRecorder struct {
	mock *MockSpotClient
}

// NewMockSpotClient creates a new mock instance.
func NewMockSpotClient(ctrl *gomock.Controller) *MockSpotClient {
	mock := &MockSpotClient{ctrl: ctrl}
	mock.recorder = &MockSpotClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSpotClient) EXPECT() *MockSpotClientMockRecorder {
	return m.recorder
}

// Account mocks base method.
func (m *MockSpotClient) Account() (*binance.AccountResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Account")
	ret0, _ := ret[0].(*binance.AccountResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Account indicates an expected call of Account.
func (mr *MockSpotClientMockRecorder) Account() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Account", reflect.TypeOf((*MockSpotClient)(nil).Account))
}

// BookTicker mocks base method.
func (m *MockSpotClient) BookTicker(param *model.BookTickerParam) (*binance.BookTickerResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BookTicker", param)
	ret0, _ := ret[0].(*binance.BookTickerResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BookTicker indicates an expected call of BookTicker.
func (mr *MockSpotClientMockRecorder) BookTicker(param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BookTicker", reflect.TypeOf((*MockSpotClient)(nil).BookTicker), param)
}

// CancelOpenOrders mocks base method.
func (m *MockSpotClient) CancelOpenOrders(includeGrid bool) (*binance.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CancelOpenOrders", includeGrid)
	ret0, _ := ret[0].(*binance.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CancelOpenOrders indicates an expected call of CancelOpenOrders.
func (mr *MockSpotClientMockRecorder) CancelOpenOrders(includeGrid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelOpenOrders", reflect.TypeOf((*MockSpotClient)(nil).CancelOpenOrders), includeGrid)
}

// CancelOrder mocks base method.
func (m *MockSpotClient) CancelOrder(param *model.CancelOrderParam) (*binance.OrderCancelResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CancelOrder", param)
	ret0, _ := ret[0].(*binance.OrderCancelResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CancelOrder indicates an expected call of CancelOrder.
func (mr *MockSpotClientMockRecorder) CancelOrder(param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelOrder", reflect.TypeOf((*MockSpotClient)(nil).CancelOrder), param)
}

// CancelOrderOnChain mocks base method.
func (m *MockSpotClient) CancelOrderOnChain(param *model.CancelOrderParam) (*binance.OrderCancelResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CancelOrderOnChain", param)
	ret0, _ := ret[0].(*binance.OrderCancelResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CancelOrderOnChain indicates an expected call of CancelOrderOnChain.
func (mr *MockSpotClientMockRecorder) CancelOrderOnChain(param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelOrderOnChain", reflect.TypeOf((*MockSpotClient)(nil).CancelOrderOnChain), param)
}

// DeleteListenKey mocks base method.
func (m *MockSpotClient) DeleteListenKey(param *model.ListenKeyParam) (*binance.EmptyResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteListenKey", param)
	ret0, _ := ret[0].(*binance.EmptyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteListenKey indicates an expected call of DeleteListenKey.
func (mr *MockSpotClientMockRecorder) DeleteListenKey(param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteListenKey", reflect.TypeOf((*MockSpotClient)(nil).DeleteListenKey), param)
}

// DepositHistory mocks base method.
func (m *MockSpotClient) DepositHistory(param *model.DepositsParam) (*binance.DepositHistoryResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DepositHistory", param)
	ret0, _ := ret[0].(*binance.DepositHistoryResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DepositHistory indicates an expected call of DepositHistory.
func (mr *MockSpotClientMockRecorder) DepositHistory(param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DepositHistory", reflect.TypeOf((*MockSpotClient)(nil).DepositHistory), param)
}

// Depth mocks base method.
func (m *MockSpotClient) Depth(param *model.DepthParam) (*binance.DepthResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Depth", param)
	ret0, _ := ret[0].(*binance.DepthResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Depth indicates an expected call of Depth.
func (mr *MockSpotClientMockRecorder) Depth(param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Depth", reflect.TypeOf((*MockSpotClient)(nil).Depth), param)
}

// ExchangeInfo mocks base method.
func (m *MockSpotClient) ExchangeInfo() (*binance.ExchangeInfoResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExchangeInfo")
	ret0, _ := ret[0].(*binance.ExchangeInfoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExchangeInfo indicates an expected call of ExchangeInfo.
func (mr *MockSpotClientMockRecorder) ExchangeInfo() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExchangeInfo", reflect.TypeOf((*MockSpotClient)(nil).ExchangeInfo))
}

// GasFee mocks base method.
func (m *MockSpotClient) GasFee() (*binance.GasFeeTokenResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GasFee")
	ret0, _ := ret[0].(*binance.GasFeeTokenResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GasFee indicates an expected call of GasFee.
func (mr *MockSpotClientMockRecorder) GasFee() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GasFee", reflect.TypeOf((*MockSpotClient)(nil).GasFee))
}

// GetBalance mocks base method.
func (m *MockSpotClient) GetBalance(param *model.AccountBalanceParam) (*binance.BalanceResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBalance", param)
	ret0, _ := ret[0].(*binance.BalanceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBalance indicates an expected call of GetBalance.
func (mr *MockSpotClientMockRecorder) GetBalance(param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBalance", reflect.TypeOf((*MockSpotClient)(nil).GetBalance), param)
}

// GetHistoryOrders mocks base method.
func (m *MockSpotClient) GetHistoryOrders(param *model.OrdersParam) (*binance.OrdersResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHistoryOrders", param)
	ret0, _ := ret[0].(*binance.OrdersResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHistoryOrders indicates an expected call of GetHistoryOrders.
func (mr *MockSpotClientMockRecorder) GetHistoryOrders(param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHistoryOrders", reflect.TypeOf((*MockSpotClient)(nil).GetHistoryOrders), param)
}

// GetOpenOrders mocks base method.
func (m *MockSpotClient) GetOpenOrders(param *model.OrdersParam) (*binance.OrdersResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOpenOrders", param)
	ret0, _ := ret[0].(*binance.OrdersResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOpenOrders indicates an expected call of GetOpenOrders.
func (mr *MockSpotClientMockRecorder) GetOpenOrders(param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOpenOrders", reflect.TypeOf((*MockSpotClient)(nil).GetOpenOrders), param)
}

// GetOrder mocks base method.
func (m *MockSpotClient) GetOrder(param *model.OrderDetailParam) (*model.OrderDetailResponse, *binance.OrderResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrder", param)
	ret0, _ := ret[0].(*model.OrderDetailResponse)
	ret1, _ := ret[1].(*binance.OrderResponse)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetOrder indicates an expected call of GetOrder.
func (mr *MockSpotClientMockRecorder) GetOrder(param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrder", reflect.TypeOf((*MockSpotClient)(nil).GetOrder), param)
}

// GetTradeFee mocks base method.
func (m *MockSpotClient) GetTradeFee(param *model.TradeFeeParam) (*binance.TradeFeeResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTradeFee", param)
	ret0, _ := ret[0].(*binance.TradeFeeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTradeFee indicates an expected call of GetTradeFee.
func (mr *MockSpotClientMockRecorder) GetTradeFee(param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTradeFee", reflect.TypeOf((*MockSpotClient)(nil).GetTradeFee), param)
}

// Klines mocks base method.
func (m *MockSpotClient) Klines(param *model.KlineParam) (*binance.KlineResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Klines", param)
	ret0, _ := ret[0].(*binance.KlineResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Klines indicates an expected call of Klines.
func (mr *MockSpotClientMockRecorder) Klines(param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Klines", reflect.TypeOf((*MockSpotClient)(nil).Klines), param)
}

// MyTrades mocks base method.
func (m *MockSpotClient) MyTrades(param *model.AccountTradesParam) (*binance.TradeResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MyTrades", param)
	ret0, _ := ret[0].(*binance.TradeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MyTrades indicates an expected call of MyTrades.
func (mr *MockSpotClientMockRecorder) MyTrades(param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MyTrades", reflect.TypeOf((*MockSpotClient)(nil).MyTrades), param)
}

// NewListenKey mocks base method.
func (m *MockSpotClient) NewListenKey() (*binance.ListenKeyResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListenKey")
	ret0, _ := ret[0].(*binance.ListenKeyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewListenKey indicates an expected call of NewListenKey.
func (mr *MockSpotClientMockRecorder) NewListenKey() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListenKey", reflect.TypeOf((*MockSpotClient)(nil).NewListenKey))
}

// NewOrder mocks base method.
func (m *MockSpotClient) NewOrder(param *model.OrderParam) (*binance.NewOrderResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewOrder", param)
	ret0, _ := ret[0].(*binance.NewOrderResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewOrder indicates an expected call of NewOrder.
func (mr *MockSpotClientMockRecorder) NewOrder(param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewOrder", reflect.TypeOf((*MockSpotClient)(nil).NewOrder), param)
}

// Ping mocks base method.
func (m *MockSpotClient) Ping() (*binance.PingResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ping")
	ret0, _ := ret[0].(*binance.PingResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Ping indicates an expected call of Ping.
func (mr *MockSpotClientMockRecorder) Ping() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*MockSpotClient)(nil).Ping))
}

// ReNewListenKey mocks base method.
func (m *MockSpotClient) ReNewListenKey(param *model.ListenKeyParam) (*binance.EmptyResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReNewListenKey", param)
	ret0, _ := ret[0].(*binance.EmptyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReNewListenKey indicates an expected call of ReNewListenKey.
func (mr *MockSpotClientMockRecorder) ReNewListenKey(param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReNewListenKey", reflect.TypeOf((*MockSpotClient)(nil).ReNewListenKey), param)
}

// Ticker24 mocks base method.
func (m *MockSpotClient) Ticker24(param *model.TickerParam) (*binance.TickerResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ticker24", param)
	ret0, _ := ret[0].(*binance.TickerResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Ticker24 indicates an expected call of Ticker24.
func (mr *MockSpotClientMockRecorder) Ticker24(param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ticker24", reflect.TypeOf((*MockSpotClient)(nil).Ticker24), param)
}

// TickerPrice mocks base method.
func (m *MockSpotClient) TickerPrice(param *model.PairPriceParam) (*binance.PairPriceResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TickerPrice", param)
	ret0, _ := ret[0].(*binance.PairPriceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TickerPrice indicates an expected call of TickerPrice.
func (mr *MockSpotClientMockRecorder) TickerPrice(param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TickerPrice", reflect.TypeOf((*MockSpotClient)(nil).TickerPrice), param)
}

// Time mocks base method.
func (m *MockSpotClient) Time() (*binance.TimeResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Time")
	ret0, _ := ret[0].(*binance.TimeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Time indicates an expected call of Time.
func (mr *MockSpotClientMockRecorder) Time() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Time", reflect.TypeOf((*MockSpotClient)(nil).Time))
}

// TokenList mocks base method.
func (m *MockSpotClient) TokenList(param *model.TokenListParam) (*model.TokensResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TokenList", param)
	ret0, _ := ret[0].(*model.TokensResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TokenList indicates an expected call of TokenList.
func (mr *MockSpotClientMockRecorder) TokenList(param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TokenList", reflect.TypeOf((*MockSpotClient)(nil).TokenList), param)
}

// Trades mocks base method.
func (m *MockSpotClient) Trades(param *model.TradeLastedParam) (*binance.TradeHistoryResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trades", param)
	ret0, _ := ret[0].(*binance.TradeHistoryResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Trades indicates an expected call of Trades.
func (mr *MockSpotClientMockRecorder) Trades(param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trades", reflect.TypeOf((*MockSpotClient)(nil).Trades), param)
}

// TradesHistory mocks base method.
func (m *MockSpotClient) TradesHistory(param *model.TradeHistoryParam) (*binance.TradeHistoryResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TradesHistory", param)
	ret0, _ := ret[0].(*binance.TradeHistoryResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TradesHistory indicates an expected call of TradesHistory.
func (mr *MockSpotClientMockRecorder) TradesHistory(param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TradesHistory", reflect.TypeOf((*MockSpotClient)(nil).TradesHistory), param)
}

// Transfer mocks base method.
func (m *MockSpotClient) Transfer(param *model.TransferParam) (*binance.TransferResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Transfer", param)
	ret0, _ := ret[0].(*binance.TransferResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Transfer indicates an expected call of Transfer.
func (mr *MockSpotClientMockRecorder) Transfer(param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Transfer", reflect.TypeOf((*MockSpotClient)(nil).Transfer), param)
}

// TransferHistory mocks base method.
func (m *MockSpotClient) TransferHistory(param *model.TransfersParam) (*binance.TransferHistoryResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransferHistory", param)
	ret0, _ := ret[0].(*binance.TransferHistoryResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TransferHistory indicates an expected call of TransferHistory.
func (mr *MockSpotClientMockRecorder) TransferHistory(param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransferHistory", reflect.TypeOf((*MockSpotClient)(nil).TransferHistory), param)
}

// Withdraw mocks base method.
func (m *MockSpotClient) Withdraw(param *model.WithdrawParam) (*binance.WithdrawResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Withdraw", param)
	ret0, _ := ret[0].(*binance.WithdrawResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Withdraw indicates an expected call of Withdraw.
func (mr *MockSpotClientMockRecorder) Withdraw(param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Withdraw", reflect.TypeOf((*MockSpotClient)(nil).Withdraw), param)
}

// WithdrawHistory mocks base method.
func (m *MockSpotClient) WithdrawHistory(param *model.WithdrawsParam) (*binance.WithdrawHistoryResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithdrawHistory", param)
	ret0, _ := ret[0].(*binance.WithdrawHistoryResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WithdrawHistory indicates an expected call of WithdrawHistory.
func (mr *MockSpotClientMockRecorder) WithdrawHistory(param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithdrawHistory", reflect.TypeOf((*MockSpotClient)(nil).WithdrawHistory), param)
}