package main

import (
	"context"
	"log"
	"net"
	"time"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"github.com/go-pg/pg" 
	"github.com/golang/protobuf/ptypes/empty"
	pb "github.com/diskordanz/coupon_service/proto"
)

type CouponService struct{
	DB *pg.DB
}

func (s CouponService) GetCoupon(ctx context.Context, req *pb.GetCouponRequest) (*pb.Coupon, error) {
	var coupon pb.Coupon
	err:=s.DB.Model(&coupon).Where("id = ?", req.Id).First()
	if err != nil {
		return nil, err
	}

	return &coupon, nil
}

func (s CouponService) ListCoupons(ctx context.Context, req *pb.ListCouponsRequest) (*pb.ListCouponsResponse, error) {
	var coupons []*pb.Coupon
	err:=s.DB.Model(&coupons).Select()
	if err != nil {
		return nil, err
	}

	return &pb.ListCouponsResponse{Coupons: coupons}, nil
}

func (s CouponService) CreateCoupon(ctx context.Context, req *pb.CreateCouponRequest) (*pb.Coupon, error) {
	err:=s.DB.Insert(req.Coupon)
	if err != nil {
		return nil, err
	}
	return &pb.Coupon{Id: req.Coupon.Id}, nil
}

func (s CouponService) UpdateCoupon(ctx context.Context, req *pb.UpdateCouponRequest) (*pb.Coupon, error) {
	err := s.DB.Model(req.Coupon).Column("name","code","description","type","status","time_from","time_to","date_from","date_to","days","value","franchise_id").Update()
	if err != nil {
		return nil, err
	}
	return &pb.Coupon{}, nil
}

func (s CouponService) DeleteCoupon(ctx context.Context, req *pb.DeleteCouponRequest) (*empty.Empty, error) {
	err:=s.DB.Delete(&pb.Coupon{Id:req.Id})
	if err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	db := pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "postgres",
		Database: "coupons_db",
		Addr:     "5555:5555",
		RetryStatementTimeout: true,
		MaxRetries:            4,
		MinRetryBackoff:       250 * time.Millisecond,
	})
	db.CreateTable(&pb.Coupon{}, nil)

	pb.RegisterCouponServiceServer(s, &CouponService{DB:db})

	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}