package main

import (
	"context"
	"time"
	"log"
	"google.golang.org/grpc"
	pb "github.com/diskordanz/coupon_service/proto"
	//"github.com/golang/protobuf/ptypes/wrappers"


)

func main() {
	conn, err := grpc.Dial("localhost:9097", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewCouponServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	

	// coup:= pb.Coupon{
	//   	Name: &wrappers.StringValue{ Value: "TestCoupon"}, 
	//   	Code: &wrappers.StringValue{ Value: "TestCODE"},
	//   	Description: &wrappers.StringValue{ Value: "some info"},
	// 	Type: 0,  
	// 	Status: &wrappers.BoolValue{ Value: true},
	// 	TimeFrom: &pb.TimeOfDay{
	// 		Hours: &wrappers.UInt32Value{ Value: 13},
	// 		Minutes: &wrappers.UInt32Value{ Value:  1},
	// 	},
	// 	TimeTo: &pb.TimeOfDay{
	// 		Hours: &wrappers.UInt32Value{ Value: 13},
	// 		Minutes: &wrappers.UInt32Value{ Value: 59},
	// 	},
	// 	DateFrom: &pb.Date{
	// 		Year: &wrappers.UInt32Value{ Value: 2019},
	// 		Month: &wrappers.UInt32Value{ Value: 2},
	// 		Day:&wrappers.UInt32Value{ Value: 1},
	// 	},
	// 	DateTo: &pb.Date{
	// 		Year: &wrappers.UInt32Value{ Value: 2019},
	// 		Month: &wrappers.UInt32Value{ Value: 5},
	// 		Day:&wrappers.UInt32Value{ Value: 30},
	// 	},
	//   	Days: []pb.Coupon_DayOfWeek{1,2,3,4,5,6,7},
	//   	Value: &wrappers.FloatValue{ Value: 10.0},
	//   	FranchiseId: &wrappers.UInt64Value{ Value: 1},
	//  }

	//  _,err = client.CreateCoupon(ctx, &pb.CreateCouponRequest{Coupon: &coup})
	//  if err != nil {
	// 	log.Fatalf("could not crete coupon: %v", err)
	//  }
	

	 
	 res, err := client.ListCoupons(ctx, &pb.ListCouponsRequest{})
	 if err != nil {
		 log.Fatalf("could not get coupons: %v", err)
	 }
	 for _,k := range res.Coupons {
		 log.Println(k)
	 }
}