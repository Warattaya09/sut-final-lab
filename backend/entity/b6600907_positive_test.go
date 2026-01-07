package entity_test

import (
	"main/entity"
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestPositive(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("Data is colleced", func(t *testing.T) {
		s := entity.Employees{
			Name:         "John",
			Salary:       20000,
			EmployeeCode: "HR-1024",
		}

		ok, err := govalidator.ValidateStruct(s)

		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())
	})
}
