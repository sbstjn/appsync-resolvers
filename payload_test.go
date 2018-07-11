package resolvers

import (
	"encoding/json"
	"reflect"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Payload", func() {
	Context("Invalid JSON", func() {
		message := payload{json.RawMessage(`{""name": "example"}`)}
		example := resolver{func(args struct {
			Name string `json:"name"`
		}) error {
			return nil
		}}

		args, err := message.parse(reflect.TypeOf(example.function).In(0))

		It("should error", func() {
			Expect(err).To(HaveOccurred())
		})

		It("should return nil", func() {
			Expect(args).To(BeNil())
		})
	})

	Context("Valid JSON and resolver with parameter", func() {
		message := payload{json.RawMessage(`{"name": "example"}`)}
		example := resolver{func(args struct {
			Name string `json:"name"`
		}) error {
			return nil
		}}

		args, err := message.parse(reflect.TypeOf(example.function).In(0))

		It("should not error", func() {
			Expect(err).NotTo(HaveOccurred())
		})

		It("should return struct", func() {
			Expect(args).NotTo(BeNil())
		})

		It("should parse data", func() {
			Expect(args).To(HaveLen(1))
			Expect(args[0].FieldByName("Name").String()).To(Equal("example"))
		})
	})
})
