package entity_test

import (
	"main/entity"
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestNegativeEmployee(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("Data is colleced", func(t *testing.T) {
		s := entity.Employees{
			Name:         "John",
			Salary:       14000,
			EmployeeCode: "hr-10240",
		}

		ok, err := govalidator.ValidateStruct(s)

		g.Expect(ok).ToNot(BeTrue())
		g.Expect(err).ToNot(BeNil())
		// g.Expect(err.Error()).To(Equal("EmployeeCode must be 2 uppercase English letters (A-Z) followed by '-' and 4 digits (0-9)"))
	})
}
