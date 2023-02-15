package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestEmployeeIDValidate(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("check employee id is valid", func(t *testing.T) {
		e := Employee{
			Name:       "test der",
			Email:      "test@mail.com",
			EmployeeID: "A123456789", // not collect
		}

		ok, err := govalidator.ValidateStruct(e)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("Employee id is valid"))
	})
}
