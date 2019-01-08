package mock_coupons

import (
	gomock "github.com/golang/mock/gomock"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
	pb "github.com/diskordanz/coupon_service/proto"
	"github.com/golang/protobuf/ptypes/empty"

)

type MockCouponServiceClient struct {
	ctrl     *gomock.Controller
	recorder *_MockCouponServiceClientRecorder
}

type _MockCouponServiceClientRecorder struct {
	mock *MockCouponServiceClient
} 

func NewMockCouponServiceClient(ctrl *gomock.Controller) *MockCouponServiceClient {
	mock := &MockCouponServiceClient{ctrl: ctrl}
	mock.recorder = &_MockCouponServiceClientRecorder{mock}
	return mock
}

func (_m *MockCouponServiceClient) EXPECT() *_MockCouponServiceClientRecorder {
	return _m.recorder
}

func (_m *MockCouponServiceClient) GetCoupon(_param0 context.Context, _param1 *pb.GetCouponRequest, _param2 ...grpc.CallOption) (*pb.Coupon, error) {
	_s := []interface{}{_param0, _param1}
	for _, _x := range _param2 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "GetCoupon", _s...)
	ret0, _ := ret[0].(*pb.Coupon)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockCouponServiceClientRecorder) GetCoupon(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetCoupon", _s...)
}

func (_m *MockCouponServiceClient) ListCoupons(_param0 context.Context, _param1 *pb.ListCouponsRequest, _param2 ...grpc.CallOption) (*pb.ListCouponsResponse, error) {
	_s := []interface{}{_param0, _param1}
	for _, _x := range _param2 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "ListCoupons", _s...)
	ret0, _ := ret[0].(*pb.ListCouponsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockCouponServiceClientRecorder) ListCoupons(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListCoupons", _s...)
}


func (_m *MockCouponServiceClient) CreateCoupon(_param0 context.Context, _param1 *pb.CreateCouponRequest, _param2 ...grpc.CallOption) (*pb.Coupon, error) {
	_s := []interface{}{_param0, _param1}
	for _, _x := range _param2 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "CreateCoupon", _s...)
	ret0, _ := ret[0].(*pb.Coupon)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockCouponServiceClientRecorder) CreateCoupon(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateCoupon", _s...)
}

func (_m *MockCouponServiceClient) UpdateCoupon(_param0 context.Context, _param1 *pb.UpdateCouponRequest, _param2 ...grpc.CallOption) (*pb.Coupon, error) {
	_s := []interface{}{_param0, _param1}
	for _, _x := range _param2 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "UpdateCoupon", _s...)
	ret0, _ := ret[0].(*pb.Coupon)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockCouponServiceClientRecorder) UpdateCoupon(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UpdateCoupon", _s...)
}

func (_m *MockCouponServiceClient) DeleteCoupon(_param0 context.Context, _param1 *pb.DeleteCouponRequest, _param2 ...grpc.CallOption) (*empty.Empty, error) {
	_s := []interface{}{_param0, _param1}
	for _, _x := range _param2 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "DeleteCoupon", _s...)
	ret0, _ := ret[0].(*empty.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockCouponServiceClientRecorder) DeleteCoupon(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteCoupon", _s...)
}

func (_m *MockCouponServiceClient) ListCouponsByFranchise(_param0 context.Context, _param1 *pb.ListCouponsByFranchiseRequest, _param2 ...grpc.CallOption) (*pb.ListCouponsResponse, error) {
	_s := []interface{}{_param0, _param1}
	for _, _x := range _param2 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "ListCouponsByFranchise", _s...)
	ret0, _ := ret[0].(*pb.ListCouponsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockCouponServiceClientRecorder) ListCouponsByFranchise(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListCouponsByFranchise", _s...)
}
