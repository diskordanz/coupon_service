package data

import(
	"time"
	"fmt"
	"log"
	"github.com/jinzhu/gorm"
	pb "github.com/diskordanz/coupon_service/proto"
	_ "github.com/lib/pq"
	"github.com/golang/protobuf/ptypes/wrappers"
	pq "github.com/mc2soft/pq-types"
)
  
type Repository interface{
	CreateCoupon(*pb.CreateCouponRequest) (*pb.Coupon, error)
    GetCoupon(*pb.GetCouponRequest) (*pb.Coupon, error)
    UpdateCoupon(*pb.UpdateCouponRequest) (*pb.Coupon, error)
    DeleteCoupon(*pb.DeleteCouponRequest) error
    ListCoupons(*pb.ListCouponsRequest) (*pb.ListCouponsResponse, error)
    ListCouponsByFranchise(*pb.ListCouponsByFranchiseRequest)(*pb.ListCouponsResponse, error)
}

type CouponRepository struct {
	DB *gorm.DB
}

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
	Value float32 
	FranchiseId uint64
	Days pq.Int32Array
}



func (db *CouponRepository) GetCoupon(req *pb.GetCouponRequest) (*pb.Coupon, error) {
	var coupon Coupon
	if err := db.DB.First(&coupon, req.Id).Error; err != nil {
		return nil, err
	}
	return ToPB(&coupon), nil
}

func (db *CouponRepository) CreateCoupon(req *pb.CreateCouponRequest) (*pb.Coupon, error) {
	var coupon Coupon
	
	if err := db.DB.Save(ToORM(req.Coupon)).Scan(&coupon).Error; err != nil {

		log.Println("no-no, when Create")
		log.Println(err)
		return nil, err
	}
	return req.Coupon, nil
}

func (db *CouponRepository) UpdateCoupon(req *pb.UpdateCouponRequest) (*pb.Coupon, error) {

	coupon := ToORM(req.Coupon)
		
	coupon.Id = req.Coupon.GetId().GetValue()

	if err := db.DB.Update(&coupon).Error; err != nil {
		return nil, err
	}
	return req.Coupon, nil
}

func (db *CouponRepository) DeleteCoupon(req *pb.DeleteCouponRequest) error {
	if err := db.DB.Delete(Coupon{Id: req.Id.GetValue()}).Error; err != nil {
		return  err
	}
	return nil
}

func (db *CouponRepository) ListCouponsByFranchise(req *pb.ListCouponsByFranchiseRequest) (*pb.ListCouponsResponse, error) {
	coupons:= []Coupon{}
	res := &pb.ListCouponsResponse{}
	
	if req.Filter == nil {
		if err := db.DB.Where("franchise_id = ?", req.Id).Find(&coupons).Error; err != nil {
			return nil, err
		}
	} else {
		filter := fmt.Sprintf("%%%s%%", req.GetFilter())
		if err := db.DB.Where("franchise_id = ? AND name LIKE ?", req.Id,filter).Find(&coupons).Error; err != nil {
			return nil, err
		}
	}

	for _, c := range coupons {
		res.Coupons = append(res.Coupons, ToPB(&c))
	}

	return res, nil
}


func (db *CouponRepository) ListCoupons(req *pb.ListCouponsRequest) (*pb.ListCouponsResponse, error) {
	var coupons []Coupon
	res := &pb.ListCouponsResponse{}

	if req.Filter == nil {
		if err := db.DB.Find(&coupons).Error; err != nil {
			return nil, err
		}
	} else {
		filter := fmt.Sprintf("%%%s%%", req.Filter)
		if err := db.DB.Where("name LIKE ?", filter).Find(&coupons).Error; err != nil {
			return nil, err
		}
	}

	for _, c := range coupons {
		res.Coupons = append(res.Coupons, ToPB(&c))
	}

	return res, nil
}



func DaysToORM(d []pb.Coupon_DayOfWeek) pq.Int32Array{

	var days []int32
	for _,k := range d {
		days = append(days, int32(pb.Coupon_DayOfWeek_value[k.String()]))
	}
	return pq.Int32Array(days)
}
func DaysToPB(d []int32) []pb.Coupon_DayOfWeek{

	var days []pb.Coupon_DayOfWeek
	for _,k := range d {
		days = append(days, pb.Coupon_DayOfWeek(k))
	}
	return days 
}

func ToORM(c *pb.Coupon) *Coupon {
	coupon := &Coupon{}
	coupon.Name = c.GetName().GetValue()
	coupon.Code = c.GetCode().GetValue()
	coupon.Description = c.GetDescription().GetValue()
	coupon.Type = int32(c.GetType())
	coupon.Status = c.GetStatus().GetValue()
	coupon.Time_from = time.Date(0, 0, 0, int(c.GetTimeFrom().GetHours().GetValue()), int(c.GetTimeFrom().GetMinutes().GetValue()), 0, 0, time.UTC)
	coupon.Time_to = time.Date( 0, 0, 0, int(c.GetTimeTo().GetHours().GetValue()), int(c.GetTimeTo().GetMinutes().GetValue()), 0, 0, time.UTC)
	coupon.Date_from = time.Date(int(c.GetDateFrom().GetYear().GetValue()), time.Month(c.GetDateFrom().GetMonth().GetValue()), int(c.GetDateFrom().GetDay().GetValue()), 0, 0, 0, 0, time.UTC)
	coupon.Date_to = time.Date(int(c.GetDateTo().GetYear().GetValue()), time.Month(c.GetDateTo().GetMonth().GetValue()), int(c.GetDateTo().GetDay().GetValue()), 0, 0, 0, 0, time.UTC)
	coupon.Days = DaysToORM(c.GetDays())
	coupon.Value = c.GetValue().GetValue()
	coupon.FranchiseId = c.GetFranchiseId().GetValue()
	return coupon
}

func TimeToPB(t *time.Time) *pb.TimeOfDay{
	time := &pb.TimeOfDay{}
	time.Hours = &wrappers.UInt32Value{ Value: uint32(t.Hour())}
	time.Minutes = &wrappers.UInt32Value{ Value: uint32(t.Minute())}
	return time
}

func DateToPB(t *time.Time) *pb.Date{ 
	date := &pb.Date{}
	date.Year = &wrappers.UInt32Value{ Value: uint32(t.Year())}
	date.Month = &wrappers.UInt32Value{ Value: uint32(t.Month())}
	date.Day = &wrappers.UInt32Value{ Value: uint32(t.Day())}
	return date
}

func ToPB(c *Coupon) *pb.Coupon{
	coupon := &pb.Coupon{}
	coupon.Name = &wrappers.StringValue{Value:c.Name}
	coupon.Code = &wrappers.StringValue{Value:c.Code}
	coupon.Description = &wrappers.StringValue{Value:c.Description}
	coupon.Type = pb.Coupon_CouponType(c.Type)
	coupon.Status = &wrappers.BoolValue{Value: c.Status}
	coupon.TimeFrom = TimeToPB(&c.Time_from)
	coupon.TimeTo = TimeToPB(&c.Time_to)
	coupon.DateFrom = DateToPB(&c.Date_from)
	coupon.DateTo = DateToPB(&c.Date_to)
	coupon.Days = DaysToPB(c.Days)
	coupon.Value = &wrappers.FloatValue{Value:c.Value}
	coupon.FranchiseId = &wrappers.UInt64Value{Value:c.FranchiseId}
	return coupon
}
