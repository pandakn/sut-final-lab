package entity

import (
	"fmt"
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestEmployeeValidate(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("check success", func(t *testing.T) {
		e := Employee{
			Name:       "test der",
			Email:      "test@mail.com",
			EmployeeID: "S12345678",
		}

		ok, err := govalidator.ValidateStruct(e)

		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())

		fmt.Print(ok)
	})
}
