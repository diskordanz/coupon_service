package model

import(
	"time"
	"fmt"
	"log"
	"github.com/jinzhu/gorm"
	pb "github.com/diskordanz/coupon_service/proto"
	_ "github.com/lib/pq"
	ptypes "github.com/golang/protobuf/ptypes"

)

type Coupon struct{ //сделать перегонку, переделать методы create, update, переписать базу, добавить пагинацию, парсер даты и времени, проверка дней
	Id uint64
	Name string
	Code string
	Description string
	Type int32
	Status bool
	Time_from time.Time
	Time_to time.Time
	Date_from time.Time 
	Date_to time.Time
	Days []int32
	Value float32
	FranchiseId uint64
}

type GormDB struct {
	DB *gorm.DB
}

func ToORM(c *pb.Coupon) (*Coupon ,error){
	coupon := &Coupon{}
	var err error

	coupon.Name = c.GetName()
	coupon.Code = c.GetCode()
	coupon.Description = c.GetDescription()
	coupon.Type = int32(c.GetType())
	coupon.Status = c.GetStatus()
	
	if coupon.Time_from, err = ptypes.Timestamp(c.GetTimeFrom()); err != nil {
		return coupon, err
	}
	if coupon.Time_to, err = ptypes.Timestamp(c.GetTimeTo()); err != nil {
		return coupon, err
	}
	if coupon.Date_from, err = ptypes.Timestamp(c.GetDateFrom()); err != nil {
		return coupon, err
	}
	if coupon.Date_to, err = ptypes.Timestamp(c.GetDateTo()); err != nil {
		return coupon, err
	}
	
	coupon.Days = DaysToORM(c.GetDays())
	coupon.Value = c.GetValue()
	coupon.FranchiseId = c.GetFranchiseId()

	return coupon, nil
}

func ToPB(c *Coupon) (*pb.Coupon,error){
	coupon := &pb.Coupon{}
	var err error

	coupon.Name = c.Name
	coupon.Code = c.Code
	coupon.Description = c.Description
	coupon.Type = pb.Coupon_CouponType(c.Type)
	coupon.Status = c.Status
	
	if coupon.TimeFrom, err = ptypes.TimestampProto(c.Time_from); err != nil {
		log.Fatalf("Error with TimeFrom: %v", err)
		return coupon, err
	}
	if coupon.TimeTo, err = ptypes.TimestampProto(c.Time_to); err != nil {
		log.Fatalf("Error with TimeTo: %v", err)
		return coupon, err
	}
	if coupon.DateFrom, err = ptypes.TimestampProto(c.Date_from); err != nil {
		log.Fatalf("Error with DateFrom: %v", err)
		return coupon, err
	}
	if coupon.DateTo, err = ptypes.TimestampProto(c.Date_to); err != nil {
		log.Fatalf("Error with DateTo: %v", err)
		return coupon, err
	}
	
	coupon.Days = DaysToPB(c.Days)
	coupon.Value = c.Value
	coupon.FranchiseId = c.FranchiseId

	log.Printf("ToPB: %v", coupon)

	return coupon, nil
}

func DaysToORM(d []pb.Coupon_DayOfWeek) []int32{

	var days []int32
	for _,k := range d {
		days = append(days, int32(k))
	}
	return days
}
func DaysToPB(d []int32) []pb.Coupon_DayOfWeek{

	var days []pb.Coupon_DayOfWeek
	for _,k := range d {
		days = append(days, pb.Coupon_DayOfWeek(k))
	}
	return days
}

func (db *GormDB) GetCoupon(req *pb.GetCouponRequest) (*pb.Coupon, error) {
	var coupon Coupon
	log.Println("Get coupon with id:", req.Id)
	if err := db.DB.First(&coupon, req.Id).Error; err != nil {
		log.Fatalf("Error with GetCoupon in gorm.go: %v", err)
		return nil, err
	}
	return ToPB(&coupon)
}

func (db *GormDB) ListCoupons(req *pb.ListCouponsRequest) (*pb.ListCouponsResponse, error) {
	coupons:= []Coupon{}
	coupon := &pb.Coupon{}
	var err error
	res := &pb.ListCouponsResponse{}
	if req.Filter == "" {
		if err = db.DB.Find(&coupons).Error; err != nil {
			log.Fatalf("Error with ListCoupons in gorm.go when Find: %v", err)
			return nil, err
		}
	} else {
		filter := fmt.Sprintf("%%%s%%", req.Filter)
		if err = db.DB.Where("name LIKE ?", filter).Find(&coupons).Error; err != nil {
			return nil, err
		}
	}

	for _, c := range coupons {
		if coupon, err = ToPB(&c); err != nil{
			log.Fatalf("Error with ListCoupons in gorm.go when ToPB: %v", err)
			return nil, err
		} 
		res.Coupons = append(res.Coupons, coupon)
	}

	return res, nil
}

func (db *GormDB) CreateCoupon(req *pb.CreateCouponRequest) (*pb.Coupon, error) {
	coupon := &pb.Coupon{Id: req.Coupon.Id}
	if err := db.DB.Create(req.Coupon).Error; err != nil {
		return nil, err
	}
	return coupon, nil
}

func (db *GormDB) UpdateCoupon(req *pb.UpdateCouponRequest) (*pb.Coupon, error) {
	if err := db.DB.Update(&req.Coupon).Error; err != nil {
		return nil, err
	}
	return req.Coupon, nil
}

func (db *GormDB) DeleteCoupon(req *pb.DeleteCouponRequest) error {
	if err := db.DB.Delete(&pb.Coupon{Id: req.Id}).Error; err != nil {
		return  err
	}
	return nil
}

func (db *GormDB) ListCouponsByFranchise(req *pb.ListCouponsByFranchiseRequest) (*pb.ListCouponsResponse, error) {
	var coupons []*pb.Coupon
	
	if req.Filter == "" {
		if err := db.DB.Where("franchise_id = ?", req.Id).Find(&coupons).Error; err != nil {
			return nil, err
		}
	} else {
		filter := fmt.Sprintf("%%%s%%", req.GetFilter())
		if err := db.DB.Where("franchise_id = ? AND name LIKE ?", req.Id,filter).Find(&coupons).Error; err != nil {
			return nil, err
		}
	}

	return &pb.ListCouponsResponse{Coupons: coupons}, nil
}
