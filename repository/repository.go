package repository

import(
	"context"
	pb "github.com/diskordanz/coupon_service/proto"
	"github.com/golang/protobuf/ptypes/empty"

)

type IRepository interface{
	CreateCoupon(*pb.CreateCouponRequest) (*pb.Coupon, error)
    GetCoupon(*pb.GetCouponRequest) (*pb.Coupon, error)
    UpdateCoupon(*pb.UpdateCouponRequest) (*pb.Coupon, error)
    DeleteCoupon(*pb.DeleteCouponRequest) error
    ListCoupons(*pb.ListCouponsRequest) (*pb.ListCouponsResponse, error)
    ListCouponsByFranchise(*pb.ListCouponsByFranchiseRequest)(*pb.ListCouponsResponse, error)
}

type CouponService struct {
	Repo IRepository
}

func (s *CouponService) CreateCoupon(ctx context.Context, req *pb.CreateCouponRequest) (*pb.Coupon, error) {
	res, err := s.Repo.CreateCoupon(req); if err != nil {
		return nil, err
	}
	return res, nil
} 

func (s *CouponService) GetCoupon(ctx context.Context, req *pb.GetCouponRequest) (*pb.Coupon, error) {
	res, err := s.Repo.GetCoupon(req); if err != nil {
		return nil, err
	}
	return res, nil
}
 
func (s *CouponService) ListCoupons(ctx context.Context, req *pb.ListCouponsRequest) (*pb.ListCouponsResponse, error) {
	res, err := s.Repo.ListCoupons(req); if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *CouponService) UpdateCoupon(ctx context.Context, req *pb.UpdateCouponRequest) (*pb.Coupon, error) {
	res, err := s.Repo.UpdateCoupon(req); if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *CouponService) DeleteCoupon(ctx context.Context, req *pb.DeleteCouponRequest) (*empty.Empty, error) {
	err := s.Repo.DeleteCoupon(req); if err != nil {
		return &empty.Empty{}, err
	}
	return &empty.Empty{}, nil
}

func (s *CouponService) ListCouponsByFranchise(ctx context.Context, req *pb.ListCouponsByFranchiseRequest) (*pb.ListCouponsResponse, error) {
	res, err := s.Repo.ListCouponsByFranchise(req); if err != nil {
		return nil, err
	}
	return res, nil
}