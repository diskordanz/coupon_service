package model

import(
	"time"
	"fmt"
	"github.com/jinzhu/gorm"
	pb "github.com/diskordanz/coupon_service/proto"
	_ "github.com/lib/pq"
)

type CouponORM struct{ //сделать перегонку, переделать методы create, update, переписать базу, добавить пагинацию, парсер даты и времени
	Id uint64
	Name string
	Code string
	Description string
	CouponType int32
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
	*gorm.DB
}

func ToORM(c *pb.Coupon) *CouponORM{
	coupon := &CouponORM{}
	coupon.Name = c.GetName()
	coupon.Code = c.GetCode()
	coupon.Description = c.GetDescription()
	coupon.CouponType = int32(c.GetType())
	coupon.Status = c.GetStatus()
	coupon.Time_from = c.GetTimeFrom()
	coupon.Time_to = c.GetTimeTo()
	coupon.Date_from = c.GetDateFrom()
	coupon.Date_to = c.GetDateTo()
	coupon.Days = CheckDays(c.GetDays())
	coupon.Value = c.GetValue()
	coupon.FranchiseId = c.GetFranchiseId()

	return coupon
}

func CheckDays(d []pb.Coupon_DayOfWeek) []int32{

	var days []int32
	for _,k := range d {
		
		days = append(days, int32(k))
	}
	return days
}

func (db *GormDB) GetCoupon(req *pb.GetCouponRequest) (*pb.Coupon, error) {
	var coupon pb.Coupon
	if err := db.First(&coupon, req.Id).Error; err != nil {
		return nil, err
	}
	return &coupon, nil
}

func (db *GormDB) ListCoupons(req *pb.ListCouponsRequest) (*pb.ListCouponsResponse, error) {
	var coupons []*pb.Coupon
	if req.Filter == "" {
		if err := db.Find(&coupons).Error; err != nil {
			return nil, err
		}
	} else {
		filter := fmt.Sprintf("%%%s%%", req.Filter)
		if err := db.Where("name LIKE ?", filter).Find(&coupons).Error; err != nil {
			return nil, err
		}
	}

	return &pb.ListCouponsResponse{Coupons: coupons}, nil
}

func (db *GormDB) CreateCoupon(req *pb.CreateCouponRequest) (*pb.Coupon, error) {
	coupon := &pb.Coupon{Id: req.Coupon.Id}
	if err := db.Create(req.Coupon).Error; err != nil {
		return nil, err
	}
	return coupon, nil
}

func (db *GormDB) UpdateCoupon(req *pb.UpdateCouponRequest) (*pb.Coupon, error) {
	if err := db.Update(&req.Coupon).Error; err != nil {
		return nil, err
	}
	return req.Coupon, nil
}

func (db *GormDB) DeleteCoupon(req *pb.DeleteCouponRequest) error {
	if err := db.Delete(&pb.Coupon{Id: req.Id}).Error; err != nil {
		return  err
	}
	return nil
}

func (db *GormDB) ListCouponsByFranchise(req *pb.ListCouponsByFranchiseRequest) (*pb.ListCouponsResponse, error) {
	var coupons []*pb.Coupon
	filter := fmt.Sprintf("%%%s%%", req.GetFilter())
	if err := db.Where("franchise_id = ? AND name LIKE ?", req.Id,filter).Find(&coupons).Error; err != nil {
		return nil, err
	}
	return &pb.ListCouponsResponse{Coupons: coupons}, nil
}
