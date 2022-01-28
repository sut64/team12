package entity

import (
	"time"
	"gorm.io/gorm"
	"github.com/asaskevich/govalidator"
	
)
type Status struct {
	gorm.Model
	Name            string
	Hospitalnet		[]Hospitalnet `gorm:"foreignKey:StatusID"`
}
type Province struct {
	gorm.Model
	Name            string
	Hospitalnet		[]Hospitalnet `gorm:"foreignKey:ProvinceID"`
}
type Genre struct {
	gorm.Model
	Name            string
	Hospitalnet		[]Hospitalnet `gorm:"foreignKey:GenreID"`
}

type Employee struct {
	gorm.Model
	Name            string
	Email           string `gorm:"uniqueIndex"`
	Password        string
	Hospitalnet		[]Hospitalnet`gorm:"goreignKey:EmployeeID"`
}

type Hospitalnet struct {
	gorm.Model
	Name            string
	Contract		float64 `valid:"IsPositive"`
	Address			string `valid:"minstringlength(4)"`//valid โดยเช็คว่ามีstring ไม่น้อยกว่า4 ตัว
	Adddate			time.Time `valid:"notpast"`

	EmployeeID		*uint
	Employee		Employee `gorm:references:id"`

	StatusID		*uint
	Status			Status `gorm:references:id"`

	ProvinceID		*uint
	Province		Province `gorm:references:id"`

	GenreID			*uint
	Genre			Genre `gorm:references:id"`

}

func init() {
	govalidator.CustomTypeTagMap.Set("IsPositive" , func(i interface{} , context interface{}) bool{
		value := i.(float64)
		return value >= 0
	})

	govalidator.CustomTypeTagMap.Set("notpast", func(i interface{} , context interface{}) bool{
		t := i.(time.Time)
		return t.After(time.Now())||t.Equal(time.Now())
	})
}