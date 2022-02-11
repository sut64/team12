package entity

import (
	"testing"
	"time"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestPaymentTimeMustNotBePast(t *testing.T) {
	g := NewGomegaWithT(t)

	ip := InvoicePayment{
		PaymentTime: time.Now(),
	}

	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(ip)

	// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
	g.Expect(ok).ToNot(BeTrue())

	// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
	g.Expect(err).ToNot(BeNil())

	// err.Error ต้องมี error message แสดงออกมา
	g.Expect(err.Error()).To(Equal("WatchedTime must not be in the past"))
}

func TestInvoiceNumberMustBeInValidPattern(t *testing.T) {
	g := NewGomegaWithT(t)

	ip := InvoicePayment{
		InvoiceNumber: "i123",
	}
	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(ip)

	// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
	g.Expect(ok).ToNot(BeTrue())

	// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
	g.Expect(err).ToNot(BeNil())

	// err.Error ต้องมี error message แสดงออกมา
	g.Expect(err.Error()).To(Equal("InvoiceNumber does not match"))
}

func TestPaymentAmountMustPositive(t *testing.T) {
	g := NewGomegaWithT(t)

	ip := InvoicePayment{
		PaymentAmount: 123,
	}
	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(ip)

	// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
	g.Expect(ok).ToNot(BeTrue())

	// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
	g.Expect(err).ToNot(BeNil())

	// err.Error ต้องมี error message แสดงออกมา
	g.Expect(err.Error()).To(Equal("PaymentAmount must be positive"))
}
