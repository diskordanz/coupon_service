package main

import (
	"context"
	"time"
	"log"
	"google.golang.org/grpc"
	pb "github.com/diskordanz/coupon_service/proto"

)

func main() {
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewCouponServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	request, err := client.GetCoupon(ctx, &pb.GetCouponRequest{Id:1})
	if err != nil {
		log.Fatalf("could not get coupons: %v", err)
	 }
	 log.Println(request)
	// coup:= pb.Coupon{
	// 	Name: "TestCoupon", 
	// 	Code: "TestCODE",
	// 	Description: "some info",
	// 	Status: false,
	// 	//Days: []pb.Coupon_DayOfWeek{1,5,6},
	// 	Value: 30.0,
	// 	FranchiseId: 2,
	// }
	// log.Println(coup)

	// request2, err := client.CreateCoupon(ctx, &pb.CreateCouponRequest{Coupon: &coup})
	// if err != nil {
	// 	log.Fatalf("could not crete coupon: %v", err)
	// }
	// log.Println(request2)
	

}