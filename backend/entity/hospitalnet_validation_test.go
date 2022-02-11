package entity

import (
	"testing"

	"time"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestHospitalContractIsPositive(t *testing.T){
	g := NewGomegaWithT(t)

	hospitalnet := Hospitalnet{
		Name: "ThaiHospital",
		Contract: -5, //ผิด
		Address	: "AAAAA",
		Adddate: time.Now(),
	}

	ok, err := govalidator.ValidateStruct(hospitalnet)

	g.Expect(ok).ToNot(BeTrue())

	g.Expect(err).ToNot(BeNil())

	g.Expect(err.Error()).To(Equal("Contract must be positive"))
}

func TestHospitalAddressMustbeInValidPattern(t *testing.T){
	g := NewGomegaWithT(t)

	hospitalnet := Hospitalnet{
		Name: "ThaiHospital",
		Contract: 5,
		Address	: "AA", //ผิด
		Adddate: time.Now(),
	}

	ok, err := govalidator.ValidateStruct(hospitalnet)

	g.Expect(ok).ToNot(BeTrue())

	g.Expect(err).ToNot(BeNil())

	g.Expect(err.Error()).To(Equal("Address must be more than 5 character"))
}

func TestHospitalAdddateMustBeNotPast(t *testing.T){
	g := NewGomegaWithT(t)

	hospitalnet := Hospitalnet{
		Name: "ThaiHospital",
		Contract: 5,
		Address	: "AABBC",
		Adddate: time.Now().AddDate(0 , 0 , -4), //ผิด
	}

	ok, err := govalidator.ValidateStruct(hospitalnet)

	g.Expect(ok).ToNot(BeTrue())

	g.Expect(err).ToNot(BeNil())

	g.Expect(err.Error()).To(Equal("Date must be not past"))
}