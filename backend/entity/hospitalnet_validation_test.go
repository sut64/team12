package entity

import (
	"testing"

	"time"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestHospitalContractIsPositive(t *testing.T) {
	g := NewGomegaWithT(t)

	hospitalnet := Hospitalnet{
		Name:     "ThaiHospital",
		Contract: -5, //ผิด
		Address:  "AAAAA",
		Adddate:  time.Now(),
	}

	ok, err := govalidator.ValidateStruct(hospitalnet)

	g.Expect(ok).ToNot(BeTrue())

	g.Expect(err).ToNot(BeNil())

	g.Expect(err.Error()).To(Equal("Contract cannot be negative or 0"))
}

func TestHospitalAddressMustbeInValidPattern(t *testing.T) {
	g := NewGomegaWithT(t)

	hospitalnet := Hospitalnet{
		Name:     "ThaiHospital",
		Contract: 5,
		Address:  "AA", //ผิด
		Adddate:  time.Now(),
	}

	ok, err := govalidator.ValidateStruct(hospitalnet)

	g.Expect(ok).ToNot(BeTrue())

	g.Expect(err).ToNot(BeNil())

	g.Expect(err.Error()).To(Equal("Adddress should more than 5 charactor"))
}

func TestHospitalAdddateMustBeNotPast(t *testing.T) {
	g := NewGomegaWithT(t)

	hospitalnet := Hospitalnet{
		Name:     "ThaiHospital",
		Contract: 5,
		Address:  "AABBC",
		Adddate:  time.Now().AddDate(0, 0, -4), //ผิด
	}

	ok, err := govalidator.ValidateStruct(hospitalnet)

	g.Expect(ok).ToNot(BeTrue())

	g.Expect(err).ToNot(BeNil())

	g.Expect(err.Error()).To(Equal("Date cannot be past"))
}
