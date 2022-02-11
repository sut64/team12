package entity

import (
	"testing"

	"time"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

// ตรวจสอบค่าข้อมูลที่ถูกต้องทั้งหมด
func TestInsuranceConveragePass(t *testing.T) {
	g := NewGomegaWithT(t)

	// ข้อมูลถูกต้องหมดทุก field
	insuranceconverage := InsuranceConverage{
		Premium:        10000,
		Protectiontype: "PPPPP",
		Datetime:       time.Now(),
	}
	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(insuranceconverage)
	// ok ต้องเป็น true แปลว่าไม่มี error
	g.Expect(ok).ToNot(BeTrue())
	// err เป็นค่า nil แปลว่าไม่มี error
	g.Expect(err).ToNot(BeNil())
}

func TestPremiummonthIsPositive(t *testing.T) {
	g := NewGomegaWithT(t)

	insuranceconverage := InsuranceConverage{
		Premium:        -10, //ผิด
		Protectiontype: "PPPPP",
		Datetime:       time.Now(),
	}
	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(insuranceconverage)
	// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
	g.Expect(ok).ToNot(BeTrue())
	// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
	g.Expect(err).ToNot(BeNil())
	// err.Error ต้องมี error message แสดงออกมา
	g.Expect(err.Error()).To(Equal("Datetime must not be in the past;Premium cannot be negative or 0"))
}

func TestProtectiontypeMustbeInValidPattern(t *testing.T) {
	g := NewGomegaWithT(t)

	insuranceconverage := InsuranceConverage{
		Premium:        10000,
		Protectiontype: "AA", //ผิด
		Datetime:       time.Now(),
	}
	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(insuranceconverage)
	// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
	g.Expect(ok).ToNot(BeTrue())
	// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
	g.Expect(err).ToNot(BeNil())
	// err.Error ต้องมี error message แสดงออกมา
	g.Expect(err.Error()).To(Equal("Datetime must not be in the past;Protectiontype should more than 5 charactor"))
}

func TestDatetimeMustBeNotPast(t *testing.T) {
	g := NewGomegaWithT(t)

	insuranceconverage := InsuranceConverage{
		Premium:        10000,
		Protectiontype: "AABBCC",
		Datetime:       time.Now().AddDate(0, 0, -4), //ผิด
	}
	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(insuranceconverage)
	// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
	g.Expect(ok).ToNot(BeTrue())
	// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
	g.Expect(err).ToNot(BeNil())
	// err.Error ต้องมี error message แสดงออกมา
	g.Expect(err.Error()).To(Equal("Datetime must not be in the past"))
}
