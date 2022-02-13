package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Invoice struct {
	gorm.Model
	PaymentType     string
	InvoicePayments []InvoicePayment `gorm:"foreignKey:InvoiceID"`
}
type Customer struct {
	gorm.Model
	Name            string
	IdNumber        string           `gorm:"uniqueIndex"`
	PolicyNumber    string           `gorm:"uniqueIndex"`
	InvoicePayments []InvoicePayment `gorm:"foreignKey:CustomerID"`
	Buyinsurance    []Buyinsurance   `gorm:"foreignKey:CustomerID"`
	Paybacks        []Payback        `gorm:"foreignKey:CustomerID"`
}
type Status struct {
	gorm.Model
	Name        string
	Hospitalnet []Hospitalnet `gorm:"foreignKey:StatusID"`
}
type Province struct {
	gorm.Model
	Name        string
	Hospitalnet []Hospitalnet `gorm:"foreignKey:ProvinceID"`
}
type Genre struct {
	gorm.Model
	Name        string
	Hospitalnet []Hospitalnet `gorm:"foreignKey:GenreID"`
}

type Employee struct {
	gorm.Model
	Name               string
	Email              string `gorm:"uniqueIndex"`
	Password           string
	Hospitalnet        []Hospitalnet        `gorm:"foreignKey:EmployeeID"`
	InvoicePayments    []InvoicePayment     `gorm:"foreignKey:EmployeeID"`
	Buyinsurance       []Buyinsurance       `gorm:"foreignKey:EmployeeID"`
	Paybacks           []Payback            `gorm:"foreignKey:EmployeeID"`
	InsuranceConverage []InsuranceConverage `gorm:"foreignKey:EmployeeID"`
	InsuranceClaim     []InsuranceClaim     `gorm:"foreignKey:EmployeeID"`
}
type InvoicePayment struct {
	gorm.Model
	PaymentTime   time.Time `valid:"notpast~PaymentTime must not be in the past"`
	InvoiceNumber string    `valid:"matches(^[i]\\d{3}$)~InvoiceNumber is not correct"`
	PaymentAmount int       `valid:"IsPositive~PaymentAmount must be positive"`

	// InvoiceID ทำหน้าที่เป็น FK
	InvoiceID *uint
	Invoice   Invoice `gorm:"references:id" valid:"-"`

	// CustomerID ทำหน้าที่เป็น FK
	CustomerID *uint
	Customer   Customer `gorm:"references:id" valid:"-"`

	// EmployeeID ทำหน้าที่เป็น FK
	EmployeeID *uint
	Employee   Employee `gorm:"references:id" valid:"-"`
}
type Hospitalnet struct {
	gorm.Model
	Name     string
	Contract int       `valid:"IsPositive~Contract cannot be negative or 0"`
	Address  string    `valid:"minstringlength(5)~Adddress should more than 5 charactor"`
	Adddate  time.Time `valid:"notpast~Date cannot be past"`

	EmployeeID *uint
	Employee   Employee `gorm:"references:id"`

	StatusID *uint
	Status   Status `gorm:"references:id"`

	ProvinceID *uint
	Province   Province `gorm:"references:id"`

	GenreID *uint
	Genre   Genre `gorm:"references:id"`
}

type Buyinsurance struct {
	gorm.Model
	Consent      string    `valid:"required~You must accept Consent"`
	Healthinfrom string    `valid:"minstringlength(5)~Healthinfrom must be 5 or more "` //valid โดยเช็คว่ามีstring ไม่น้อยกว่า5 ตัว
	Adddate      time.Time `valid:"notpastnow~Time must not be in the past"`
	// InsuranceCoverageID ทำหน้าที่เป็น FK
	InsuranceConverageID *uint
	InsuranceConverage   InsuranceConverage `gorm:"references:id"`
	// CustomerID ทำหน้าที่เป็น FK
	EmployeeID *uint
	Employee   Employee `gorm:"references:id"`
	// EmployeeID ทำหน้าที่เป็น FK
	CustomerID *uint
	Customer   Customer `gorm:"references:id"`
}

