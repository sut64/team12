package entity

import (
	"testing"
	"time"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestPaymentTimeMustNotBePast(t *testing.T) {
	g := NewGomegaWithT(t)

	ip1 := InvoicePayment{
		PaymentTime:   time.Date(1999, 10, 5, 0, 0, 0, 0, time.UTC), //ผิด
		InvoiceNumber: "i234",
		PaymentAmount: 123,
	}

	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(ip1)

	// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
	g.Expect(ok).ToNot(BeTrue())

	// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
	g.Expect(err).ToNot(BeNil())

	// err.Error ต้องมี error message แสดงออกมา
	g.Expect(err.Error()).To(Equal("PaymentTime must not be in the past"))
}

func TestInvoiceNumberMustBeInValidPattern(t *testing.T) {
	g := NewGomegaWithT(t)

	ip2 := InvoicePayment{
		PaymentTime:   time.Now(),
		InvoiceNumber: "i23456", //ผิด
		PaymentAmount: 123,
	}
	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(ip2)

	// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
	g.Expect(ok).ToNot(BeTrue())

	// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
	g.Expect(err).ToNot(BeNil())

	// err.Error ต้องมี error message แสดงออกมา
	g.Expect(err.Error()).To(Equal("InvoiceNumber is not correct"))
}

func TestPaymentAmountMustPositive(t *testing.T) {
	g := NewGomegaWithT(t)

	ip3 := InvoicePayment{
		PaymentTime:   time.Now(),
		InvoiceNumber: "i234",
		PaymentAmount: -123, //ผิด
	}
	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(ip3)

	// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
	g.Expect(ok).ToNot(BeTrue())

	// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
	g.Expect(err).ToNot(BeNil())

	// err.Error ต้องมี error message แสดงออกมา
	g.Expect(err.Error()).To(Equal("PaymentAmount must be positive"))
}
func TestPayment(t *testing.T) {
	g := NewGomegaWithT(t)

	ip4 := InvoicePayment{
		PaymentTime:   time.Now(),
		InvoiceNumber: "i234",
		PaymentAmount: 123,
	}
	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(ip4)

	// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
	g.Expect(ok).To(BeTrue())

	// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
	g.Expect(err).To(BeNil())
}
