package entity

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func SetupDatabase() {
	database, err := gorm.Open(sqlite.Open("projectSE.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	database.AutoMigrate(
		&Employee{},
		&Buyinsurance{},
		&Hospitalnet{},
		&Genre{},
		&Status{},
		&Motive{},
		&Province{},
		&Invoice{},
		&Customer{},
		&InsuranceClaim{},
		&InvoicePayment{},
		&InsuranceConverage{},
		&PackageInsur{},
		&Protection{},
		&Totallist{},
	)

	db = database

	password, err := bcrypt.GenerateFromPassword([]byte("123456"), 14)

	// -- Employee Data
	db.Model(&Employee{}).Create(&Employee{
		Name:     "นายเอ",
		Email:    "aaa@example.com",
		Password: string(password),
	})
	db.Model(&Employee{}).Create(&Employee{
		Name:     "นายบี",
		Email:    "bbb@example.com",
		Password: string(password),
	})

	var a Employee
	var b Employee
	db.Raw("Select * FROM employees WHERE email = ?", "aaa@example.com").Scan(&a)
	db.Raw("Select * FROM employees WHERE email = ?", "bbb@example.com").Scan(&b)

	// -- Invoice Data
	i1 := Invoice{
		PaymentType: "Credit",
	}
	db.Model(&Invoice{}).Create(&i1)

	i2 := Invoice{
		PaymentType: "Cash",
	}
	db.Model(&Invoice{}).Create(&i2)

	// -- Customer Data
	c := Customer{
		Name:         "นายซี",
		IdNumber:     "1329911111111",
		PolicyNumber: "00000001",
	}
	db.Model(&Customer{}).Create(&c)

	d := Customer{
		Name:         "นายดี",
		IdNumber:     "1329922222222",
		PolicyNumber: "00000002",
	}
	db.Model(&Customer{}).Create(&d)

	f := Customer{
		Name:         "นายเอฟ",
		IdNumber:     "1329933333333",
		PolicyNumber: "00000003",
	}
	db.Model(&Customer{}).Create(&f)
	Protection1 := Protection{
		Protection_name: "ประกันชีวิต",
	}
	db.Model(&Protection{}).Create(&Protection1)

	Protection2 := Protection{
		Protection_name: "ประกันสุขภาพ",
	}
	db.Model(&Protection{}).Create(&Protection2)

	Protection3 := Protection{
		Protection_name: "ประกันอุบัติเหตุ",
	}
	db.Model(&Protection{}).Create(&Protection3)

	package1 := PackageInsur{
		Package_name: "แพ็คเกจ 1",
	}
	db.Model(&PackageInsur{}).Create(&package1)
	Package2 := PackageInsur{
		Package_name: "แพ็คเกจ 2",
	}
	db.Model(&PackageInsur{}).Create(&Package2)
	Package3 := PackageInsur{
		Package_name: "แพ็คเกจ 3",
	}
	db.Model(&PackageInsur{}).Create(&Package3)

	totallist1 := Totallist{
		Totallist_cost: "1,000,000 ต่อปี",
	}
	db.Model(&Totallist{}).Create(&totallist1)

	totallist2 := Totallist{
		Totallist_cost: "1,500,000 ต่อปี",
	}
	db.Model(&Totallist{}).Create(&totallist2)

	m1 := Motive{
		Name: "รถล้ม",
	}
	db.Model(&Motive{}).Create(&m1)

	m2 := Motive{
		Name: "ขาหัก",
	}
	db.Model(&Motive{}).Create(&m2)

	P1 := Province{
		Name: "กรุงเทพมหานคร",
	}
	db.Model(&Province{}).Create(&P1)

	P2 := Province{
		Name: "นครราชสีมา",
	}
	db.Model(&Province{}).Create(&P2)

	P3 := Province{
		Name: "ขอนแก่น",
	}
	db.Model(&Province{}).Create(&P3)

	G1 := Genre{
		Name: "รัฐ",
	}
	db.Model(&Genre{}).Create(&G1)

	G2 := Genre{
		Name: "เอกชน",
	}
	db.Model(&Genre{}).Create(&G2)

	S1 := Status{
		Name: "โรงพยาบาล",
	}
	db.Model(&Status{}).Create(&S1)

	S2 := Status{
		Name: "คลินิก",
	}
	db.Model(&Status{}).Create(&S2)

	// InsuranceConverage 1
	db.Model(&InsuranceConverage{}).Create(&InsuranceConverage{
		Premium:        10000,
		Protectiontype: "ผู้ป่วยนอก",
		Datetime:       time.Now(),
		Employee:       a,
		Protection:     Protection1,
		PackageInsur:   package1,
		Totallist:      totallist1,
	})
	// InsuranceConverage2
	db.Model(&InsuranceConverage{}).Create(&InsuranceConverage{
		Premium:        15000,
		Protectiontype: "ผู้ป่วยใน",
		Datetime:       time.Now(),
		Employee:       b,
		Protection:     Protection2,
		PackageInsur:   Package2,
		Totallist:      totallist2,
	})
}