type Payback struct {
	gorm.Model

	IDcard string  `gorm:"uniqueIndex" valid:"matches(^[0123456789]{13}$)~กรุณากรอกบัตรประจำตัวประชาชนให้ถูกต้อง"`
	Accout string  `valid:"matches(^[0123456789]{10}$)~กรุณากรอกเลขบัญชีให้ถูกต้อง"`
	Year   float32 `valid:"minamount~Year must not be negotive, required~Year must not be zero"`

	EmployeeID *uint
	Employee   Employee `gorm:"references:id" valid:"-"`

	ProtectionID *uint
	Protection   Protection `gorm:"references:id"`

	BankID *uint
	Bank   Bank `gorm:"references:id" valid:"-"`

	CustomerID *uint
	Customer   Customer `gorm:"references:id"`
}
type Bank struct {
	gorm.Model
	Bank_name string
	Payback   []Payback `gorm:"foreignKey:BankID"`
}
type Protection struct {
	gorm.Model
	Protection_name    string
	InsuranceConverage []InsuranceConverage `gorm:"foreignKey:ProtectionID"`
}
type PackageInsur struct {
	gorm.Model
	Package_name       string
	InsuranceConverage []InsuranceConverage `gorm:"foreignKey:PackageInsurID"`
}
type Totallist struct {
	gorm.Model
	Totallist_cost     string
	InsuranceConverage []InsuranceConverage `gorm:"foreignKey:TotallistID"`
}

type Motive struct {
	gorm.Model
	Name           string
	InsuranceClaim []InsuranceClaim `gorm:"foreignKey:MotiveID"`
}
type InsuranceClaim struct {
	gorm.Model
	Compensation int       `valid:"IsPositive~Compensation cannot be negative"`
	Insdate      time.Time `valid:"notpastnow~InsurDate must not be in the past"`
	Patient      string    `valid:"matches(^[0123456789]{13}$)~กรุณากรอกบัตรประจำตัวประชาชนให้ถูกต้อง"`
	EmployeeID   *uint
	Employee     Employee `gorm:"references:id" valid:"-"`

	MotiveID *uint
	Motive   Motive `gorm:"references:id" valid:"-"`
}

type InsuranceConverage struct {
	gorm.Model
	Protectiontype string    `valid:"minstringlength(5)~Protectiontype should more than 5 charactor"`
	Premium        int       `valid:"IsPositive~Premium cannot be negative or 0"`
	Datetime       time.Time `valid:"notpast~Datetime must not be in the past"`

	EmployeeID *uint
	Employee   Employee `gorm:"references:id"`

	ProtectionID *uint
	Protection   Protection `gorm:"references:id"`

	PackageInsurID *uint
	PackageInsur   PackageInsur `gorm:"references:id"`

	TotallistID *uint
	Totallist   Totallist `gorm:"references:id"`
}

func init() {
	govalidator.CustomTypeTagMap.Set("notpastnow", func(i interface{}, o interface{}) bool {
		t := i.(time.Time)
		// ย้อนหลังไม่เกิน 1 วัน
		return t.After(time.Now().AddDate(0, 0, -1))
	})
	govalidator.CustomTypeTagMap.Set("notpast",
		func(i interface{}, context interface{}) bool {
			t := i.(time.Time)
			if t.Year() >= time.Now().Year() {
				if int(t.Month()) >= int(time.Now().Month()) {
					if t.Day() >= time.Now().Day() {
						return true
					}
				}
			}
			return false
		})
	govalidator.CustomTypeTagMap.Set("IsPositive", func(i interface{}, context interface{}) bool {
		value := i.(int)
		return value >= 0
	})
	govalidator.CustomTypeTagMap.Set("past", func(i interface{}, context interface{}) bool {
		t := i.(time.Time)
		now := time.Now()
		return now.After(t)
	})
	govalidator.CustomTypeTagMap.Set("future", func(i interface{}, context interface{}) bool {
		t := i.(time.Time)
		now := time.Now()
		return now.Before(time.Time(t))
	})
	govalidator.CustomTypeTagMap.Set("present",
		func(i interface{}, context interface{}) bool {
			t := i.(time.Time)
			if t.Year() == time.Now().Year() {
				if int(t.Month()) == int(time.Now().Month()) {
					if t.Day() == time.Now().Day() {
						return true
					}
				}
			}
			return false
		})

	govalidator.CustomTypeTagMap.Set("minamount", func(i, o interface{}) bool {
		a := i.(float32)
		return a >= 1
	})

}
