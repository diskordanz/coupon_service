package main

import (
	
	pb "github.com/diskordanz/coupon_service/proto"
	"log"
	"net"
	"os"
	//"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	repository "github.com/diskordanz/coupon_service/repository"
	model "github.com/diskordanz/coupon_service/model"


	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)
 

func main() {
	//time.Sleep(5*time.Second)
	var err error
	var lis net.Listener
	if port := os.Getenv("COUPON_SERVICE_ADDRESS"); port != "" {
		lis, err = net.Listen("tcp", port)
	} else {
		lis, err = net.Listen("tcp", "localhost:50052")
	}

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=coupons3_db sslmode=disable password=postgres")

	if err != nil {
		log.Fatal(err)
	}

	pb.RegisterCouponServiceServer(s, &repository.CouponService{&model.GormDB{db}})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	} 
}
