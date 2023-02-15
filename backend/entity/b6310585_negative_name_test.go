package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestEmployeeNameValidate(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("check name cannot be blank", func(t *testing.T) {
		e := Employee{
			Name:       "", // not collect
			Email:      "test@mail.com",
			EmployeeID: "S12345678",
		}

		ok, err := govalidator.ValidateStruct(e)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("Name: non zero value required"))
	})
}
