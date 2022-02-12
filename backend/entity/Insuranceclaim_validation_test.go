package entity

import (
	"testing"

	"time"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestInsuranceClaimComp(t *testing.T) {
	g := NewGomegaWithT(t)

	comp := InsuranceClaim{
		Compensation: -100, //ผิด
		Patient:      "1321300115420",
		Insdate:      time.Now(),
	}

	ok, err := govalidator.ValidateStruct(comp)

	g.Expect(ok).ToNot(BeTrue())

	g.Expect(err).ToNot(BeNil())

	g.Expect(err.Error()).To(Equal("Compensation cannot be negative"))
}

func TestInsuranceClaimPat(t *testing.T) {
	g := NewGomegaWithT(t)

	pat := InsuranceClaim{
		Compensation: 100,
		Patient:      "13213001154205556", //ผิด
		Insdate:      time.Now(),
	}

	ok, err := govalidator.ValidateStruct(pat)

	g.Expect(ok).ToNot(BeTrue())

	g.Expect(err).ToNot(BeNil())

	g.Expect(err.Error()).To(Equal("กรุณากรอกบัตรประจำตัวประชาชนให้ถูกต้อง"))
}

func TestInsuranceClaimInsD(t *testing.T) {
	g := NewGomegaWithT(t)

	insd := InsuranceClaim{
		Compensation: 100,
		Patient:      "1321300115420",
		Insdate:      time.Date(1999, 10, 5, 0, 0, 0, 0, time.UTC), //ผิด
	}

	ok, err := govalidator.ValidateStruct(insd)

	g.Expect(ok).ToNot(BeTrue())

	g.Expect(err).ToNot(BeNil())

	g.Expect(err.Error()).To(Equal("InsurDate must not be in the past"))
}
func TestInsuranceClaim(t *testing.T) {
	g := NewGomegaWithT(t)

	// ข้อมูลถูกต้องหมดทุก field
	insr := InsuranceClaim{
		Compensation: 100,
		Patient:      "1321300115420",
		Insdate:      time.Now(),
	}
	//ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(insr)

	//ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
	g.Expect(ok).To(BeTrue())

	// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
	g.Expect(err).To(BeNil())
}
