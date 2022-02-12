package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestPaybackIDcard(t *testing.T) {
	g := NewGomegaWithT(t)

	ID := Payback{
		IDcard: "1111", //ผิด
		Accout: "1234567890",
		Year:   10.2,
	}

	ok, err := govalidator.ValidateStruct(ID)

	g.Expect(ok).ToNot(BeTrue())

	g.Expect(err).ToNot(BeNil())

	g.Expect(err.Error()).To(Equal("กรุณากรอกบัตรประจำตัวประชาชนให้ถูกต้อง"))
}

func TestPaybackAccout(t *testing.T) {
	g := NewGomegaWithT(t)

	Ac := Payback{
		IDcard: "1111222222444",
		Accout: "1234", //ผิด
		Year:   10.1,
	}

	ok, err := govalidator.ValidateStruct(Ac)

	g.Expect(ok).ToNot(BeTrue())

	g.Expect(err).ToNot(BeNil())

	g.Expect(err.Error()).To(Equal("กรุณากรอกเลขบัญชีให้ถูกต้อง"))
}

func TestPaybackYear(t *testing.T) {
	g := NewGomegaWithT(t)

	Y := Payback{
		IDcard: "1111222222444",
		Accout: "1234111111",
		Year:   0, //ผิด
	}

	ok, err := govalidator.ValidateStruct(Y)

	g.Expect(ok).ToNot(BeTrue())

	g.Expect(err).ToNot(BeNil())

	if err.Error() == "Year must not be zero" {
		g.Expect(err.Error()).To(Equal("Year must not be zero"))
	} else if err.Error() == "Year must not be negotive" {
		g.Expect(err.Error()).To(Equal("Year must not be negotive"))
	}
}
func TestPayback(t *testing.T) {
	g := NewGomegaWithT(t)

	// ข้อมูลถูกต้องหมดทุก field
	insr := Payback{
		IDcard: "1111222222423",
		Accout: "1234122111",
		Year:   5.2,
	}
	//ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(insr)

	//ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
	g.Expect(ok).To(BeTrue())

	// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
	g.Expect(err).To(BeNil())
}
