package entity_test

import (
	"main/entity"
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestNegativeSalary(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("Data is colleced", func(t *testing.T) {
		s := entity.Employees{
			Name:         "John",
			Salary:       14000,
			EmployeeCode: "HR-1024",
		}

		ok, err := govalidator.ValidateStruct(s)

		g.Expect(ok).ToNot(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("Salary must be between 15000 and 200000"))
	})
}
