package main

import (
	"context"
	"time"
	"log"
	"google.golang.org/grpc"
	pb "github.com/diskordanz/coupon_service/proto"

)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewCouponServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	coupon:=pb.Coupon{
		Name:"coupon 2",
		Code: "CODE",
		Description: "",
		Type: 1,
		Status: true,
		Value: 10.0,
		FranchiseId:1 }

	request, err := client.CreateCoupon(ctx, &pb.CreateCouponRequest{Coupon:&coupon})
	if err != nil {
		log.Fatalf("could not create coupon: %v", err)
	}
	log.Printf("New coupon: %s", request)
 
	request2, err := client.ListCoupons(ctx, &pb.ListCouponsRequest{})
	if err != nil {
		log.Fatalf("could not get coupons list: %v", err)
	}
	log.Printf("All coupons: %s", request2)


}