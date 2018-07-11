package resolvers

import (
	"errors"
	"reflect"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Validate", func() {
	var item reflect.Type

	Context("Fail validation", func() {
		list := validateList{
			func(handler reflect.Type) error {
				return errors.New("Failed")
			},
		}

		It("should err", func() {
			Expect(list.run(item)).To(HaveOccurred())
		})
	})

	Context("Pass validation", func() {
		list := validateList{
			func(handler reflect.Type) error {
				return nil
			},
		}

		It("should not err", func() {
			Expect(list.run(item)).NotTo(HaveOccurred())
		})
	})
})
